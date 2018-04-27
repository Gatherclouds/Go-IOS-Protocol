package p2p

import (
	"github.com/iost-official/prototype/common/mclock"
	"github.com/iost-official/prototype/event"
	"github.com/iost-official/prototype/p2p/discover"
	"io"
	"log"
	"os"
	"sort"
	"sync"
	"time"
)

const (
	baseProtocolVersion    = 5
	baseProtocolLength     = uint64(16)
	baseProtocolMaxMsgSize = 2 * 1024
	snappyProtocolVersion  = 5
	pingInterval           = 15 * time.Second
)

