package main

import (
	"flag"
	"fmt"
	"net"
	"Go-IOS-Protocol/network/discover"
	"Go-IOS-Protocol/core/message"
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

//Filter The filter used by Router
// Rulers :
//     1. if both white list and black list are nil, this filter is all-pass
//     2. if one of those is not nil, filter as it is
//     3. if both of those list are not nil, filter as white list
type Filter struct {
	WhiteList  []message.Message
	BlackList  []message.Message
	RejectType []ReqType
	AcceptType []ReqType
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
