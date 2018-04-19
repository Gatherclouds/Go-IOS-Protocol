package core

type Request struct {
	Time    int64
	From    string
	To      string
	ReqType int
	Body    []byte
}

//type Response struct {
//	Time        int64
//	From        string
//	To          string
//	Code        int
//	Description string
//}
