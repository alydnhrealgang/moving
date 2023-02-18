package items

func NewMemory() *Memory {
	return &Memory{items: make(map[string]*ItemData)}
}

type Memory struct {
	items map[string]*ItemData
}

func (m *Memory) QueryItems(itemType string, name string, value string, index int64, size int64) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Memory) SaveItem(data *ItemData) error {
	m.items[data.Code] = data
	return nil
}

func (m *Memory) GetItem(code string) (*ItemData, error) {
	return m.items[code], nil
}

func (m *Memory) GetChildren(code string) (items []string, err error) {
	for _, item := range m.items {
		if item.ParentCode == code {
			items = append(items, item.Code)
		}
	}
	return
}

func (m *Memory) UpdateItemsParentCode(moved []*ItemData, code string) error {
	for _, data := range moved {
		m.items[data.Code].ParentCode = code
	}
	return nil
}
