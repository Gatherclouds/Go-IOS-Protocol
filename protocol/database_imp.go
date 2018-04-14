package protocol

import "fmt"

func DatabaseFactory(target string, chain iosbase.BlockChain, pool iosbase.StatePool) (Database, error) {
	switch target {
	case "base":
		return &DatabaseImpl{
			bc:         chain,
			sp:         pool,
			chViewList: []chan View{},
		}, nil
	}
	return nil, fmt.Errorf("target Database not found")
}

type DatabaseImpl struct {
	bc   iosbase.BlockChain
	sp   iosbase.StatePool
	view View

	chViewList []chan View
}

func (d *DatabaseImpl) NewViewSignal() (chan View, error) {
	chView := make(chan View)
	d.chViewList = append(d.chViewList, chView)
	return chView, nil
}

func (d *DatabaseImpl) VerifyTx(tx iosbase.Tx) error {
	// here only existence of Tx inputs will be verified
	for _, in := range tx.Inputs {
		if _, err := d.sp.Find(in.StateHash); err != nil {
			return fmt.Errorf("some input not found")
		}
	}
	return nil
}
func (d *DatabaseImpl) VerifyTxWithCache(tx iosbase.Tx, cachePool iosbase.TxPool) error {
	err := d.VerifyTx(tx)
	if err != nil {
		return err
	}
	txs, _ := cachePool.GetSlice()
	for _, existedTx := range txs {
		if iosbase.Equal(existedTx.Hash(), tx.Hash()) {
			return fmt.Errorf("has included")
		}
		if txConflict(existedTx, tx) {
			return fmt.Errorf("conflicted")
		} else if sliceIntersect(existedTx.Inputs, tx.Inputs) {
			return fmt.Errorf("conflicted")
		}
	}
	return nil
}
func (d *DatabaseImpl) VerifyBlock(block *iosbase.Block) error {
	var blkTxPool iosbase.TxPool
	blkTxPool.Decode(block.Content)

	txs, _ := blkTxPool.GetSlice()
	for i, tx := range txs {
		if i == 0 { // verify coinbase tx
			continue
		}
		err := d.VerifyTx(tx)
		if err != nil {
			return err
		}
	}
	return nil
}
func (d *DatabaseImpl) VerifyBlockWithCache(block *iosbase.Block, cachePool iosbase.TxPool) error {
	var blkTxPool iosbase.TxPool
	blkTxPool.Decode(block.Content)

	txs, _ := blkTxPool.GetSlice()
	for i, tx := range txs {
		if i == 0 { // TODO: verify coinbase tx
			continue
		}
		err := d.VerifyTxWithCache(tx, cachePool)
		if err != nil {
			return err
		}
	}
	return nil
}
func (d *DatabaseImpl) PushBlock(block *iosbase.Block) error {
	d.bc.Push(block)
	d.sp.Transact(block)
	var err error
	d.view, err = ViewFactory("dpos")
	if err != nil {
		return err
	}
	d.view.Init(d.bc)

	for _, chv := range d.chViewList {
		chv <- d.view
	}
	return nil
}
func (d *DatabaseImpl) GetStatePool() (iosbase.StatePool, error) {
	return d.sp, nil
}
func (d *DatabaseImpl) GetBlockChain() (iosbase.BlockChain, error) {
	return d.bc, nil
}
func (d *DatabaseImpl) GetCurrentView() (View, error) {
	return d.view, nil
}

func sliceEqualI(a, b []iosbase.TxInput) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !iosbase.Equal(a[i].Hash(), b[i].Hash()) {
			return false
		}
	}
	return true
}

func sliceEqualS(a, b []iosbase.State) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !iosbase.Equal(a[i].Hash(), b[i].Hash()) {
			return false
		}
	}
	return true
}

func sliceIntersect(a []iosbase.TxInput, b []iosbase.TxInput) bool {
	for _, ina := range a {
		for _, inb := range b {
			if iosbase.Equal(ina.Hash(), inb.Hash()) {
				return true
			}
		}
	}
	return false
}

func txConflict(a, b iosbase.Tx) bool {
	if sliceEqualI(a.Inputs, b.Inputs) &&
		sliceEqualS(a.Outputs, b.Outputs) &&
		a.Recorder != b.Recorder {
		return true
	} else {
		return false
	}
}

