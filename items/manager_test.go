package items

import (
	"github.com/alydnhrealgang/moving/common/utils"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManagerSaveAndGet(t *testing.T) {
	store := NewMemory()
	manager := NewManager(store)

	box698500001 := NewItem("698500001", utils.EmptyString, "box", utils.EmptyString, utils.EmptyString, nil, nil, 1)
	box698500002 := NewItem("698500002", utils.EmptyString, "box", utils.EmptyString, utils.EmptyString, nil, nil, 1)
	err := manager.SaveItem(box698500001)
	assert.Nil(t, err)

	err = manager.SaveItem(box698500002)
	assert.Nil(t, err)

	item, err := manager.GetItem(box698500001.code)
	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.EqualValues(t, item, box698500001)

	item, err = manager.GetItem(box698500002.code)
	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.EqualValues(t, item, box698500002)

	box698500001.name = "box001"
	box698500001.description = "des001"
	box698500002.name = "box002"
	box698500002.description = "des002"
	err = manager.SaveItem(box698500001)
	assert.Nil(t, err)
	err = manager.SaveItem(box698500002)
	assert.Nil(t, err)

	manager = NewManager(store)

	item, err = manager.GetItem(box698500001.code)
	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.EqualValues(t, item, box698500001)
	assert.NotSame(t, item, box698500001)
	assert.NotEmpty(t, box698500001.name)
	assert.NotEmpty(t, box698500001.description)
	assert.Equal(t, box698500001.name, item.name)
	assert.Equal(t, box698500001.description, item.description)

	item, err = manager.GetItem(box698500002.code)
	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.EqualValues(t, item, box698500002)
	assert.NotSame(t, item, box698500002)
	assert.NotEmpty(t, box698500002.name)
	assert.NotEmpty(t, box698500002.description)
	assert.Equal(t, box698500002.name, item.name)
	assert.Equal(t, box698500002.description, item.description)
}

func TestManagerItemChildren(t *testing.T) {
	store := NewMemory()
	manager := NewManager(store)

	box698500001 := NewItem("698500001", utils.EmptyString, "box", utils.EmptyString, utils.EmptyString, nil, nil, 1)
	box698500002 := NewItem("698500002", utils.EmptyString, "box", utils.EmptyString, utils.EmptyString, nil, nil, 1)

	item003 := NewItem("698500003", box698500001.code, "article", utils.EmptyString, utils.EmptyString, nil, nil, 1)
	item004 := NewItem("698500004", box698500001.code, "article", utils.EmptyString, utils.EmptyString, nil, nil, 1)
	item005 := NewItem("698500005", box698500002.code, "article", utils.EmptyString, utils.EmptyString, nil, nil, 1)

	items := []*Item{box698500001, box698500002, item003, item004, item005}
	for _, item := range items {
		err := manager.SaveItem(item)
		assert.Nil(t, err)
	}

	children, exists, err := manager.GetChildren(box698500001.code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Len(t, children, 2)
	assert.True(t, lo.SomeBy(children, func(item *Item) bool {
		return item.code == item003.code
	}))
	assert.True(t, lo.SomeBy(children, func(item *Item) bool {
		return item.code == item004.code
	}))

	children, exists, err = manager.GetChildren(box698500002.code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Len(t, children, 1)
	assert.True(t, lo.SomeBy(children, func(item *Item) bool {
		return item.code == item005.code
	}))

	manager = NewManager(store)
	children, exists, err = manager.GetChildren(box698500001.code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Len(t, children, 2)
	assert.True(t, lo.SomeBy(children, func(item *Item) bool {
		return item.code == item003.code
	}))
	assert.True(t, lo.SomeBy(children, func(item *Item) bool {
		return item.code == item004.code
	}))

	children, exists, err = manager.GetChildren(box698500002.code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Len(t, children, 1)
	assert.True(t, lo.SomeBy(children, func(item *Item) bool {
		return item.code == item005.code
	}))

	moved, codesNotFound, err := manager.Move([]string{item003.code, item004.code}, box698500002.code)
	assert.Nil(t, err)
	assert.Empty(t, codesNotFound)
	assert.Len(t, moved, 2)

	children, exists, err = manager.GetChildren(box698500001.code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Len(t, children, 0)

	manager = NewManager(store)
	children, exists, err = manager.GetChildren(box698500001.code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Len(t, children, 0)

	children, exists, err = manager.GetChildren(box698500002.code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Len(t, children, 3)
	assert.True(t, lo.SomeBy(children, func(item *Item) bool {
		return item.code == item003.code
	}))
	assert.True(t, lo.SomeBy(children, func(item *Item) bool {
		return item.code == item004.code
	}))
	assert.True(t, lo.SomeBy(children, func(item *Item) bool {
		return item.code == item005.code
	}))

	moved, codesNotFound, err = manager.Move([]string{item003.code, item004.code, item005.code}, box698500001.code)
	assert.Nil(t, err)
	assert.Empty(t, codesNotFound)
	assert.Len(t, moved, 3)

	children, exists, err = manager.GetChildren(box698500002.code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Len(t, children, 0)
}
