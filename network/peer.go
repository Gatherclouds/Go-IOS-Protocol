package network

import (
	"net"
	"Go-IOS-Protocol/common/mclock"
	"log"
	"os"

	"Go-IOS-Protocol/network/discover"
)

type Peer struct {
	conn    net.Conn
	log     log.Logger
	local   string
	remote  string
	created mclock.AbsTime
	closed  chan struct{}
}

func (p *Peer) Disconnect() {
	if p != nil && p.conn != nil {
		p.conn.Close()
	}
}

func newPeer(conn net.Conn, local, remote string) *Peer {
	return &Peer{
		conn:    conn,
		local:   local,
		remote:  remote,
		log:     *log.New(os.Stderr, "", 0), //TODO: 写专门的logger
		created: mclock.Now(),
		closed:  make(chan struct{}),
	}
}

func (ps *peerSet) Get(node *discover.Node) *Peer {
	ps.lock.RLock()
	defer ps.lock.RUnlock()
	if ps.peers == nil {
		return nil
	}
	peer, ok := ps.peers[node.String()]
	if !ok {
		return nil
	}
	return peer
}

func (ps *peerSet) SetAddr(addr string, p *Peer) {
	ps.lock.Lock()
	defer ps.lock.Unlock()
	p.remote = addr
	if ps.peers == nil {
		ps.peers = make(map[string]*Peer)
	}
	ps.peers[addr] = p
	return
}

func (ps *peerSet) Set(node *discover.Node, p *Peer) {
	ps.lock.Lock()
	defer ps.lock.Unlock()
	p.remote = node.String()
	if ps.peers == nil {
		ps.peers = make(map[string]*Peer)
	}
	ps.peers[node.String()] = p
	return
}

func (ps *peerSet) RemoveByNodeStr(nodeStr string) {
	node, _ := discover.ParseNode(nodeStr)
	ps.Remove(node)
}

func (ps *peerSet) Remove(node *discover.Node) {
	ps.lock.Lock()
	defer ps.lock.Unlock()
	ps.peers[node.String()].Disconnect()
	delete(ps.peers, node.String())
	return
}

//todo ping - pong check is conn  alive, remove it when it's dead