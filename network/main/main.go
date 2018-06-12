package main

import (
	"flag"
	"fmt"
	"net"
	"Go-IOS-Protocol/network/discover"
	"Go-IOS-Protocol/core/message"
	"strconv"
	"time"
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

func bootnodeStart() {
	node, err := discover.ParseNode("84a8ecbeeb6d3f676da1b261c35c7cd15ae17f32b659a6f5ce7be2d60f6c16f9@0.0.0.0:30304")
	if err != nil {
		fmt.Errorf("parse boot node got err:%v", err)
	}
	router, _ := RouterFactory("base")
	conf := initNetConf()
	conf.SetNodeID(string(node.ID))
	baseNet, err := NewBaseNetwork(conf)
	if err != nil {
		fmt.Println("NewBaseNetwork ", err)
		return
	}
	err = router.Init(baseNet, node.TCP)
	if err != nil {
		fmt.Println("Init ", err)
		return
	}
	go router.Run()
	fmt.Println("server starting", node.Addr())
	select {}
}

func testBaseNetwork() {
	rs := make([]Router, 0)
	for i := 0; i < 2; i++ {
		router, _ := RouterFactory("base")
		baseNet, _ := NewBaseNetwork(&NetConifg{NodeTablePath: "node_table_" + strconv.Itoa(i), ListenAddr: "192.168.1.34"})
		err := router.Init(baseNet, uint16(20002+i))
		if err != nil {
			fmt.Println("init net got err", err)
			return
		}
		router.Run()
		rs = append(rs, router)
	}
	time.Sleep(15 * time.Second)
	

}

func (f *Filter) check(req message.Message) bool {
	var memberCheck, typeCheck byte
	if f.WhiteList == nil && f.BlackList == nil {
		memberCheck = byte(0)
	} else if f.WhiteList != nil {
		memberCheck = byte(1)
	} else {
		memberCheck = byte(2)
	}
	if f.AcceptType == nil && f.RejectType == nil {
		typeCheck = byte(0)
	} else if f.AcceptType != nil {
		typeCheck = byte(1)
	} else {
		typeCheck = byte(2)
	}
	var m, t bool

	switch memberCheck {
	case 0:
		m = true
	case 1:
		m = memberContain(req.From, f.WhiteList)
	case 2:
		m = !memberContain(req.From, f.BlackList)
	}

	switch typeCheck {
	case 0:
		t = true
	case 1:
		t = reqTypeContain(req.ReqType, f.AcceptType)
	case 2:
		t = !reqTypeContain(req.ReqType, f.RejectType)
	}

	return m && t
}

func memberContain(a string, c []message.Message) bool {
	for _, m := range c {
		if m.From == a {
			return true
		}
	}
	return false
}

func reqTypeContain(a int32, c []ReqType) bool {
	for _, t := range c {
		if int32(t) == a {
			return true
		}
	}
	return false

}
