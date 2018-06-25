package p2p

import (
	"net"
	"io/ioutil"
	"os"
	"fmt"
	"bytes"
	"encoding/binary"
	"strconv"
	"time"
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

func (nn *NaiveNetwork) Close(port uint16) error {
	port = 3 // 避免出现unused variable
	nn.done = true
	return nn.listen.Close()
}

func (nn *NaiveNetwork) Listen(port uint16) (<-chan Request, error) {
	var err error
	nn.listen, err = net.Listen("tcp", ":"+strconv.Itoa(int(port)))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return nil, err
	}
	fmt.Println("Listening on " + ":" + strconv.Itoa(int(port)))
	req := make(chan Request)

	conn := make(chan net.Conn)

	// For every listener spawn the following routine
	go func(l net.Listener) {
		for {
			//fmt.Println("new conn1", port)
			c, err := l.Accept()
			//fmt.Println("new conn", port)
			if err != nil {
				// handle error
				conn <- nil
				return
			}
			conn <- c
		}
	}(nn.listen)

	go func() {
		for {
			select {
			case c := <-conn:
				if c == nil {
					if nn.done {
						return
					}
					fmt.Println("Error accepting: ")
					break
				}

				go func(conn net.Conn) {
					defer conn.Close()
					// Make a buffer to hold incoming data.
					buf := make([]byte, HEADLENGTH)
					// Read the incoming connection into the buffer.
					_, err := conn.Read(buf)
					if err != nil {
						fmt.Println("Error reading request head:", err.Error())
					}

				}(c)
		}
	}
	return req, nil
}
