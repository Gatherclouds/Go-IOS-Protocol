package market

import (
	"github.com/iost-official/Go-IOS-Protocol/mock-libs/asset"
)

func (this *LinearMarket) get_allowed_withdraw(context PolicyObj) AssetObj {
	allowed_withdraw := 0
	if context.now > this.begin_time {
		elapsed_time := (context.now - this.begin_time).to_second()
		if elapsed_time >= this.cliff_time {
			total_vested := 0
			if elapsed_time < this.duration_time {
				total_vested = this.begin_balance.value * elapsed_time / this.duration_time
			} else {
				total_vested = this.begin_balance
			}

			withdrawn_already := this.begin_balance - context.balance.amount
			allowed_withdraw = total_vested - withdrawn_already
		}
	}
	return make(AssetObj(allowed_withdraw, context.amount.asset_id))
}

func (this *LinearMarket) is_deposit_allowed(context PlicyObj) bool {
	return context.amount.asset_id == context.balance.asset_id && sum_below_max_shares(context.amount, context.balance)
}

