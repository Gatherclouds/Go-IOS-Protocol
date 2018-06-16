package transaction

import (
	"fmt"
	"math"

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


func (this *TransferFrom) validate() {
	debug.assert(this.fee.amount >= 0, "fee < 0")
	debug.assert(this.amount.amount > 0, "amount <= 0")
	debug.assert(len(this.input) > 0, "input size <= 0")
	debug.aseert(this.amount.asset_id == this.fee.asset_id, "fee must be payed by asset owner")

	debug.assert(len(in) != 0, "must be at least on input")
}
