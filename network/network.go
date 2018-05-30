package network

import (
	"Go-IOS-Protocol/core/message"
	"Go-IOS-Protocol/db"
	"net"
	"io/ioutil"
	"os"
	"strconv"
)

type RequestHead struct {
	Length uint32 // length of Request
}

const (
	HEADLENGTH               = 4
	CheckKnownNodeInterval   = 10
	NodeLiveThresholdSeconds = 20
	MaxDownloadRetry         = 2
)

type Response struct {
	From        string
	To          string
	Code        int    // like http status code
	Description string // code status description
}

//Network api
type Network interface {
	Broadcast(req message.Message)
	Send(req message.Message)
	Listen(port uint16) (<-chan message.Message, error)
	Close(port uint16) error
	Download(start, end uint64) error
	CancelDownload(start, end uint64) error
}

type NaiveNetwork struct {
	db     *db.LDBDatabase //database of known nodes
	listen net.Listener
	conn   net.Conn
	done   bool
}

//NewNaiveNetwork create n peers
func NewNaiveNetwork(n int) (*NaiveNetwork, error) {
	dirname, err := ioutil.TempDir(os.TempDir(), "p2p_test_")
	if err != nil {
		return nil, err
	}
	db, err := db.NewLDBDatabase(dirname, 0, 0)
	if err != nil {
		return nil, err
	}
	nn := &NaiveNetwork{
		db:     db,
		listen: nil,
		done:   false,
	}
	for i := 1; i <= n; i++ {
		nn.db.Put([]byte(string(i)), []byte("0.0.0.0:"+strconv.Itoa(11036+i)))
	}
	return nn, nil
}

