package network

import (
	"Go-IOS-Protocol/core/message"
	"Go-IOS-Protocol/db"
	"Go-IOS-Protocol/network/discover"
	"Go-IOS-Protocol/log"
	"net"
	"io/ioutil"
	"os"
	"strconv"
	"fmt"
	"sync"
	"encoding/binary"
	"time"
	"bytes"
	"errors"
)

type RequestHead struct {
	Length uint32 // length of Request
}

const (
	HEADLENGTH               = 4
	CheckKnownNodeInterval   = 10
	NodeLiveThresholdSeconds = 20
	MaxDownloadRetry         = 2
)

type Response struct {
	From        string
	To          string
	Code        int    // like http status code
	Description string // code status description
}

//Network api
type Network interface {
	Broadcast(req message.Message)
	Send(req message.Message)
	Listen(port uint16) (<-chan message.Message, error)
	Close(port uint16) error
	Download(start, end uint64) error
	CancelDownload(start, end uint64) error
}

type NaiveNetwork struct {
	db     *db.LDBDatabase //database of known nodes
	listen net.Listener
	conn   net.Conn
	done   bool
}

//NewNaiveNetwork create n peers
func NewNaiveNetwork(n int) (*NaiveNetwork, error) {
	dirname, err := ioutil.TempDir(os.TempDir(), "p2p_test_")
	if err != nil {
		return nil, err
	}
	db, err := db.NewLDBDatabase(dirname, 0, 0)
	if err != nil {
		return nil, err
	}
	nn := &NaiveNetwork{
		db:     db,
		listen: nil,
		done:   false,
	}
	for i := 1; i <= n; i++ {
		nn.db.Put([]byte(string(i)), []byte("0.0.0.0:"+strconv.Itoa(11036+i)))
	}
	return nn, nil
}

func (nn *NaiveNetwork) Close(port uint16) error {
	port = 3 // 避免出现unused variable
	nn.done = true
	return nn.listen.Close()
}

func (nn *NaiveNetwork) Broadcast(req message.Message) error {
	iter := nn.db.NewIterator()
	for iter.Next() {
		addr, _ := nn.db.Get([]byte(string(iter.Key())))
		conn, err := net.Dial("tcp", string(addr))
		if err != nil {
			if dErr := nn.db.Delete([]byte(string(iter.Key()))); dErr != nil {
				fmt.Errorf("failed to delete peer : k= %v, v = %v, err:%v\n", iter.Key(), addr, err.Error())
			}
			fmt.Errorf("dialing to %v encounter err : %v\n", addr, err.Error())
			continue
		}
		nn.conn = conn
		go nn.Send(req)
	}
	iter.Release()
	return iter.Error()
}

func (nn *NaiveNetwork) Send(req message.Message) {
	if nn.conn == nil {
		fmt.Errorf("no connect in network")
		return
	}
	defer nn.conn.Close()

	reqHeadBytes, reqBodyBytes, err := reqToBytes(req)
	if err != nil {
		fmt.Errorf("reqToBytes encounter err : %v\n", err.Error())
		return
	}
	if _, err = nn.conn.Write(reqHeadBytes); err != nil {
		fmt.Errorf("sending request head encounter err :%v\n", err.Error())
	}
	if _, err = nn.conn.Write(reqBodyBytes[:]); err != nil {
		fmt.Errorf("sending request body encounter err : %v\n", err.Error())
	}
	return
}

func (nn *NaiveNetwork) Listen(port uint16) (<-chan message.Message, error) {
	var err error
	nn.listen, err = net.Listen("tcp", ":"+strconv.Itoa(int(port)))
	if err != nil {
		return nil, fmt.Errorf("Error listening: %v", err.Error())
	}
	fmt.Println("Listening on " + ":" + strconv.Itoa(int(port)))
	req := make(chan message.Message, 100)

	conn := make(chan net.Conn, 10)

	// For every listener spawn the following routine
	go func(l net.Listener) {
		for {
			c, err := l.Accept()
			if err != nil {
				// handle error
				conn <- nil
				return
			}
			conn <- c
		}
	}(nn.listen)
	go func() {
		for {
			select {
			case c := <-conn:
				if c == nil {
					if nn.done {
						return
					}
					fmt.Println("Error accepting: ")
					break
				}

				go func(conn net.Conn) {
					defer conn.Close()
					// Make a buffer to hold incoming data.
					buf := make([]byte, HEADLENGTH)
					// Read the incoming connection into the buffer.
					_, err := conn.Read(buf)
					if err != nil {
						fmt.Errorf("Error reading request head:%v", err.Error())
					}
					length := binary.BigEndian.Uint32(buf)
					_buf := make([]byte, length)
					_, err = conn.Read(_buf)

					if err != nil {
						fmt.Errorf("Error reading request body:%v", err.Error())
					}
					var received message.Message
					received.Unmarshal(_buf)
					req <- received
					// Send a response back to person contacting us.
					//conn.Write([]byte("Message received."))
				}(c)
			case <-time.After(1000.0 * time.Second):
				fmt.Println("accepting time out..")
			}
		}
	}()
	return req, nil
}

func reqToBytes(req message.Message) ([]byte, []byte, error) {
	reqBodyBytes, err := req.Marshal(nil)
	if err != nil {
		return nil, reqBodyBytes, err
	}
	reqHead := new(bytes.Buffer)
	if err := binary.Write(reqHead, binary.BigEndian, int32(len(reqBodyBytes))); err != nil {
		return nil, reqBodyBytes, err
	}
	return reqHead.Bytes(), reqBodyBytes, nil
}

//NetConfig p2p net config
type NetConifg struct {
	LogPath       string
	NodeTablePath string
	NodeID        string
	ListenAddr    string
}

func (conf *NetConifg) SetLogPath(path string) *NetConifg {
	if path == "" {
		fmt.Errorf("path of log should not be empty")
	}
	conf.LogPath = path
	return conf
}

func (conf *NetConifg) SetNodeTablePath(path string) *NetConifg {
	if path == "" {
		fmt.Errorf("path of node table should not be empty")
	}
	conf.NodeTablePath = path
	return conf
}

func (conf *NetConifg) SetNodeID(id string) *NetConifg {
	if id == "" {
		fmt.Errorf("node id should not be empty")
	}
	conf.NodeID = id
	return conf
}

func (conf *NetConifg) SetListenAddr(addr string) *NetConifg {
	if addr == "" {
		fmt.Errorf("listen addr should not be empty")
	}
	conf.ListenAddr = addr
	return conf
}

//BaseNetwork boot node maintain all node table, and distribute the node table to all node
type BaseNetwork struct {
	nodeTable  *db.LDBDatabase //all known node except remoteAddr
	neighbours map[string]*discover.Node
	lock       sync.RWMutex
	peers      peerSet // manage all connection
	RecvCh     chan message.Message
	listener   net.Listener

	NodeHeightMap map[string]uint64 //maintain all height of nodes higher than current height
	localNode     *discover.Node

	DownloadHeights map[uint64]uint8 //map[height]retry_times
	log             *log.Logger
}

// NewBaseNetwork ...
func NewBaseNetwork(conf *NetConifg) (*BaseNetwork, error) {
	recv := make(chan message.Message, 100)
	var err error
	if conf.LogPath == "" {
		conf.LogPath, err = ioutil.TempDir(os.TempDir(), "iost_log_")
		if err != nil {
			return nil, fmt.Errorf("iost_log_path err: %v", err)
		}
	}
	if conf.NodeTablePath == "" {
		conf.NodeTablePath = "iost_node_table_"
	}
	srvLog, err := log.NewLogger(conf.LogPath)
	if err != nil {
		return nil, fmt.Errorf("failed to init log %v", err)
	}

	nodeTable, err := db.NewLDBDatabase(conf.NodeTablePath, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to init db %v", err)
	}
	neighbours := make(map[string]*discover.Node, 0)
	NodeHeightMap := make(map[string]uint64, 0)
	if conf.NodeID == "" {
		conf.NodeID = string(discover.GenNodeId())
	}
	localNode := &discover.Node{ID: discover.NodeID(conf.NodeID), IP: net.ParseIP(conf.ListenAddr)}
	downloadHeights := make(map[uint64]uint8, 0)
	s := &BaseNetwork{
		nodeTable:       nodeTable,
		RecvCh:          recv,
		localNode:       localNode,
		neighbours:      neighbours,
		log:             srvLog,
		NodeHeightMap:   NodeHeightMap,
		DownloadHeights: downloadHeights,
	}
	return s, nil
}

// Listen listen local port, find neighbours
func (bn *BaseNetwork) Listen(port uint16) (<-chan message.Message, error) {
	bn.localNode.TCP = port
	bn.log.D("[net] listening %v", bn.localNode)
	var err error
	bn.listener, err = net.Listen("tcp", bn.localNode.Addr())
	if err != nil {
		return bn.RecvCh, errors.New("failed to listen addr, err  = " + fmt.Sprintf("%v", err))
	}
	go func() {
		for {
			conn, err := bn.listener.Accept()
			if err != nil {
				bn.log.E("[net] accept downStream node err:%v", err)
				continue
			}
			go bn.receiveLoop(conn)
		}
	}()
	//register
	go bn.registerLoop()
	go bn.nodeCheckLoop()
	return bn.RecvCh, nil
}

//Broadcast msg to all node in the node table
func (bn *BaseNetwork) Broadcast(msg message.Message) {
	neighbours := bn.neighbours
	bn.lock.Lock()
	for _, node := range neighbours {
		bn.log.D("[net] broad msg: type= %v, from=%v,to=%v,time=%v, to node: %v", msg.ReqType, msg.From, msg.To, msg.Time, node.String())
		msg.To = node.String()
		go bn.broadcast(msg)
	}
	bn.lock.Unlock()
}

//broadcast broadcast to all neighbours, stop broadcast when msg already broadcast
func (bn *BaseNetwork) broadcast(msg message.Message) {
	if msg.TTL == 0 {
		return
	} else {
		msg.TTL = msg.TTL - 1
	}
	data, err := msg.Marshal(nil)
	if err != nil {
		bn.log.E("[net] marshal request encountered err:%v", err)
	}
	req := newRequest(BroadcastMessage, bn.localNode.String(), data)
	conn, err := bn.dial(msg.To)
	if err != nil {
		bn.log.E("[net] broadcast dial tcp got err:%v", err)
		return
	}
	if er := bn.send(conn, req); er != nil {
		bn.peers.RemoveByNodeStr(msg.To)
	}
}

func (bn *BaseNetwork) dial(nodeStr string) (net.Conn, error) {
	bn.lock.Lock()
	defer bn.lock.Unlock()
	node, _ := discover.ParseNode(nodeStr)
	peer := bn.peers.Get(node)
	if peer == nil {
		bn.log.D("[net] dial to %v", node.Addr())
		conn, err := net.Dial("tcp", node.Addr())
		if err != nil {
			bn.log.E("[net] dial tcp %v got err:%v", node.Addr(), err)
			return conn, fmt.Errorf("dial tcp %v got err:%v", node.Addr(), err)
		}
		go bn.receiveLoop(conn)
		peer := newPeer(conn, bn.localNode.String(), nodeStr)
		bn.peers.Set(node, peer)
	}

	return bn.peers.Get(node).conn, nil
}

//Send msg to msg.To
func (bn *BaseNetwork) Send(msg message.Message) {
	if msg.TTL == 0 {
		return
	} else {
		msg.TTL = msg.TTL - 1
	}
	data, err := msg.Marshal(nil)
	if err != nil {
		bn.log.E("[net] marshal request encountered err:%v", err)
	}
	bn.log.D("[net] send msg: type= %v, from=%v,to=%v,time=%v, to node: %v", msg.ReqType, msg.From, msg.To, msg.Time)
	req := newRequest(Message, bn.localNode.String(), data)
	conn, err := bn.dial(msg.To)
	if err != nil {
		bn.log.E("[net] Send, dial tcp got err:%v", err)
		return
	}
	if er := bn.send(conn, req); er != nil {
		bn.peers.RemoveByNodeStr(msg.To)
	}
}

// Close all connection
func (bn *BaseNetwork) Close(port uint16) error {
	if bn.listener != nil {
		bn.listener.Close()
	}
	return nil
}
