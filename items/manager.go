package items

import (
	"fmt"
	"github.com/alydnhrealgang/moving/logs"
	"github.com/samber/lo"
	"sync"
)

func NewManager(store Store) *Manager {
	return &Manager{
		store:  store,
		items:  map[string]*Item{},
		lock:   sync.RWMutex{},
		logger: logs.Scope("ITEM-MGR"),
	}
}

type Manager struct {
	store  Store
	items  map[string]*Item
	lock   sync.RWMutex
	logger *logs.LogrusScope
}

func (m *Manager) Move(sourceCodes []string, toCode string) (moved []*Item, codesNotFound []string, err error) {
	ls := m.logger.F("sourceCodes", sourceCodes).F("toCode", toCode).M("Move")
	parent, err := m.ensureItem(toCode, true, ls)
	if nil != err {
		return
	}
	if nil == parent {
		codesNotFound = append(codesNotFound, toCode)
		return
	}
	moved = make([]*Item, 0, len(sourceCodes))
	m.lock.Lock()
	defer m.lock.Unlock()
	for _, code := range sourceCodes {
		item, iErr := m.ensureItem(code, false, ls)
		if nil != iErr {
			err = iErr
			return
		}
		if nil == item {
			codesNotFound = append(codesNotFound, code)
		}
		if item.parentCode == toCode {
			continue
		}
		moved = append(moved, item)
	}

	err = m.store.UpdateItemsParentCode(
		lo.Map(moved, func(item *Item, _ int) *ItemData { return item.ToData() }),
		toCode)

	if nil != err {
		ls.WM("m.store.UpdateItemsParentCode").Error(err)
		return
	}

	for index, item := range moved {
		if item.HasParent() {
			originalParent, ok := m.items[item.parentCode]
			if ok && nil != originalParent {
				originalParent.DeleteChildren(item.code)
			}
		}
		parent.AddChildrenIfLoaded(item)
		item.parentCode = toCode
		moved[index] = item.Clone()
	}
	return
}

func (m *Manager) GetItem(code string) (item *Item, err error) {
	ls := m.logger.F("code", code).M("GetItem")
	item, err = m.ensureItem(code, true, ls)
	if nil != item {
		m.lock.RLock()
		defer m.lock.RUnlock()
		item = item.Clone()
	}
	return
}

func (m *Manager) QueryItems(itemType, tagName, keywordOrTagValue string, startIndex, fetchSize int64) (
	items []*Item, err error) {

	ls := m.logger.F("itemType", itemType).
		F("tagName", tagName).
		F("keywordOrTagValue", keywordOrTagValue).
		F("startIndex", startIndex).
		F("fetchSize", fetchSize).
		M("QueryItems")

	codes, err := m.store.QueryItems(itemType, tagName, keywordOrTagValue, startIndex, fetchSize)
	if nil != err {
		ls.WM("m.store.QueryItem").Error(err)
		return
	}

	for _, code := range codes {
		item, err := m.ensureItem(code, true, ls)
		if nil != err {
			ls.F("code", "code").WM("m.ensureItem(code, true, ls)").Error(err)
			continue
		}
		items = append(items, item.Clone())
	}

	return
}

func (m *Manager) GetChildren(code string) (children []*Item, itemExists bool, err error) {
	ls := m.logger.F("code", code).M("GetChildren")
	item, err := m.ensureItem(code, true, ls)
	if nil != err {
		return
	}
	if nil == item {
		return
	}

	itemExists = true

	m.lock.RLock()
	defer m.lock.RUnlock()
	if !item.ChildrenLoaded() {
		m.lock.RUnlock()
		defer m.lock.RLock()
		err = m.loadChildren(item, ls)
		if nil != err {
			ls.WM("m.loadChildren(item, ls)").Error(err)
			return
		}
	}
	children = item.GetChildren(true)
	return
}

func (m *Manager) loadChildren(item *Item, ls *logs.LogrusScope) error {
	if !item.ChildrenLoaded() {
		childCodes, err := m.store.GetChildren(item.code)
		if nil != err {
			ls.WM("m.store.GetChildren(code)").Error(err)
			return err
		}
		item.children = make(map[string]*Item)
		if nil == childCodes || len(childCodes) == 0 {
			return nil
		}
		for _, code := range childCodes {
			child, err := m.ensureItem(code, false, ls)
			if nil != err {
				ls.WM("m.ensureItem(code, true, ls)").Error(err)
				continue
			}
			item.AddChildrenIfLoaded(child)
		}
	}
	return nil
}

func (m *Manager) SaveItem(item *Item) error {
	ls := m.logger.F("code", item.code).M("SaveItem")
	existItem, err := m.ensureItem(item.code, true, ls)
	if nil != existItem {
		err = func() error {
			m.lock.RLock()
			defer m.lock.RUnlock()
			if existItem.itemType != item.itemType {
				return fmt.Errorf("the itemType of a saved item cannot be altered")
			}
			if existItem.parentCode != item.parentCode {
				return fmt.Errorf("the parentCode of a saved item cannot be altered")
			}
			return nil
		}()
		if nil != err {
			return err
		}
	}
	var parent *Item
	if item.HasParent() {
		parent, err = m.ensureItem(item.parentCode, true, ls)
		if nil != err {
			ls.WM("m.ensureItem(item.parentCode, ls)").Error(err)
			return err
		}
		if nil == parent {
			ls.WM("m.ensureItem(item.parentCode, ls)").Error("parent cannot be nil")
			return fmt.Errorf("parent item not found")
		}
	}
	data := item.ToData()
	err = m.store.SaveItem(data)
	if nil != err {
		ls.WM("m.store.SaveItem(item.ToData())").Error(err)
		return err
	}
	if nil != existItem {
		m.lock.Lock()
		defer m.lock.Unlock()
		existItem.Apply(item)
	} else {
		m.lock.Lock()
		defer m.lock.Unlock()
		item.id = data.ID
		item = item.Clone()
		m.items[item.code] = item
		if item.HasParent() {
			parent.AddChildrenIfLoaded(item)
		}
	}

	return nil
}

func (m *Manager) ensureItem(code string, doLock bool, ls *logs.LogrusScope) (item *Item, err error) {
	ls.F("code", code).M("ensureItem")

	if doLock {
		m.lock.RLock()
		defer m.lock.RUnlock()
	}

	item, ok := m.items[code]
	if !ok || nil == item {
		if doLock {
			m.lock.RUnlock()
			defer m.lock.RLock()

			m.lock.Lock()
			defer m.lock.Unlock()
		}

		item, ok = m.items[code]
		if !ok || nil == item {
			var data *ItemData
			data, err = m.store.GetItem(code)
			if nil != err {
				ls.WM("m.store.GetItem(code)").Error(err)
				return
			}
			if nil == data {
				return
			}
			item = data.ToItem()
			m.items[code] = item
		}
	}
	return
}
