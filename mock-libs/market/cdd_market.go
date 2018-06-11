package market

import (
	"math"

	"github.com/iost-official/Go-IOS-Protocol/mock-libs/asset"
)

func (this *CddMarket) calc_benefit(context PolicyObj) int {
	delta_time := (context.now - this.last_update).to_second()

	delta_coin := context.balance.amount.value
	delta_coin *= delta_time

	benefit := context.balance.amount.value
	benefit *= math.max(this.vesting_second, 1)

	return math.min(this.coin_second_earned+delta_coin, benefit)
}

func (this *CddMarket) update_benefit(context PolicyObj) {
	this.last_update = context.now
	this.coin_second_earned = this.calc_benefit(context)
}

func (this *CddMarket) get_allowed_withdraw(context PolicyObj) AssetObj {

	if context.now <= this.start_clain {
		return make(AssetObj(0, context.balance.asset_id))
	}

	cs_earned := this.calc_benefit(context)
	withdraw_available := cs_earned / math.max(this.vesting_second, 1)

	return make(AssetObj(withdraw_available, context.balance.asset_id))
}

func (this *CddMarket) is_withdraw_allowed(context PlicyObj) bool {
	return context.amount <= this.get_allowed_withdraw(context)
}

func (this *CddMarket) on_withdraw(context PolicyObj) {
	this.update_benefit(context)
	coin_seconds_needed := context.amount.amount.value
	coin_seconds_needed *= math.max(this.vesting_second, 1)

	this.coin_second_earned -= coin_seconds_needed
}
