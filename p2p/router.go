package p2p

import "fmt"

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

func RouterFactory(target string) (Router, error) {
	switch target {
	case "base":
		return &RouterImpl{}, nil
	}
	return nil, fmt.Errorf("target Router not found")
}

type RouterImpl struct {
	base core.Network

	chIn  <-chan core.Request
	chOut chan<- core.Request

	filterList  []Filter
	filterMap   map[int]chan core.Request
	knownMember []string
	ExitSignal  chan bool
}

func (r *RouterImpl) Init(base core.Network, port uint16) error {
	var err error

	r.base = base
	r.filterList = make([]Filter, 0)
	r.filterMap = make(map[int]chan core.Request)
	r.knownMember = make([]string, 0)
	r.ExitSignal = make(chan bool)

	r.chIn, err = r.base.Listen(port)
	if err != nil {
		return err
	}

	return nil
}

/*
Get filtered request channel
*/
func (r *RouterImpl) FilteredChan(filter Filter) (chan core.Request, error) {
	chReq := make(chan core.Request)

	r.filterList = append(r.filterList, filter)
	r.filterMap[len(r.filterList)-1] = chReq

	return chReq, nil
}