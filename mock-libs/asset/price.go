package asset

func (this *PriceObj) is_null() bool {
	return this == nil
}

func (this *PriceObj) is_for(asset_id AssetIdType) bool {
	if !this.settlement_price.is_null() {
		return settlement_price.base.asset_id == asset_id
	} else if !this.core_exange_rate.is_null() {
		return core_exange_rate.base.asset_id == asset_id
	}
	return true
}
