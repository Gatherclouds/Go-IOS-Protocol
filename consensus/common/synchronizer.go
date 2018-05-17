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

