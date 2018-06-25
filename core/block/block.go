package block

import (
	"Go-IOS-Protocol/common"
	"Go-IOS-Protocol/core/tx"
)

//go:generate gencode go -schema=structs.schema -package=block

// Block 是一个区块的结构体定义
type Block struct {
	Head    BlockHead
	Content []tx.Tx
}


// Encode 是区块的序列化方法
func (d *Block) Encode() []byte {
	c := make([][]byte, 0)
	for _, t := range d.Content {
		c = append(c, t.Encode())
	}
	br := BlockRaw{d.Head, c}
	b, err := br.Marshal(nil)
	if err != nil {
		panic(err)
	}
	return b
}

// Decode 是区块的反序列方法
func (d *Block) Decode(bin []byte) error {
	var br BlockRaw
	_, err := br.Unmarshal(bin)
	d.Head = br.Head
	for _, t := range br.Content {
		var tt tx.Tx
		err = tt.Decode(t)
		if err != nil {
			return err
		}
		d.Content = append(d.Content, tt)
	}
	return nil
}

func (d *Block) Hash() []byte {
	return common.Sha256(d.Encode())
}

func (d *Block) HeadHash() []byte {
	return d.Head.Hash()
}

func (d *BlockHead) Encode() []byte {
	bin, err := d.Marshal(nil)
	if err != nil {
		panic(err)
	}
	return bin
}

func (d *BlockHead) Decode(bin []byte) error {
	_, err := d.Unmarshal(bin)
	return err
}
func (d *BlockHead) Hash() []byte {
	return common.Sha256(d.Encode())
}