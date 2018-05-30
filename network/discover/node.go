package discover

import (
	"net"
	"time"
	"Go-IOS-Protocol/common"
)

type NodeID string

type Node struct {
	IP       net.IP // len 4 for IPv4 or 16 for IPv6
	UDP, TCP uint16 // port numbers
	ID       NodeID // the node's public key

	sha common.Hash

	// Time when the node was added to the table.
	addedAt time.Time
}
