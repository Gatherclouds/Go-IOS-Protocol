package consensus

type TxStatus int

const (
	ACCEPT TxStatus = iota
	CACHED
	POOL
	REJECT
	EXPIRED
)

type Consensus interface {
	Run()
	Stop()

	GetBlockChain() block.Chain
	GetCachedBlockChain() block.Chain
}
