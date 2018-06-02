package main

import "flag"

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

