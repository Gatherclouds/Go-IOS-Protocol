package p2p

import (
	"time"
	"log"
	"sync"
)

const (
	baseProtocolVersion    = 5
	baseProtocolLength     = uint64(16)
	baseProtocolMaxMsgSize = 2 * 1024
	snappyProtocolVersion  = 5
	pingInterval           = 15 * time.Second
)

type Peer struct {
	rw       *conn
	running  map[string]*protoRW
	log      log.Logger
	created  mclock.AbsTime
	wg       sync.WaitGroup
	protoErr chan error
	closed   chan struct{}
	disc     chan DiscReason // diconnect reason
	events   *event.Feed     // events receives message send / receive events if set
}

const (
	// devp2p message codes
	handshakeMsg = 0x00
	discMsg      = 0x01
	pingMsg      = 0x02
	pongMsg      = 0x03
	getPeersMsg  = 0x04
	peersMsg     = 0x05
)

type protoRW struct {
	Protocol
	in     chan Msg        // receives read messages
	closed <-chan struct{} // receives when peer is shutting down
	wstart <-chan struct{} // receives when write may start
	werr   chan<- error    // for write results
	offset uint64
	w      MsgWriter
}
