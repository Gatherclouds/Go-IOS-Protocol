package db

type Request struct {
	Time    int64
	From    string
	To      string
	ReqType int
	Body    []byte
}


