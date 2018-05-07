package console

import "sync"

type Console struct {
	cmds    []Cmd
	running bool
}

var Wg sync.WaitGroup
var Done = make(chan struct{})
var Nn *p2p.NaiveNetwork
var Db *iostdb.LDBDatabase
