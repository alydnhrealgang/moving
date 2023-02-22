package store

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/alydnhrealgang/moving/assets"
	"github.com/alydnhrealgang/moving/common/utils"
	"github.com/alydnhrealgang/moving/items"
	"github.com/alydnhrealgang/moving/logs"
	"github.com/dgraph-io/badger/v3"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func NewBadger(path string) (*Badger, error) {
	logger := logs.Scope("BADGER")
	options := badger.DefaultOptions(path)
	options.WithLogger(logger)
	db, err := badger.Open(options)
	if nil != err {
		return nil, err
	}
	idSequence, err := db.GetSequence([]byte("ID_SEQ$$"), 1)
	journalPath := filepath.Join(path, "journey.log")
	journalLogWriter, err := os.OpenFile(journalPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_SYNC, os.FileMode(0660))
	if nil != err {
		return nil, err
	}

	return &Badger{
		db:               db,
		logger:           logger,
		idSequencer:      idSequence,
		journalLogWriter: journalLogWriter,
		journalLock:      sync.Mutex{},
	}, nil
}

type Badger struct {
	db               *badger.DB
	logger           *logs.LogrusScope
	idSequencer      *badger.Sequence
	journalLogWriter *os.File
	journalLock      sync.Mutex
}

func (s *Badger) GetSuggestionTexts(key, text string) (suggestionTexts []string, err error) {
	prefix := s.buildSuggestionPrefix(key)
	ls := s.logger.F("key", key).
		F("text", text).
		F("prefix", string(prefix)).
		Method("GetSuggestionTexts")

	err = s.db.View(func(txn *badger.Txn) error {
		options := badger.DefaultIteratorOptions
		options.PrefetchValues = false
		iter := txn.NewIterator(options)
		defer iter.Close()
		for iter.Seek(prefix); iter.ValidForPrefix(prefix); iter.Next() {
			if !iter.Valid() {
				continue
			}
			texts := strings.Split(string(iter.Item().Key()), "$$$$")
			if len(texts) == 2 {
				if len(text) == 0 || strings.Contains(texts[1], text) {
					suggestionTexts = append(suggestionTexts, texts[1])
				}
			}
		}
		return nil
	})

	if nil != err {
		ls.WM("s.db.View(func(txn *badger.Txn) error").Error(err)
		return
	}

	return
}

// The QueryItems function searches for items based on their itemType, tagName, and keywordOrTagValue,
// and allows for paging using startIndex and fetchSize parameters. The itemType parameter must be specified.
// Results returned with high performance include only the tags that have keys and values that fully match
// the tagName and keywordOrTagValue parameters.
// Results returned with low performance include items whose name, description or in props are fuzzy matched by the
// keywordOrTagValue parameter when the tagName is not specified.
func (s *Badger) QueryItems(itemType, tagName, keywordOrTagValue string, startIndex, fetchSize int64) (
	data []string, err error) {
	var prefix []byte
	queryByTag := !utils.EmptyOrWhiteSpace(tagName)
	if queryByTag {
		prefix = s.buildTagPrefix(itemType, tagName, keywordOrTagValue)
	} else {
		prefix = s.buildItemTypePrefix(itemType)
	}
	ls := s.logger.F("itemType", itemType).
		F("tagName", tagName).
		F("keywordOrTagValue", keywordOrTagValue).
		F("startIndex", startIndex).
		F("fetchSize", fetchSize).
		F("prefix", string(prefix)).
		M("QueryItems")

	err = s.db.View(func(txn *badger.Txn) error {
		options := badger.DefaultIteratorOptions
		iter := txn.NewIterator(options)
		defer iter.Close()
		currentIndex := int64(0)
		for iter.Seek(prefix); iter.ValidForPrefix(prefix); iter.Next() {
			if !iter.Valid() {
				continue
			}
			currentIndex++
			if currentIndex <= startIndex {
				continue
			}
			item := &items.ItemData{}
			key := string(iter.Item().Key())
			err := iter.Item().Value(func(val []byte) error {
				err := json.Unmarshal(val, item)
				if nil != err {
					ls.F("key", key).WM("json.Unmarshal(val, item)").Error(err)
					return nil
				}
				if queryByTag || item.FuzzyMatch(keywordOrTagValue) {
					data = append(data, item.Code)
				}
				return nil
			})

			if nil != err {
				ls.F("key", key).WM("iter.Item().Value(func(val []byte) error").Error(err)
				continue
			}

			if int64(len(data)) == fetchSize {
				break
			}
		}
		return nil
	})
	if nil != err {
		ls.WM("s.db.View(func(txn *badger.Txn) error").Error(err)
		err = fmt.Errorf("failed to query items from badger")
		return
	}
	return
}

func (s *Badger) DeleteItem(data *items.ItemData) (err error) {
	ls := s.logger.F("code", data.Code).M("DeleteItem")
	dbData, err := json.Marshal(data)
	if nil != err {
		ls.WM("json.Marshal(data)").Error(err)
		return fmt.Errorf("marshal data to json failed")
	}
	err = s.db.Update(func(txn *badger.Txn) error {
		changedTags, _ := data.DiffTags(nil)
		changedProps, _ := data.DiffProps(nil)
		key, pathKey, typeKey := s.buildItemKey(data.Code, data.ParentCode, data.ItemType)
		err := txn.Delete(key)
		if nil != err {
			ls.F("key", key).WM("txn.Delete(key)").Error(err)
			return fmt.Errorf("delete item from badger failed")
		}
		err = txn.Delete(pathKey)
		if nil != err {
			ls.F("pathKey", pathKey).WM("txn.Delete(pathKey)").Error(err)
			return fmt.Errorf("delete item path from badger failed")
		}
		err = txn.Delete(typeKey)
		if nil != err {
			ls.F("typeKey", typeKey).WM("txn.Delete(typeKey)").Error(err)
			return fmt.Errorf("delete item type from badger failed")
		}
		if nil != changedTags {
			for k, v := range changedTags {
				tagKey := s.buildTagKey(data.ItemType, k, v, data.Code)
				err = txn.Delete(tagKey)
				if nil != err {
					ls.F("tagKey", string(tagKey)).WM("txn.Set(tagKey, dbData)").Error(err)
					return fmt.Errorf("delete item tag failed")
				}
			}
		}
		if nil != changedProps {
			for k, v := range changedProps {
				err = s.incrementSuggestionRef(k, v, -1, txn, ls)
				if nil != err {
					ls.WM("s.incrementSuggestionRef(k, v, 1, txn, ls)").Error(err)
					return fmt.Errorf("decrement suggestion ref to badger failed")
				}
			}
		}
		return nil
	})

	if nil != err {
		ls.WM("s.db.Update(func(txn *badger.Txn) error").Error(err)
		return err
	}

	s.writeJournalLine("DeleteItem:", string(dbData))
	return nil
}

func (s *Badger) SaveItem(data *items.ItemData) (err error) {
	ls := s.logger.F("code", data.Code).M("SaveItem")
	dbData, err := json.Marshal(data)
	if nil != err {
		ls.WM("json.Marshal(data)").Error(err)
		return fmt.Errorf("marshal data to json failed")
	}
	codeData := []byte(data.Code)
	if nil != err {
		ls.WM("json.Marshal(data)").Error(err)
		return
	}
	err = s.db.Update(func(txn *badger.Txn) error {
		originalItem, err := s.getItem(data.Code, txn, ls)
		if nil != err {
			if err != badger.ErrKeyNotFound {
				ls.WM("s.getItem(data.Code, txn, ls)").Error(err)
				return fmt.Errorf("get item from badger failed")
			}
		}
		if nil == originalItem {
			id, err := s.idSequencer.Next()
			if nil != err {
				ls.WM("s.idSequencer.Next()").Error(err)
				return fmt.Errorf("generate id failed")
			}
			data.ID = fmt.Sprintf("BADGER-%d", id)
		}

		changedTags, deletedTags := data.DiffTags(originalItem)
		changedProps, deleteProps := data.DiffProps(originalItem)
		key, pathKey, typeKey := s.buildItemKey(data.Code, data.ParentCode, data.ItemType)

		err = txn.Set(key, dbData)
		if nil != err {
			ls.F("key", key).WM("txn.Set(key, dbData)").Error(err)
			return fmt.Errorf("set item by code to badger failed")
		}

		err = txn.Set(pathKey, codeData)
		if nil != err {
			ls.F("pathKey", pathKey).WM("txn.Set(pathKey, dbData)").Error(err)
			return fmt.Errorf("set item by path to badger failed")
		}

		err = txn.Set(typeKey, dbData)
		if nil != err {
			ls.F("typeKey", typeKey).WM("txn.Set(typeKey, dbData)").Error(err)
			return fmt.Errorf("set item based on typeKey to badger failed")
		}

		if nil != changedTags {
			for k, v := range changedTags {
				tagKey := s.buildTagKey(data.ItemType, k, v, data.Code)
				err = txn.Set(tagKey, dbData)
				if nil != err {
					ls.F("tagKey", string(tagKey)).WM("txn.Set(tagKey, dbData)").Error(err)
					return fmt.Errorf("set item tag to bager failed")
				}
			}
		}
		if nil != deletedTags {
			for k, v := range deletedTags {
				tagKey := s.buildTagKey(data.ItemType, k, v, data.Code)
				err = txn.Delete(tagKey)
				if nil != err {
					ls.F("tagKey", tagKey).WM("txn.Delete(tagKey)").Error(err)
					return fmt.Errorf("delete item tag from badger failed")
				}
			}
		}

		if nil != changedProps {
			for k, v := range changedProps {
				err = s.incrementSuggestionRef(k, v, 1, txn, ls)
				if nil != err {
					ls.WM("s.incrementSuggestionRef(k, v, 1, txn, ls)").Error(err)
					return fmt.Errorf("increment suggestion ref to badger failed")
				}
			}
		}

		if nil != deleteProps {
			for k, v := range deleteProps {
				err = s.incrementSuggestionRef(k, v, -1, txn, ls)
				if nil != err {
					ls.WM("s.incrementSuggestionRef(k,v,-1,txn,ls)").Error(err)
					return fmt.Errorf("decrement suggestion ref to badger failed")
				}
			}
		}

		return nil
	})

	if nil != err {
		ls.WM("s.db.Update(func(txn *badger.Txn) error").Error(err)
		return err
	}

	s.writeJournalLine("SaveItem:", string(dbData))
	return nil
}

func (s *Badger) incrementSuggestionRef(
	key string, text string, refValue int64, txn *badger.Txn, ls *logs.LogrusScope) error {

	suggestionKey := s.buildSuggestionKey(key, text)
	dbData, err := txn.Get(suggestionKey)
	ref := int64(0)
	if nil != err {
		if err != badger.ErrKeyNotFound {
			ls.F("suggestionKey", suggestionKey).WM("txn.Get(suggestionKey)").Error(err)
			return err
		} else if refValue <= 0 {
			// It will do nothing if refValue is negative or zero when suggestionKey doesn't exist
			return nil
		}
	} else {
		err = dbData.Value(func(val []byte) error {
			ref, _ = binary.Varint(val)
			return nil
		})
		if nil != err {
			ls.F("suggestionKey", suggestionKey).WM("dbData.Value(func(val []byte) error").Error(err)
			return err
		}
	}
	ref += refValue
	if ref > 0 {
		buf := make([]byte, binary.MaxVarintLen64)
		binary.PutVarint(buf, ref)
		err = txn.Set(suggestionKey, buf)
		if nil != err {
			ls.F("suggestionKey", suggestionKey).F("ref", ref).
				WM("txn.Set(suggestionKey, buf)").Error(err)
			return err
		}
	} else {
		// when the ref value less or equals than zero, it will be deleted.
		err = txn.Delete(suggestionKey)
		if nil != err {
			ls.F("suggestionKey", suggestionKey).WM("txn.Delete(suggestionKey)").Error(err)
			return err
		}
	}
	return nil
}

func (s *Badger) GetItem(code string) (item *items.ItemData, err error) {
	ls := s.logger.F("code", code).M("GetItem")
	err = s.db.View(func(txn *badger.Txn) error {
		item, err = s.getItem(code, txn, ls)
		if err != badger.ErrKeyNotFound {
			return err
		}
		return nil
	})
	if nil != err {
		ls.WM("s.db.View(func(txn *badger.Txn) error").Error(err)
		return
	}

	return
}

func (s *Badger) GetChildren(code string) (childCodes []string, err error) {
	prefix := s.buildItemPathPrefix(code)
	ls := s.logger.F("prefix", string(prefix)).M("GetChildren")
	err = s.db.View(func(txn *badger.Txn) error {
		options := badger.DefaultIteratorOptions
		iter := txn.NewIterator(options)
		defer iter.Close()
		for iter.Seek(prefix); iter.ValidForPrefix(prefix); iter.Next() {
			if !iter.Valid() {
				continue
			}
			_ = iter.Item().Value(func(val []byte) error {
				childCodes = append(childCodes, string(val))
				return nil
			})
		}
		return nil
	})
	if nil != err {
		ls.WM("s.db.View(func(txn *badger.Txn) error").Error(err)
		err = fmt.Errorf("get children from badger failed")
		return
	}

	return
}

// UpdateItemsParentCode will delete all items according to the original parent path of each item.
// Then, It will save them according to the new parent path of each item.
// Therefore, when GetChildren is called by the new code, these moved items will be fetched, but they cannot be fetched
// by original parent code.
func (s *Badger) UpdateItemsParentCode(moved []*items.ItemData, code string) (err error) {
	codesToMove := lo.Map(moved, func(item *items.ItemData, _ int) string { return item.Code })
	ls := s.logger.F("codesToMove", codesToMove).F("code", code).M("UpdateItemsParentCode")
	err = s.db.Update(func(txn *badger.Txn) error {
		for _, moveItem := range moved {
			if !utils.EmptyOrWhiteSpace(moveItem.ParentCode) {
				key, pathKey, _ := s.buildItemKey(moveItem.Code, moveItem.ParentCode, utils.EmptyString)
				err := txn.Delete(pathKey)
				if nil != err {
					ls.F("pathKey", string(pathKey)).WM("txn.Delete(pathKey)").Error(err)
					return fmt.Errorf("delete original path from badger failed")
				}
				_, pathKey, _ = s.buildItemKey(moveItem.Code, code, utils.EmptyString)
				moveItem = moveItem.Clone()
				moveItem.ParentCode = code
				data, err := json.Marshal(moveItem)
				if nil != err {
					ls.F("pathKey", string(pathKey)).WM("json.Marshal(moveItem)").Error(err)
					return fmt.Errorf("marhsal moveItem to json failed")
				}
				err = txn.Set(key, data)
				if nil != err {
					ls.F("key", string(key)).WM("txn.Set(pathKey)").Error(err)
					return fmt.Errorf("set item to badger failed")
				}

				err = txn.Set(pathKey, []byte(moveItem.Code))
				if nil != err {
					ls.F("pathKey", string(pathKey)).WM("txn.Set(pathKey)").Error(err)
					return fmt.Errorf("set item parent key to badger failed")
				}
			}
		}
		return nil
	})
	if nil != err {
		ls.WM("s.db.Update(func(txn *badger.Txn) error").Error(err)
		return fmt.Errorf("update items parent code failed")
	}
	s.writeJournalF("UpdateItemsParentCode:[%s] => [%s]\n", strings.Join(codesToMove, ","), code)
	return nil
}

func (s *Badger) getItem(code string, txn *badger.Txn, ls *logs.LogrusScope) (item *items.ItemData, err error) {
	key := s.buildItemCodeKey(code)
	dbItem, err := txn.Get(key)
	if nil != err {
		if err == badger.ErrKeyNotFound {
			ls.F("key", string(key)).Info("item not found")
			return
		}
		ls.F("key", string(key)).WM("txn.Get(key)").Error(err)
		err = errors.Errorf("get item by %s from badger failed", key)
		return
	}
	item = &items.ItemData{}
	err = dbItem.Value(func(val []byte) error {
		err := json.Unmarshal(val, item)
		if nil != err {
			ls.F("key", string(key)).WM("json.Unmarshal(val, item)").Error(err)
			return err
		}
		return nil
	})
	if nil != err {
		ls.F("key", string(key)).WM("dbItem.Value(func(val []byte) error").Error(err)
		return
	}
	return
}

func (s *Badger) Shutdown() error {
	err := s.journalLogWriter.Close()
	if nil != err {
		s.logger.M("Shutdown").WM("s.journalLogWriter.Close()").Error(err)
	}
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

func (s *Badger) DeleteAsset(asset *assets.AssetData) error {
	err := s.db.Update(func(txn *badger.Txn) error {
		key := s.buildAssetKey(asset.Path, asset.Name)
		return txn.Delete(key)
	})
	if nil != err {
		return errors.Wrap(err, "DeleteAssetFailed")
	}
	return nil
}

func (s *Badger) QueryAssets(path string) (data []*assets.AssetData, err error) {
	ls := s.logger.F("path", path).M("QueryAssets")
	data = make([]*assets.AssetData, 0, 1)
	err = s.db.View(func(txn *badger.Txn) error {
		options := badger.DefaultIteratorOptions
		key := s.buildAssetKey(path, utils.EmptyString)
		iter := txn.NewIterator(options)
		defer iter.Close()
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

func (s *Badger) buildItemKey(code, parentCode, itemType string) (key, pathKey, typeKey []byte) {
	key = s.buildItemCodeKey(code)
	parent := parentCode
	if utils.EmptyOrWhiteSpace(parent) {
		parent = "ROOT"
	}
	pathKey = []byte(fmt.Sprintf("ITEM_PATH$$%s$$%s$$", parent, code))
	typeKey = []byte(fmt.Sprintf("ITEM_TYPE$$%s$$%s$$", itemType, code))
	return
}

func (s *Badger) buildTagKey(itemType, k, v, code string) []byte {
	return []byte(fmt.Sprintf("ITEM_TAG$$%s$$%s$$%s$$%s$$", itemType, k, v, code))
}

func (s *Badger) buildTagPrefix(itemType, k, v string) []byte {
	return []byte(fmt.Sprintf("ITEM_TAG$$%s$$%s$$%s$$", itemType, k, v))
}

func (s *Badger) buildItemTypePrefix(itemType string) []byte {
	return []byte(fmt.Sprintf("ITEM_TYPE$$%s$$", itemType))
}

func (s *Badger) buildItemPathPrefix(code string) []byte {
	return []byte(fmt.Sprintf("ITEM_PATH$$%s$$", code))
}

func (s *Badger) buildItemCodeKey(code string) []byte {
	return []byte(fmt.Sprintf("ITEM$$%s$$", code))
}

func (s *Badger) buildSuggestionPrefix(key string) []byte {
	return []byte(fmt.Sprintf("SUGGEST$$%s$$$$", key))
}

func (s *Badger) buildSuggestionKey(key, text string) []byte {
	return []byte(fmt.Sprintf("SUGGEST$$%s$$$$%s", key, text))
}

func (s *Badger) writeJournal(args ...string) {
	s.journalLock.Lock()
	defer s.journalLock.Unlock()
	text := strings.Join(args, "")
	_, err := s.journalLogWriter.WriteString(text)
	if nil != err {
		panic(err)
	}
}

func (s *Badger) writeJournalLine(args ...string) {
	s.journalLock.Lock()
	defer s.journalLock.Unlock()
	text := strings.Join(args, " ")
	_, err := s.journalLogWriter.WriteString(text + "\n")
	if nil != err {
		panic(err)
	}
}

func (s *Badger) writeJournalF(f string, args ...any) {
	s.journalLock.Lock()
	defer s.journalLock.Unlock()
	text := fmt.Sprintf(f, args...)
	_, err := s.journalLogWriter.WriteString(text)
	if nil != err {
		panic(err)
	}
}

func (s *Badger) DumpAllKeys() error {
	return s.db.View(func(txn *badger.Txn) error {
		options := badger.DefaultIteratorOptions
		key := []byte("ITEM_PATH$$")
		iter := txn.NewIterator(options)
		defer iter.Close()
		for iter.Seek(key); iter.ValidForPrefix(key); iter.Next() {
			if !iter.Valid() {
				continue
			}
			s.logger.Info(string(iter.Item().Key()))
		}
		return nil
	})
}
