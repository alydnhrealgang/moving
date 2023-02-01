package cache

import (
	"github.com/alydnhrealgang/moving/assets"
	"sync"
)

func CreateMemory() *Memory {
	return &Memory{
		pathAssets: make(PathAssets),
		lock:       sync.RWMutex{},
	}
}

type Memory struct {
	pathAssets PathAssets
	lock       sync.RWMutex
}

func (m *Memory) CacheAsset(a *assets.Asset) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	am := m.pathAssets[a.Path()]
	if nil == am {
		am = make(assets.AssetMap)
		m.pathAssets[a.Path()] = am
	}
	am[a.Name()] = a
	return nil
}

func (m *Memory) GetAssets(path string) (assets.AssetMap, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.pathAssets[path], nil
}

func (m *Memory) CacheAssets(path string, am assets.AssetMap) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.pathAssets[path] = am
	return nil
}

type PathAssets map[string]assets.AssetMap
