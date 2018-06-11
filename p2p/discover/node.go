package discover

import (
	"net"
	"time"
	"github.com/ethereum/go-ethereum/common"
	"errors"
	"strconv"
	"fmt"
)

type NodeID [NodeIDBits / 8]byte

const NodeIDBits = 512

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

// NewNode creates a new node. It is mostly meant to be used for
// testing purposes.
func NewNode(id NodeID, ip net.IP, udpPort, tcpPort uint16) *Node {
	if ipv4 := ip.To4(); ipv4 != nil {
		ip = ipv4
	}
	return &Node{
		IP:  ip,
		UDP: udpPort,
		TCP: tcpPort,
		ID:  id,
		//sha: crypto.Keccak256Hash(id[:]),
		//TODO: implement Keccak256
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

func ParseNode(nodeStr string) (node *Node, err error) {
	node = &Node{}
	nodeIdStrs := strings.Split(nodeStr, "@")
	if len(nodeIdStrs) < 2 {
		return node, fmt.Errorf("miss nodeId")
	}
	node.ID = NodeID(nodeIdStrs[0])
	tcpStr := strings.Split(nodeIdStrs[1], ":")
	node.IP = net.ParseIP(tcpStr[0])
	tcp, err := strconv.Atoi(tcpStr[1])
	if err != nil {
		return
	}
	node.TCP = uint16(tcp)
	return node, nil

}