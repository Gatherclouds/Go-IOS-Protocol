package network

import (
	"sync"
	"Go-IOS-Protocol/core/message"
)

type ReqType int32

const (
	ReqPublishTx     ReqType = iota
	ReqBlockHeight           //The height of the request to block
	RecvBlockHeight          //The height of the receiving block
	ReqNewBlock              // recieve a new block or a response for download block
	ReqDownloadBlock         // request for the height of block is equal to target

	MsgMaxTTL = 2
)

//Router Forwarding specific request to other components and sending messages for them
type Router interface {
	Init(base Network, port uint16) error
	FilteredChan(filter Filter) (chan message.Message, error)
	Run()
	Stop()
	Send(req message.Message)
	Broadcast(req message.Message)
	Download(start, end uint64) error
	CancelDownload(start, end uint64) error
}

var Route Router
var once sync.Once


