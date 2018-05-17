package common

var (
	SyncNumber        = 10
	MaxDownloadNumber = 10
)

type Synchronizer interface {
	StartListen() error
	NeedSync() (bool, uint64, uint64)
	SyncBlocks(startNumber uint64, endNumber uint64) error
}

type SyncImpl struct {
	blockCache   BlockCache
	router       Router
	heightChan   chan message.Message
	blkSyncChain chan message.Message
}

func NewSynchronizer(bc BlockCache, router Router) *SyncImpl {
	sync := &SyncImpl{
		blockCache: bc,
		router:     router,
	}
	var err error
	sync.heightChan, err = sync.router.FilteredChan(Filter{
		WhiteList:  []message.Message{},
		BlackList:  []message.Message{},
		RejectType: []ReqType{},
		AcceptType: []ReqType{
			ReqBlockHeight,
		}})
	if err != nil {
		return nil
	}

	sync.blkSyncChain, err = sync.router.FilteredChan(Filter{
		WhiteList:  []message.Message{},
		BlackList:  []message.Message{},
		RejectType: []ReqType{},
		AcceptType: []ReqType{
			ReqDownloadBlock,
		}})
	if err != nil {
		return nil
	}
	return sync
}

