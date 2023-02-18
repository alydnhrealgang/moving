package assets

import (
	"github.com/alydnhrealgang/moving/common/utils"
	"github.com/alydnhrealgang/moving/server/models"
	"io"
)

type Asset struct {
	path         string
	name         string
	size         int64
	contentType  string
	lastModified string
}

func (a Asset) Path() string {
	return a.path
}

func (a Asset) Name() string {
	return a.name
}

func (a Asset) Size() int64 {
	return a.size
}

func (a Asset) ContentType() string {
	return a.contentType
}

func (a Asset) LastModified() string {
	return a.lastModified
}

func (a Asset) ToData() *AssetData {
	return &AssetData{
		Path:         a.path,
		Name:         a.name,
		Size:         a.size,
		ContentType:  a.contentType,
		LastModified: a.lastModified,
	}
}

func (a Asset) ToModel() *models.AssetData {
	return &models.AssetData{
		LastModified: a.lastModified,
		Name:         a.name,
		Path:         a.path,
		Size:         a.size,
		ContentType:  a.contentType,
	}
}

type AssetStore interface {
	SaveAsset(asset *AssetData) error
	QueryAssets(path string) ([]*AssetData, error)
	DeleteAsset(asset *AssetData) error
}

type AssetData struct {
	Path         string
	Name         string
	Size         int64
	ContentType  string
	LastModified string
}

func ParseAsset(asset *Asset) *AssetData {
	return &AssetData{
		Path:         asset.Path(),
		Name:         asset.Name(),
		Size:         asset.Size(),
		ContentType:  asset.ContentType(),
		LastModified: asset.LastModified(),
	}
}

type Cache interface {
	CacheAsset(a *Asset) error
	GetAssets(path string) (AssetMap, error)
	DeleteAssets(path string) error
	CacheAssets(path string, am AssetMap) error
}

type ContentStore interface {
	UploadFrom(path, contentType string, r io.Reader) (size int64, err error)
	DownloadTo(path string, size int64, w io.Writer) error
}

type AssetMap map[string]*Asset

func (am AssetMap) ToSlice(name string) []*Asset {
	if !utils.EmptyOrWhiteSpace(name) {
		a, exists := am[name]
		if exists && nil != a {
			return []*Asset{a}
		}

		return nil
	}
	return utils.MapValues(am).([]*Asset)
}
