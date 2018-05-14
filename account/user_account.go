package account

type UserAccount struct {
	addr    *Address
	balance *info.UserBalance
	nonce   *info.UserNonce
	// todo: contract_code, storage_db
}

