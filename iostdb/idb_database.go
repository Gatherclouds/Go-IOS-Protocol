package iostdb

import (
	"sync"
	"fmt"
)

type LDBDatabase struct {
	fn string
	db *leveldb.DB

	quitLock sync.Mutex
	quitChan chan chan error
}

func NewLDBDatabase(file string, cache int, handles int) (*LDBDatabase, error) {
	db, err := leveldb.OpenFile(file, &opt.Options{
		OpenFilesCacheCapacity: handles,
		BlockCacheCapacity:     cache / 2 * opt.MiB,
		WriteBuffer:            cache / 4 * opt.MiB,
		Filter:                 filter.NewBloomFilter(10),
	})

	fmt.Println(file, err)
	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(file, nil)
	}
	if err != nil {
		return nil, err
	}
	return &LDBDatabase{
		fn: file,
		db: db,
	}, nil
}

