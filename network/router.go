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

