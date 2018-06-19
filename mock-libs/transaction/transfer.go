package transaction

import (
	"crypto"
	"github.com/iost-official/Go-IOS-Protocol/iosbase/debug"
)

func (this *TransferTo) validate() {

	debug.assert(this.fee.amount >= 0, "fee < 0")
	debug.assert(this.amount.amount > 0, "amount <= 0")

	var in, out []CommitmentType
	net_public := amount.amount.value
	public_c := make(crypto.blind(this.bf, net_public))

}
