package block

type SignatureState struct {
	available_address_sigs map[Address]PublicKeyType
	provided_address_sigs  map[Address]PublicKeyType

	available_keys []PublicKeyType

	provided_signatures map[PublicKeyType]bool

	approved_by   map[AccountIdType]bool
	max_recursion int
}


	return total_weight >= auth.weight_threshold
}

