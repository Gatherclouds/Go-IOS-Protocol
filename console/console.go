package console

import (
	"sync"
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Console struct {
	cmds    []Cmd
	running bool
}

var Wg sync.WaitGroup
var Done = make(chan struct{})
var Nn *p2p.NaiveNetwork
var Db *iostdb.LDBDatabase

func (c *Console) Listen(prompt string) {
	for c.running {
		var cmd string
		fmt.Print(prompt)
		reader := bufio.NewReader(os.Stdin)
		cmd, _ = reader.ReadString('\n')
		args := strings.Fields(cmd)
		if len(args) == 0 {
			continue
		}
		fmt.Print(c.Run(args[0], args[1:]))
	}
}

