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
// NewDPoS: 新建一个DPoS实例
// acc: 节点的Coinbase账户, bc: 基础链(从数据库读取), witnessList: 见证节点列表
func NewDPoS(acc Account, bc block.Chain, witnessList []string /*, network core.Network*/) (*DPoS, error) {
	p := DPoS{}
	p.Account = acc
	p.blockCache = NewBlockCache(bc, len(witnessList)*2/3+1)

	var err error
	p.router, err = RouterFactory("base")
	if err != nil {
		return nil, err
	}

	p.synchronizer = NewSynchronizer(p.blockCache, p.router)
	if p.synchronizer == nil {
		return nil, err
	}

	//	Tx chan init
	p.chTx, err = p.router.FilteredChan(Filter{
		WhiteList:  []message.Message{},
		BlackList:  []message.Message{},
		RejectType: []ReqType{},
		AcceptType: []ReqType{
			ReqPublishTx,
			reqTypeVoteTest, // Only for test
		}})
	if err != nil {
		return nil, err
	}

	//	Block chan init
	p.chBlock, err = p.router.FilteredChan(Filter{
		WhiteList:  []message.Message{},
		BlackList:  []message.Message{},
		RejectType: []ReqType{},
		AcceptType: []ReqType{ReqNewBlock}})
	if err != nil {
		return nil, err
	}

	p.initGlobalProperty(p.Account, witnessList)
	return &p, nil
}

func (p *DPoS) initGlobalProperty(acc Account, witnessList []string) {
	p.globalStaticProperty = newGlobalStaticProperty(acc, witnessList)
	p.globalDynamicProperty = newGlobalDynamicProperty()
}

