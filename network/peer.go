package network

import (
	"net"
	"Go-IOS-Protocol/common/mclock"
	"log"
	"os"
)

type Peer struct {
	conn    net.Conn
	log     log.Logger
	local   string
	remote  string
	created mclock.AbsTime
	closed  chan struct{}
}

func (p *Peer) Disconnect() {
	if p != nil && p.conn != nil {
		p.conn.Close()
	}
}

func newPeer(conn net.Conn, local, remote string) *Peer {
	return &Peer{
		conn:    conn,
		local:   local,
		remote:  remote,
		log:     *log.New(os.Stderr, "", 0), //TODO: 写专门的logger
		created: mclock.Now(),
		closed:  make(chan struct{}),
	}
}
