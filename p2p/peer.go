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

