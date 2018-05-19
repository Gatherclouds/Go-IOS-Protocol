package tx

import "github.com/ethereum/go-ethereum/core/vm"

type Tx struct {
	Time      int64
	Contract  vm.Contract
	Signs     []common.Signature
	Publisher []common.Signature
}

func (t *Tx) Encode() []byte {
	return nil
}
func (t *Tx) Decode([]byte) error {
	return nil
}
func (t *Tx) Hash() []byte {
	return nil
}