package network

import (
	"sync"
	"Go-IOS-Protocol/core/message"
	"fmt"
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

//GetInstance get singleton of network, [NOTE] conf.ListenAddr = your ip, port = !30304
func GetInstance(conf *NetConifg, target string, port uint16) (Router, error) {
	var err error
	once.Do(func() {
		baseNet, er := NewBaseNetwork(conf)
		if er != nil {
			err = er
			return
		}
		if target == "" {
			target = "base"
		}
		Route, err = RouterFactory(target)
		if err != nil {
			return
		}
		err = Route.Init(baseNet, port)
		if err != nil {
			return
		}
		Route.Run()
	})
	return Route, err
}

func RouterFactory(target string) (Router, error) {
	switch target {
	case "base":
		return &RouterImpl{}, nil
	}
	return nil, fmt.Errorf("target Router not found")
}

type RouterImpl struct {
	base Network

	chIn  <-chan message.Message
	chOut chan<- message.Message

	filterList  []Filter
	filterMap   map[int]chan message.Message
	knownMember []string
	ExitSignal  chan bool

	port uint16
}

func (r *RouterImpl) Init(base Network, port uint16) error {
	var err error
	r.base = base
	r.filterList = make([]Filter, 0)
	r.filterMap = make(map[int]chan message.Message)
	r.knownMember = make([]string, 0)
	r.ExitSignal = make(chan bool)
	r.port = port
	r.chIn, err = r.base.Listen(port)
	if err != nil {
		return err
	}
	return nil
}

//FilteredChan Get filtered request channel
func (r *RouterImpl) FilteredChan(filter Filter) (chan message.Message, error) {
	chReq := make(chan message.Message, 100)

	r.filterList = append(r.filterList, filter)
	r.filterMap[len(r.filterList)-1] = chReq

	return chReq, nil
}

func (r *RouterImpl) receiveLoop() {
	for true {
		select {
		case <-r.ExitSignal:
			r.base.Close(r.port)
			return
		case req := <-r.chIn:
			for i, f := range r.filterList {
				if f.check(req) {
					r.filterMap[i] <- req
				}
			}
		}
	}
}

func (r *RouterImpl) Run() {
	go r.receiveLoop()
}

func (r *RouterImpl) Stop() {
	r.ExitSignal <- true
}

func (r *RouterImpl) Send(req message.Message) {
	req.TTL = MsgMaxTTL
	r.base.Send(req)
}

// Broadcast to all known members
func (r *RouterImpl) Broadcast(req message.Message) {
	req.TTL = MsgMaxTTL
	r.base.Broadcast(req)
}

//download block with height >= start && height <= end
func (r *RouterImpl) Download(start uint64, end uint64) error {
	fmt.Println("sync:", start, end)
	if end < start {
		return fmt.Errorf("end should be greater than start")
	}
	return r.base.Download(start, end)
}

//CancelDownload cancel download
func (r *RouterImpl) CancelDownload(start uint64, end uint64) error {
	return r.base.CancelDownload(start, end)
}
