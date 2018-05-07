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

func (c *Console) Init(cmds ...Cmd) error {
	c.cmds = make([]Cmd, 0)
	c.running = true
	for _, cc := range cmds {
		c.RegisterCmd(cc)
	}
	return nil
}

func (c *Console) RegisterCmd(cmd Cmd) {
	c.cmds = append(c.cmds, cmd)
}
