package assets

import (
	"github.com/alydnhrealgang/moving/common/utils"
	"github.com/alydnhrealgang/moving/logs"
	"time"
)

func Create(store AssetStore, cache Cache) *Manager {
	return &Manager{
		logger: logs.Scope("ASSET-MGR"),
		cache:  cache,
		store:  store,
	}
}

type Manager struct {
	logger *logs.LogrusScope
	cache  Cache
	store  AssetStore
}

func (m *Manager) Save(path, name, contentType string, size int64) (*Asset, error) {
	ls := m.logger.WithField("path", path).
		WithField("name", name).
		WithField("contentType", contentType).
		WithField("size", size).
		Method("Save")

	a := &Asset{
		path:         path,
		name:         name,
		size:         size,
		contentType:  contentType,
		lastModified: time.Now().In(utils.China).Format(utils.DateTimeWithMicroLayout),
	}
	err := m.store.SaveAsset(a.ToData())
	if nil != err {
		return nil, err
	}

	err = m.cache.CacheAsset(a)
	if nil != err {
		ls.WithMethod("m.cache.CacheAssets(a)").Error(err)
	}

	return a, nil
}

func (m *Manager) Query(path, name string) ([]*Asset, error) {
	ls := m.logger.WithField("path", path).WithField("name", name).Method("query")
	am, err := m.ensure(path, ls)
	if nil != err {
		ls.Error(err)
		return nil, err
	}

	return am.ToSlice(name), nil
}

func (m *Manager) ensure(path string, ls *logs.LogrusScope) (AssetMap, error) {
	am, err := m.cache.GetAssets(path)
	if nil != err {
		ls.WithMethod("m.cache.GetAssets(path)").Error(err)
	}

	if nil == am || len(am) == 0 {
		data, err := m.store.QueryAssets(path)
		if nil != err {
			return nil, err
		}

		if utils.EmptyArray(data) {
			return nil, nil
		}

		am = make(AssetMap)
		for _, data := range data {
			am[data.Name] = &Asset{
				path:         data.Path,
				name:         data.Name,
				size:         data.Size,
				contentType:  data.ContentType,
				lastModified: data.LastModified,
			}
		}

		err = m.cache.CacheAssets(path, am)
		if nil != err {
			ls.WithMethod("m.cache.CacheAssets(am)").Error(err)
		}
	}

	return am, nil
}
