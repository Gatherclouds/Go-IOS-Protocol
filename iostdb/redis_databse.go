package iostdb

import (
	"github.com/gomodule/redigo/redis"
)

const (
	Conn   = "tcp"
	DBAddr = "localhost:6379"
)

type RedisDatabase struct {
	cli redis.Conn
}

func NewRedisDatabase() (*RedisDatabase, error) {
	dial, _ := redis.Dial(Conn, DBAddr)
	return &RedisDatabase{cli: dial}, nil
}

func (rdb *RedisDatabase) Put(key []byte, value []byte) error {
	_, err := rdb.cli.Do("SET", interface{}(key), interface{}(value))
	return err
}

