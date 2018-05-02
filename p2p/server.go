package p2p

import "net"

type connFlag int

// Conn represents a connection between two nodes in the network
type conn struct {
	fd net.Conn
	transport
	flags connFlag
	cont  chan error      // The run loop uses cont to signal errors to SetupConn.
	id    discover.NodeID // valid after the encryption handshake
	caps  []Cap           // valid after the protocol handshake
	name  string          // valid after the protocol handshake
}

