package network

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
