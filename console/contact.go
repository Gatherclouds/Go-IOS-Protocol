package console

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"os"
)

func Connect() Cmd {
	c := Cmd{
		Name:  "connect",
		Usage: `connect PORT - Connect to the network. Listen to PORT`,
	}
	c.Exec = func(host *Console, args []string) string {
		if len(args) != 1 {
			return "Invalid arguments!\n"
		}
		port, err := strconv.Atoi(args[0])
		if err != nil {
			return "Invalid arguments!\n"
		}

		dirname, _ := ioutil.TempDir(os.TempDir(), min_framework.DbFile)
		Db, err = iostdb.NewLDBDatabase(dirname, 0, 0)
		if err != nil {
			return "Can't open database"
		}

		Nn = p2p.NewNaiveNetwork()
		lis, err := Nn.Listen(uint16(port))
		if err != nil {
			return fmt.Sprint(err) + "\n"
		}

	}
	return c
}
