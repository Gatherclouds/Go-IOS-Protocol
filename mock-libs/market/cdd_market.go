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

