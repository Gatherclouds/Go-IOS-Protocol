package block

import (
	"crypto"
)

func (this *BlockHeader) b_id() BlockIdType {
	this_hash := crypto.sha224.hash(this)
	this_hash.hash[0] = crypto.reverse(this.block_num)
	var result BlockIdType
	result.hash = this_hash.hash
	return result
}

	var ids []DigestType

	return make(ChecksumType(crypto.make_hash(ids[0])))
}
