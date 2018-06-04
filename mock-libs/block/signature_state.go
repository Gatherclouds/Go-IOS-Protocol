package block

type SignatureState struct {
	available_address_sigs map[Address]PublicKeyType
	provided_address_sigs  map[Address]PublicKeyType

	available_keys []PublicKeyType

	provided_signatures map[PublicKeyType]bool

	approved_by   map[AccountIdType]bool
	max_recursion int
}

func (this *SignatureState) check_authority_by_authority(au *AuthorityType, depth int) bool {
	if au == nil {
		return false
	}

	auth := *au
	total_weight := 0

	for k := range auth.key_auths {
		if signed_by_key(k.key) {
			total_weight += k.value
			if total_weight >= auth.weight_threshold {
				return true
			}
		}
	}

	for k := range auths.address_auths {
		if signed_by_address(k.key) {
			total_weight += k.value
			if total_weight >= auth.weight_threshold {
				return true
			}
		}
	}

	for a := range auth.account_auths {
		approve, exist := this.approved_by[a.key]
		if !exist {
			if depth == this.max_recursion {
				return false
			}
			if check_authority_by_authority(get_active(a.key), depth+1) {
				this.approved_by[a.key] = true
				total_weight += a.value
				if total_weight >= auth.weight_threshold {
					return true
				}
			}
		} else {
			total_weight += a.value
			if total_weight >= auth.weight_threshold {
				return true
			}
		}
	}
	return total_weight >= auth.weight_threshold
}

