package iostdb

// Block chain

type BlockChain interface {
	Get(layer int) (*Block, error)
	Push(block *Block) error
	Length() int
	Top() *Block
	Init() error
	Close() error
}


