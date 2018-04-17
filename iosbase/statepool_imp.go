package iosbase

import (
	"github.com/gomodule/redigo/redis"
)

type StatePoolImpl struct {
	cli redis.Conn
}

const (
	Conn   = "tcp"
	DBAddr = "localhost:6379"
)

