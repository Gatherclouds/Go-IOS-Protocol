package p2p

import (
	"net"
	"io/ioutil"
	"os"
)

type RequestHead struct {
	Length uint32 // Request的长度信息
}

const HEADLENGTH = 4

type Response struct {
	From        string
	To          string
	Code        int // http-like的状态码和描述
	Description string
}

// 最基本网络的模块API，之后gossip协议，虚拟的网络延迟都可以在模块内部实现
type Network interface {
	Send(req Request)
	Listen(port uint16) (<-chan Request, error)
	Close(port uint16) error
}

type NaiveNetwork struct {
	peerList *iostdb.LDBDatabase
	listen   net.Listener
	done     bool
}

func NewNaiveNetwork() *NaiveNetwork {
	dirname, _ := ioutil.TempDir(os.TempDir(), "p2p_test_")
	db, _ := iostdb.NewLDBDatabase(dirname, 0, 0)
	nn := &NaiveNetwork{
		//peerList: []string{"1.1.1.1", "2.2.2.2"},
		peerList: db,
		listen:   nil,
		done:     false,
	}
	nn.peerList.Put([]byte("1"), []byte("127.0.0.1:11037"))
	nn.peerList.Put([]byte("2"), []byte("127.0.0.1:11038"))
	return nn
}
