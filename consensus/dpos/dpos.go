package dpos

type DPoS struct {
	account    Account
	blockCache BlockCache
	router     Router
	synchronizer	Synchronizer
	globalStaticProperty
	globalDynamicProperty

	//测试用，保存投票状态，以及投票消息内容的缓存
	votedStats map[string][]string
	infoCache  [][]byte

	exitSignal chan bool
	chTx       chan message.Message
	chBlock    chan message.Message
}

