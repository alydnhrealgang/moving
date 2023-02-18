package store

import (
	"github.com/alydnhrealgang/moving/common/utils"
	"github.com/alydnhrealgang/moving/items"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func TestManagerSaveAndGet(t *testing.T) {
	store, err := NewBadger("test")
	assert.Nil(t, err)
	defer func() {
		err := store.Shutdown()
		assert.Nil(t, err)
		err = os.RemoveAll("test")
		assert.Nil(t, err)
	}()

	manager := items.NewManager(store)

	box698500001 := items.NewItem("698500001", utils.EmptyString, "box", utils.EmptyString, utils.EmptyString, nil, nil, 1)
	box698500002 := items.NewItem("698500002", utils.EmptyString, "box", utils.EmptyString, utils.EmptyString, nil, nil, 1)
	err = manager.SaveItem(box698500001)
	assert.Nil(t, err)

	err = manager.SaveItem(box698500002)
	assert.Nil(t, err)

	item, err := manager.GetItem(box698500001.Code())
	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.EqualValues(t, item, box698500001)

	item, err = manager.GetItem(box698500002.Code())
	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.EqualValues(t, item, box698500002)

	box698500001.SetName("box001")
	box698500001.SetDescription("des001")
	box698500002.SetName("box002")
	box698500002.SetDescription("des002")
	err = manager.SaveItem(box698500001)
	assert.Nil(t, err)
	err = manager.SaveItem(box698500002)
	assert.Nil(t, err)

	manager = items.NewManager(store)

	item, err = manager.GetItem(box698500001.Code())
	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.EqualValues(t, item, box698500001)
	assert.NotSame(t, item, box698500001)
	assert.NotEmpty(t, box698500001.Name())
	assert.NotEmpty(t, box698500001.Description())
	assert.Equal(t, box698500001.Name(), item.Name())
	assert.Equal(t, box698500001.Description(), item.Description())

	item, err = manager.GetItem(box698500002.Code())
	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.EqualValues(t, item, box698500002)
	assert.NotSame(t, item, box698500002)
	assert.NotEmpty(t, box698500002.Name())
	assert.NotEmpty(t, box698500002.Description())
	assert.Equal(t, box698500002.Name(), item.Name())
	assert.Equal(t, box698500002.Description(), item.Description())
}

func TestQuery(t *testing.T) {
	store, err := NewBadger("test")
	assert.Nil(t, err)
	defer func() {
		err := store.Shutdown()
		assert.Nil(t, err)
		err = os.RemoveAll("test")
		assert.Nil(t, err)
	}()

	manager := items.NewManager(store)
	boxes := make([]*items.Item, 0, 10)
	for i := 0; i < 10; i++ {
		boxes = append(boxes, items.NewItem(
			"box-"+strconv.Itoa(i),
			utils.EmptyString,
			"box",
			"box-name-"+strconv.Itoa(i),
			"desc-"+strconv.Itoa(i),
			map[string]string{"a": "b", "b": "c"},
			map[string]string{"type": "box-prop"},
			1,
		))
	}
	for _, box := range boxes {
		err = manager.SaveItem(box)
		assert.Nil(t, err)
	}

	suggestionTexts, err := store.GetSuggestionTexts("type", "box")
	assert.Nil(t, err)
	assert.Len(t, suggestionTexts, 1)

	queriedBoxes, err := manager.QueryItems("box", "a", "c", 0, 10)
	assert.Nil(t, err)
	assert.Empty(t, queriedBoxes)

	queriedBoxes, err = manager.QueryItems("box", "b", "c", 0, 5)
	assert.Nil(t, err)
	assert.Len(t, queriedBoxes, 5)

	queriedBoxes, err = manager.QueryItems("box", "b", "c", 5, 4)
	assert.Nil(t, err)
	assert.Len(t, queriedBoxes, 4)

	queriedBoxes, err = manager.QueryItems("box", "b", "c", 9, 2)
	assert.Nil(t, err)
	assert.Len(t, queriedBoxes, 1)

	queriedBoxes, err = manager.QueryItems("box", "", "desc", 0, 11)
	assert.Nil(t, err)
	assert.Len(t, queriedBoxes, 10)

	queriedBoxes, err = manager.QueryItems("box", "", "name", 0, 11)
	assert.Nil(t, err)
	assert.Len(t, queriedBoxes, 10)

	queriedBoxes, err = manager.QueryItems("box", "", "prop", 0, 11)
	assert.Nil(t, err)
	assert.Len(t, queriedBoxes, 10)

	for index, box := range boxes {
		box.SetName("箱子" + box.Code())
		box.SetDescription("描述" + box.Code())
		box.SetTags("a", utils.EmptyString)
		box.SetTags("b", utils.EmptyString)
		box.SetProp("type", utils.EmptyString)
		if index%2 == 0 {
			box.SetTags("大小", "很大")
			box.SetProp("材料", "布")
		} else {
			box.SetTags("大小", "很小")
			box.SetProp("材料", "铁")
		}
		err = manager.SaveItem(box)
		assert.Nil(t, err)
	}

	queriedBoxes, err = manager.QueryItems("box", "a", "b", 0, 10)
	assert.Nil(t, err)
	assert.Empty(t, queriedBoxes)

	queriedBoxes, err = manager.QueryItems("box", "b", "c", 0, 10)
	assert.Nil(t, err)
	assert.Empty(t, queriedBoxes)

	queriedBoxes, err = manager.QueryItems("box", "大小", "很大", 0, 5)
	assert.Nil(t, err)
	assert.Len(t, queriedBoxes, 5)

	queriedBoxes, err = manager.QueryItems("box", "大小", "很小", 0, 5)
	assert.Nil(t, err)
	assert.Len(t, queriedBoxes, 5)

	queriedBoxes, err = manager.QueryItems("box", "", "布", 0, 5)
	assert.Nil(t, err)
	assert.Len(t, queriedBoxes, 5)

	queriedBoxes, err = manager.QueryItems("box", "", "铁", 0, 5)
	assert.Nil(t, err)
	assert.Len(t, queriedBoxes, 5)

	suggestionTexts, err = store.GetSuggestionTexts("type", "box")
	assert.Nil(t, err)
	assert.Len(t, suggestionTexts, 0)

	suggestionTexts, err = store.GetSuggestionTexts("材料", "")
	assert.Nil(t, err)
	assert.Len(t, suggestionTexts, 2)
}

func TestManagerItemChildren(t *testing.T) {
	store, err := NewBadger("test")
	assert.Nil(t, err)
	defer func() {
		err := store.Shutdown()
		assert.Nil(t, err)
		err = os.RemoveAll("test")
		assert.Nil(t, err)
	}()

	manager := items.NewManager(store)

	box698500001 := items.NewItem("698500001", utils.EmptyString, "box", utils.EmptyString, utils.EmptyString, nil, nil, 1)
	box698500002 := items.NewItem("698500002", utils.EmptyString, "box", utils.EmptyString, utils.EmptyString, nil, nil, 1)

	item003 := items.NewItem("698500003", box698500001.Code(), "article", utils.EmptyString, utils.EmptyString, nil, nil, 1)
	item004 := items.NewItem("698500004", box698500001.Code(), "article", utils.EmptyString, utils.EmptyString, nil, nil, 1)
	item005 := items.NewItem("698500005", box698500002.Code(), "article", utils.EmptyString, utils.EmptyString, nil, nil, 1)

	allItems := []*items.Item{box698500001, box698500002, item003, item004, item005}
	for _, item := range allItems {
		err := manager.SaveItem(item)
		assert.Nil(t, err)
	}

	children, exists, err := manager.GetChildren(box698500001.Code())
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Len(t, children, 2)
	assert.True(t, lo.SomeBy(children, func(item *items.Item) bool {
		return item.Code() == item003.Code()
	}))
	assert.True(t, lo.SomeBy(children, func(item *items.Item) bool {
		return item.Code() == item004.Code()
	}))

	children, exists, err = manager.GetChildren(box698500002.Code())
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Len(t, children, 1)
	assert.True(t, lo.SomeBy(children, func(item *items.Item) bool {
		return item.Code() == item005.Code()
	}))

	manager = items.NewManager(store)
	children, exists, err = manager.GetChildren(box698500001.Code())
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Len(t, children, 2)
	assert.True(t, lo.SomeBy(children, func(item *items.Item) bool {
		return item.Code() == item003.Code()
	}))
	assert.True(t, lo.SomeBy(children, func(item *items.Item) bool {
		return item.Code() == item004.Code()
	}))

	children, exists, err = manager.GetChildren(box698500002.Code())
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Len(t, children, 1)
	assert.True(t, lo.SomeBy(children, func(item *items.Item) bool {
		return item.Code() == item005.Code()
	}))

	moved, codesNotFound, err := manager.Move([]string{item003.Code(), item004.Code()}, box698500002.Code())
	assert.Nil(t, err)
	assert.Empty(t, codesNotFound)
	assert.Len(t, moved, 2)

	children, exists, err = manager.GetChildren(box698500001.Code())
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Len(t, children, 0)

	manager = items.NewManager(store)
	children, exists, err = manager.GetChildren(box698500001.Code())
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Len(t, children, 0)

	children, exists, err = manager.GetChildren(box698500002.Code())
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Len(t, children, 3)
	assert.True(t, lo.SomeBy(children, func(item *items.Item) bool {
		return item.Code() == item003.Code()
	}))
	assert.True(t, lo.SomeBy(children, func(item *items.Item) bool {
		return item.Code() == item004.Code()
	}))
	assert.True(t, lo.SomeBy(children, func(item *items.Item) bool {
		return item.Code() == item005.Code()
	}))

	moved, codesNotFound, err = manager.Move([]string{item003.Code(), item004.Code(), item005.Code()}, box698500001.Code())
	assert.Nil(t, err)
	assert.Empty(t, codesNotFound)
	assert.Len(t, moved, 3)

	children, exists, err = manager.GetChildren(box698500002.Code())
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Len(t, children, 0)
}

func TestSaveItemWithChangedParentCode(t *testing.T) {
	store, err := NewBadger("test")
	assert.Nil(t, err)
	defer func() {
		err := store.Shutdown()
		assert.Nil(t, err)
		err = os.RemoveAll("test")
		assert.Nil(t, err)
	}()

	manager := items.NewManager(store)

	box698500000 := items.NewItem("698500000", utils.EmptyString, "box", utils.EmptyString, utils.EmptyString, nil, nil, 1)
	box698500001 := items.NewItem("698500001", utils.EmptyString, "box", utils.EmptyString, utils.EmptyString, nil, nil, 1)
	box698500002 := items.NewItem("698500002", "698500001", "box", utils.EmptyString, utils.EmptyString, nil, nil, 1)

	err = manager.SaveItem(box698500000)
	assert.Nil(t, err)
	err = manager.SaveItem(box698500001)
	assert.Nil(t, err)
	err = manager.SaveItem(box698500002)
	assert.Nil(t, err)

	box698500002.SetParentCode("698500000")
	err = manager.SaveItem(box698500002)
	assert.NotNil(t, err)

	box698500002.SetParentCode("698500001")
	err = manager.SaveItem(box698500002)
	assert.Nil(t, err)

}
