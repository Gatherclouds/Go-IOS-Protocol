package main

import (
	"flag"
	"fmt"
	"net"
)

var serverAddr = flag.String("s", "", "server port 30304, or other ports that have already started.")
var conf = initNetConf()

func initNetConf() *NetConifg {
	conf := &NetConifg{}
	conf.SetLogPath("/tmp")
	conf.SetNodeTablePath("/tmp")
	conf.SetListenAddr("0.0.0.0")
	return conf
}

func main() {
	bootnodeStart()
	//testBaseNetwork()
	//testBootNodeConn()
}

func testBootNodeConn() {
	node, err := discover.ParseNode("84a8ecbeeb6d3f676da1b261c35c7cd15ae17f32b659a6f5ce7be2d60f6c16f9@18.219.254.124:30304")

	if err != nil {
		fmt.Errorf("parse boot node got err:%v", err)
	}
	conn, err := net.Dial("tcp", node.Addr())
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Write([]byte("111"))
	conn.Close()
	fmt.Println("conn to remote success!")
}

//CancelDownload cancel downloading block with height between start and end
func (bn *BaseNetwork) CancelDownload(start, end uint64) error {
	bn.lock.Lock()
	defer bn.lock.Unlock()
	for ; start <= end; start++ {
		delete(bn.DownloadHeights, start)
	}
	return nil
}

//sendTo send request to the address
func (bn *BaseNetwork) sendTo(addr string, req *Request) {
	conn, err := bn.dial(addr)
	if err != nil {
		bn.log.E("[net] dial tcp got err:%v", err)
		return
	}
	if er := bn.send(conn, req); er != nil {
		bn.peers.RemoveByNodeStr(addr)
	}
}

