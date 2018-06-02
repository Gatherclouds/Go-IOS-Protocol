package network

import (
	"time"
	"bytes"
	"encoding/binary"
	"Go-IOS-Protocol/core/message"
	"strings"
	"net"
)

type NetReqType int16

const (
	Message NetReqType = iota + 1
	MessageReceived
	BroadcastMessage
	BroadcastMessageReceived
	Ping
	Pong
	ReqNodeTable
	NodeTable
)

// Request data structure exchanged by nodes
type Request struct {
	Version   [4]byte
	Length    int32 // length of request
	Timestamp int64
	Type      NetReqType
	FromLen   int16
	From      []byte
	Body      []byte
}

var NET_VERSION = [4]byte{'i', 'o', 's', 't'}

func isNetVersionMatch(buf []byte) bool {
	if len(buf) >= len(NET_VERSION) {
		return buf[0] == NET_VERSION[0] &&
			buf[1] == NET_VERSION[1] &&
			buf[2] == NET_VERSION[2] &&
			buf[3] == NET_VERSION[3]
	}
	return false
}

func newRequest(typ NetReqType, from string, data []byte) *Request {
	r := &Request{
		Version:   NET_VERSION,
		Timestamp: time.Now().UnixNano(),
		Type:      typ,
		FromLen:   int16(len(from)),
		From:      []byte(from),
		Body:      data,
	}
	//len(timestamp) + len(type) + len(fromLen) + len(from) + len(body)
	r.Length = int32(8 + 2 + 2 + len(r.From) + len(data))

	return r
}

func (r *Request) Pack() ([]byte, error) {
	var err error
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, &r.Version)
	err = binary.Write(buf, binary.BigEndian, &r.Length)
	err = binary.Write(buf, binary.BigEndian, &r.Timestamp)
	err = binary.Write(buf, binary.BigEndian, &r.Type)
	err = binary.Write(buf, binary.BigEndian, &r.FromLen)
	err = binary.Write(buf, binary.BigEndian, &r.From)
	err = binary.Write(buf, binary.BigEndian, &r.Body)
	return buf.Bytes(), err
}

//handle broadcast node's height
func (r *Request) msgHandle(net *BaseNetwork) {
	msg := &message.Message{}
	if _, err := msg.Unmarshal(r.Body); err == nil {
		switch msg.ReqType {
		case int32(RecvBlockHeight):
			var rh message.ResponseHeight
			rh.Decode(msg.Body)
			net.SetNodeHeightMap(string(r.From), rh.BlockHeight)
		default:
		}
	}
}

func (r *Request) handle(base *BaseNetwork, conn net.Conn) {
	switch r.Type {
	case Message:
		appReq := &message.Message{}
		if _, err := appReq.Unmarshal(r.Body); err == nil {
			base.log.D("[net] msg from =%v, to = %v, typ = %v,  ttl = %v", appReq.From, appReq.To, appReq.ReqType, appReq.TTL)
			base.RecvCh <- *appReq
		} else {
			base.log.E("[net] failed to unmarshal recv msg:%v, err:%v", r, err)
		}
		if er := base.send(conn, newRequest(MessageReceived, base.localNode.String(), common.Int64ToBytes(r.Timestamp))); er != nil {
			conn.Close()
		}
		r.msgHandle(base)
	case MessageReceived:
		base.log.D("[net] MessageReceived: %v", string(r.From), common.BytesToInt64(r.Body))
	case BroadcastMessage:
		appReq := &message.Message{}
		if _, err := appReq.Unmarshal(r.Body); err == nil {
			base.log.D("[net] msg from =%v, to = %v, typ = %v,  ttl = %v", appReq.From, appReq.To, appReq.ReqType, appReq.TTL)
			if appReq.ReqType == int32(ReqBlockHeight) {
				appReq.From = string(r.From)
			}
			base.RecvCh <- *appReq
			base.Broadcast(*appReq)
		}
		r.msgHandle(base)
	case BroadcastMessageReceived:
		//request for nodeTable
	case ReqNodeTable:
		base.putNode(string(r.From))
		base.peers.SetAddr(string(r.From), newPeer(conn, base.localNode.String(), conn.RemoteAddr().String()))
		addrs, err := base.AllNodesExcludeAddr(string(r.From))
		if err != nil {
			base.log.E("[net] failed to nodetable ", err)
		}
		req := newRequest(NodeTable, base.localNode.String(), []byte(strings.Join(addrs, ",")))
		if er := base.send(conn, req); er != nil {
			conn.Close()
		}
		//got nodeTable and save
	case NodeTable:
		base.putNode(string(r.Body))
	default:
		base.log.E("[net] wrong request :", r)
	}
}