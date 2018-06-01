package network

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
