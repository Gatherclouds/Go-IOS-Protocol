package discover

import (
	"net"
	"time"
	"Go-IOS-Protocol/common"
	"errors"
	"strconv"
	"fmt"
)

type NodeID string

type Node struct {
	IP       net.IP // len 4 for IPv4 or 16 for IPv6
	UDP, TCP uint16 // port numbers
	ID       NodeID // the node's public key

	// This is a cached copy of sha3(ID) which is used for node
	// distance calculations. This is part of Node in order to make it
	// possible to write tests that need a node at a certain distance.
	// In those tests, the content of sha will not actually correspond
	// with ID.
	sha common.Hash

	// Time when the node was added to the table.
	addedAt time.Time
}

// NewNode creates a new node
func NewNode(id NodeID, ip net.IP, udpPort, tcpPort uint16) *Node {
	if ipv4 := ip.To4(); ipv4 != nil {
		ip = ipv4
	}
	return &Node{
		IP:  ip,
		UDP: udpPort,
		TCP: tcpPort,
		ID:  id,
	}
}

// Incomplete returns true for nodes with no IP address.
func (n *Node) Incomplete() bool {
	return n.IP == nil
}

// checks whether n is a valid complete node.
func (n *Node) validateComplete() error {
	if n.Incomplete() {
		return errors.New("incomplete node")
	}
	if n.UDP == 0 {
		return errors.New("missing UDP port")
	}
	if n.TCP == 0 {
		return errors.New("missing TCP port")
	}
	if n.IP.IsMulticast() || n.IP.IsUnspecified() {
		return errors.New("invalid IP (multicast/unspecified)")
	}
	return nil
}

func (n *Node) String() string {
	return string(n.ID) + "@" + n.IP.String() + ":" + strconv.Itoa(int(n.TCP))
}

func (n *Node) Addr() string {
	return n.IP.String() + ":" + strconv.Itoa(int(n.TCP))
}

// NodeID prints as a long hexadecimal number.
func (n NodeID) String() string {
	return fmt.Sprintf("%s", string(n))
}

func GenNodeId() NodeID {
	id := common.ToHex(common.Sha256(common.Int64ToBytes(time.Now().UnixNano())))
	return NodeID(id)
}
