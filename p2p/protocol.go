package p2p

// Protocol represents a P2P subprotocol implementation.
type Protocol struct {

	Name string
	Version uint
	Length uint64

	Run func(peer *Peer, rw MsgReadWriter) error
	NodeInfo func() interface{}
	PeerInfo func(id discover.NodeID) interface{}
}
