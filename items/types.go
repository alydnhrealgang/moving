package items

import (
	"github.com/alydnhrealgang/moving/common/utils"
	"github.com/samber/lo"
	"strings"
)

type Items map[string]*Item
type Suggests map[string][]string

type Store interface {
	SaveItem(data *ItemData) error
	GetItem(code string) (*ItemData, error)
	GetChildren(code string) ([]string, error)
	UpdateItemsParentCode(moved []*ItemData, code string) error
	QueryItems(itemType string, name string, value string, index int64, size int64) ([]string, error)
}

type ItemData struct {
	ID          string            `json:"id"`
	Code        string            `json:"code"`
	ParentCode  string            `json:"parentCode"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	ItemType    string            `json:"itemType"`
	Tags        map[string]string `json:"tags"`
	Props       map[string]string `json:"props"`
	Quantity    int               `json:"quantity"`
}

func (id *ItemData) DiffTags(other *ItemData) (changed, deleted map[string]string) {
	if nil == other {
		changed = id.Tags
		return
	}

	return diffMap(id.Tags, other.Tags)
}

func (id *ItemData) DiffProps(other *ItemData) (changed, deleted map[string]string) {
	if nil == other {
		changed = id.Props
		return
	}

	return diffMap(id.Props, other.Props)
}

func diffMap(m, o map[string]string) (changed, deleted map[string]string) {
	changed = make(map[string]string, 0)
	deleted = make(map[string]string, 0)
	if nil == m && nil == o {
		return
	}
	if nil == m {
		changed = o
		return
	}
	if nil == o {
		deleted = m
		return
	}

	for k, v := range m {
		if utils.EmptyOrWhiteSpace(v) && !utils.EmptyOrWhiteSpace(o[k]) {
			deleted[k] = o[k]
			continue
		}
		if v != o[k] {
			changed[k] = v
		}
	}

	for k, v := range o {
		if utils.EmptyOrWhiteSpace(m[k]) {
			delete(changed, k)
			deleted[k] = v
		}
	}
	return
}

type Item struct {
	id          string
	code        string
	parentCode  string
	name        string
	description string
	itemType    string
	tags        map[string]string
	props       map[string]string
	quantity    int
	children    map[string]*Item
}

func (i *Item) ID() string {
	return i.id
}

func (i *Item) Type() string {
	return i.itemType
}

func (i *Item) Code() string {
	return i.code
}

func (i *Item) Quantity() int {
	return i.quantity
}

func (i *Item) ParentCode() string {
	return i.parentCode
}

func (i *Item) Name() string {
	return i.name
}

func (i *Item) Description() string {
	return i.description
}

func (i *Item) SetName(name string) {
	i.name = name
}

func (i *Item) SetParentCode(code string) {
	i.parentCode = code
}

func (i *Item) SetDescription(desc string) {
	i.description = desc
}

func (i *Item) SetTags(k, v string) {
	if nil == i.tags {
		i.tags = make(map[string]string)
	}
	if utils.EmptyOrWhiteSpace(v) {
		delete(i.tags, k)
	} else {
		i.tags[k] = v
	}
}

func (i *Item) SetProp(k, v string) {
	if nil == i.props {
		i.props = make(map[string]string)
	}
	if utils.EmptyOrWhiteSpace(v) {
		delete(i.props, k)
	} else {
		i.props[k] = v
	}
}

func (i *Item) Tags() map[string]string {
	return utils.CopyMap(i.tags)
}

func (i *Item) Props() map[string]string {
	return utils.CopyMap(i.props)
}

func (i *Item) HasParent() bool {
	return !utils.EmptyOrWhiteSpace(i.parentCode)
}

func (i *Item) GetChildren(clone bool) []*Item {
	if nil == i.children {
		return nil
	}
	return lo.MapToSlice(i.children, func(_ string, child *Item) *Item {
		if clone {
			return child.Clone()
		}
		return child
	})
}

func (i *Item) ChildrenLoaded() bool {
	return nil != i.children
}

func (i *Item) DeleteChildren(code string) {
	if i.ChildrenLoaded() {
		delete(i.children, code)
	}
}

func (i *Item) AddChildrenIfLoaded(item *Item) {
	if i.ChildrenLoaded() {
		i.children[item.code] = item
	}
}

func (i *Item) Clone() *Item {
	return &Item{
		id:          i.id,
		code:        i.code,
		parentCode:  i.parentCode,
		name:        i.name,
		description: i.description,
		itemType:    i.itemType,
		tags:        utils.CopyMap(i.tags),
		props:       utils.CopyMap(i.props),
		quantity:    i.quantity,
		children:    utils.CopyMap(i.children),
	}
}

func (i *Item) Apply(other *Item) {
	if other == nil {
		return
	}

	i.name = other.name
	i.description = other.description
	i.quantity = other.quantity
	i.tags = utils.CopyMap(other.tags)
	i.props = utils.CopyMap(other.props)
}

func (id *ItemData) Clone() *ItemData {
	return &ItemData{
		ID:          id.ID,
		Code:        id.Code,
		ParentCode:  id.ParentCode,
		Name:        id.Name,
		Description: id.Description,
		ItemType:    id.ItemType,
		Tags:        utils.CopyMap(id.Tags),
		Props:       utils.CopyMap(id.Props),
		Quantity:    id.Quantity,
	}
}

func (id *ItemData) ToItem() *Item {
	var tags map[string]string
	if id.Tags != nil {
		tags = make(map[string]string, len(id.Tags))
		for k, v := range id.Tags {
			tags[k] = v
		}
	}

	var props map[string]string
	if id.Props != nil {
		props = make(map[string]string, len(id.Props))
		for k, v := range id.Props {
			props[k] = v
		}
	}

	return &Item{
		id:          id.ID,
		code:        id.Code,
		parentCode:  id.ParentCode,
		name:        id.Name,
		description: id.Description,
		itemType:    id.ItemType,
		tags:        utils.CopyMap(id.Tags),
		props:       utils.CopyMap(id.Props),
		quantity:    id.Quantity,
	}
}

func (id *ItemData) FuzzyMatch(value string) bool {
	if value == "*" {
		return true
	}
	if strings.Contains(id.Name, value) || strings.Contains(id.Description, value) {
		return true
	}

	if nil != id.Props {
		for _, v := range id.Props {
			if strings.Contains(v, value) {
				return true
			}
		}
	}

	return false
}

func (i *Item) ToData() *ItemData {
	return &ItemData{
		ID:          i.id,
		Code:        i.code,
		ParentCode:  i.parentCode,
		Name:        i.name,
		Description: i.description,
		ItemType:    i.itemType,
		Tags:        utils.CopyMap(i.tags),
		Props:       utils.CopyMap(i.props),
		Quantity:    i.quantity,
	}
}

func NewItem(code, parentCode, itemType, name, description string, tags, props map[string]string, count int) *Item {
	return &Item{
		code:        code,
		parentCode:  parentCode,
		name:        name,
		description: description,
		tags:        utils.CopyMap(tags),
		props:       utils.CopyMap(props),
		quantity:    count,
		itemType:    itemType,
		children:    nil, // Assume there are no children initially
	}
}
