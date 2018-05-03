package iostdb

type StatePoolImpl struct {
	cli redis.Conn
}

const (
	Conn   = "tcp"
	DBAddr = "localhost:6379"
)

