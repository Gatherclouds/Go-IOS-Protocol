package iostdb

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type BlockChainImpl struct {
	db     *leveldb.DB
	redis  redis.Conn
	length int
}

const (
	DBPath   = "savedata/"
	IndexKey = "block_chain_index"
)

