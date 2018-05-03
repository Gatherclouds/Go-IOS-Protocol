package discover

import (
	"net"
	"time"
)

type NodeID [NodeIDBits / 8]byte

const NodeIDBits = 512

type Node struct {
	IP       net.IP
	UDP, TCP uint16
	ID       NodeID

	sha common.Hash
	
	addedAt time.Time
}
