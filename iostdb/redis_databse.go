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

func (rdb *RedisDatabase) PutHM(key []byte, args ...[]byte) error {
	newArgs := make([]interface{}, len(args)+1)
	newArgs[0] = key
	for i, v := range args {
		newArgs[i+1] = v
	}
	_, err := rdb.cli.Do("HMSET", newArgs...)
	return err
}
