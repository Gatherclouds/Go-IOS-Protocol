package transaction

import (
	"github.com/iost-official/Go-IOS-Protocol/mock-libs/block"
)

func (this *SignedTransaction) get_required_signatures(chain_id ChainIdType, available_keys map[PublicKeyType]bool) map[PublicKeyType]bool {

	var required_active, required_owner map[AccountIdType]bool
	var other []authority
	this.get_required_authorities(required_active, required_owner, other)

	for active := range required_active {
		s.check_authority(active)
	}

	return result
}
