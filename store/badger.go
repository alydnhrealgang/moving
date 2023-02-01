package store

import (
	"encoding/json"
	"fmt"
	"github.com/alydnhrealgang/moving/assets"
	"github.com/alydnhrealgang/moving/common/utils"
	"github.com/alydnhrealgang/moving/logs"
	"github.com/dgraph-io/badger/v3"
	"github.com/pkg/errors"
)

func NewBadger(path string) (*Badger, error) {
	logger := logs.Scope("BADGER")
	options := badger.DefaultOptions(path)
	options.WithLogger(logger)
	db, err := badger.Open(options)
	if nil != err {
		return nil, err
	}
	return &Badger{
		db:     db,
		logger: logger,
	}, nil
}

type Badger struct {
	db     *badger.DB
	logger *logs.LogrusScope
}

func (s *Badger) Shutdown() error {
	return s.db.Close()
}

func (s *Badger) SaveAsset(asset *assets.AssetData) error {
	err := s.db.Update(func(txn *badger.Txn) error {
		key := s.buildAssetKey(asset.Path, asset.Name)
		bytes, _ := json.Marshal(asset)
		return txn.Set(key, bytes)
	})
	if nil != err {
		return errors.Wrap(err, "SaveAssetFailed")
	}
	return nil
}

func (s *Badger) QueryAssets(path string) (data []*assets.AssetData, err error) {
	ls := s.logger.F("path", path).M("QueryAssets")
	data = make([]*assets.AssetData, 0, 1)
	err = s.db.View(func(txn *badger.Txn) error {
		options := badger.DefaultIteratorOptions
		key := s.buildAssetKey(path, utils.EmptyString)
		iter := txn.NewKeyIterator(key, options)
		for iter.Seek(key); iter.ValidForPrefix(key); iter.Next() {
			if !iter.Valid() {
				continue
			}
			err = iter.Item().Value(func(val []byte) error {
				assetData := &assets.AssetData{}
				err = json.Unmarshal(val, assetData)
				if nil != err {
					ls.WM("json.Unmarshal(val, assetData)").Warning(err)
				} else {
					data = append(data, assetData)
				}
				return nil
			})
		}

		return nil
	})

	if nil != err {
		ls.WM("s.db.View(func(txn *badger.Txn) error").Error(err)
		err = errors.Wrap(err, "QueryAssetsFailed")
	}

	return
}

func (s *Badger) buildAssetKey(path, name string) []byte {
	return []byte(fmt.Sprintf("%s@%s", path, name))
}
