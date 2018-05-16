package iostdb

import (
	"github.com/golang/mock/gomock"
)

// NewMockStatePool creates a new mock instance
func NewMockStatePool(ctrl *gomock.Controller) *MockStatePool {
	mock := &MockStatePool{ctrl: ctrl}
	mock.recorder = &MockStatePoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStatePool) EXPECT() *MockStatePoolMockRecorder {
	return m.recorder
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












