package network

import (
	"io"
	"time"
)

type Msg struct {
	Code       uint64
	Size       uint32 // size of the paylod
	Payload    io.Reader
	ReceivedAt time.Time
}
type MsgReader interface {
	ReadMsg() (Msg, error)
}
