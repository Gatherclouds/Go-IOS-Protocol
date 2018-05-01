package p2p


/*
Marked request types using by protocol
*/
type ReqType int

const (
	ReqPublishTx ReqType = iota
	ReqNewBlock
)

//go:generate mockgen -destination mocks/mock_router.go -package protocol_mock github.com/iost-official/PrototypeWorks/protocol Router

/*
Forwarding specific request to other components and sending messages for them
*/
type Router interface {
	Init(base core.Network, port uint16) error
	FilteredChan(filter Filter) (chan core.Request, error)
	Run()
	Stop()
	Send(req core.Request)
	Broadcast(req core.Request)
	Download (req core.Request) chan []byte
}
