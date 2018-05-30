package network

import "net"

type Peer struct {
	conn    net.Conn
	log     log.Logger
	local   string
	remote  string
	created mclock.AbsTime
	closed  chan struct{}
}

