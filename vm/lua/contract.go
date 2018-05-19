package lua

type Contract struct {
	info vm.ContractInfo
	code string
	main Method
	apis map[string]Method
}

func (c *Contract) Info() vm.ContractInfo {
	return c.info
}
