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

		Wg.Add(1)
		
		return fmt.Sprintf("Connected with port %d successfully!\n", port)
	}
	return c
}
