package protocol

import (
	"fmt"
	"github.com/iost-official/Go-IOS-Protocol/iosbase"
)

/*
Marked request types using by protocol
*/
type ReqType int

const (
	ReqPrePrepare ReqType = iota
	ReqPrepare
	ReqCommit
	ReqSubmitTxPack
	ReqPublishTx
	ReqNewBlock
)

type ResState int

const (
	Accepted ResState = iota
	Reject
	Error
)

//go:generate mockgen -destination mocks/mock_router.go -package protocol_mock github.com/iost-official/Go-IOS-Protocol/protocol Router

/*
Forwarding specific request to other components and sending messages for them
*/
type Router interface {
	Init(base iosbase.Network, port uint16) error
	FilteredChan(filter Filter) (chan iosbase.Request, chan iosbase.Response, error)
	Run()
	Stop()
	Send(req iosbase.Request) chan iosbase.Response
	Broadcast(req iosbase.Request)
}



