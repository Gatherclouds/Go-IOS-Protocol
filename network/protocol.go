package network

import (
	"Go-IOS-Protocol/network/discover"
)

type Protocol struct {

	Name string
	Version uint
	Length uint64

	Run func(peer *Peer, rw MsgReadWriter) error

	NodeInfo func() interface{}

	PeerInfo func(id discover.NodeID) interface{}
}
