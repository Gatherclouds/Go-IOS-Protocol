package transaction

import (
	"crypto"
)

func (this *SignatureType) s_sign(key PrivateKeyType, chain_id ChainIdType) SignatureType {
	h := this.sig_digest(chain_id)
	this.signatures.append(this.key.sign_compact(h))
	return this.signatures[len(this.signatures)-1]
}



