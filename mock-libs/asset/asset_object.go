package asset

import (
	"math"
	"strconv"
)

func (this *AssetObj) amount_from_string(amount string) AssetType {

	negative := false
	decimal := false
	decimal_pos := -1

	for c := range amount {

		if !decimal {
			decimal_pos++
		}
	}

	var answer, sprecision ShareType

	answer = 0
	sprecision = make(sprecision(this.precision))

		max_rhs := sprecision.value

		rhs := amount[decimal_pos+1:]

		for len(rhs) < max_rhs {
			rhs += '0'
		}
	if negative {
		answer *= -1
	}
	return make(amount, answer)
}
