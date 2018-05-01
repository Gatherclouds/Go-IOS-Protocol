package p2p


/*
Marked request types using by protocol
*/
type ReqType int

const (
	ReqPublishTx ReqType = iota
	ReqNewBlock
)

