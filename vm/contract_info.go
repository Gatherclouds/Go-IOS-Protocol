package vm

type ContractInfo struct {
	Prefix   string
	Language string
	Version  int8

	GasLimit int64
	Price    float64

	Signers []IOSTAccount
	Sender  IOSTAccount
}

func (c *ContractInfo) toRaw() contractInfoRaw {
	return contractInfoRaw{
		Language: c.Language,
		Version:  c.Version,
		GasLimit: c.GasLimit,
		Price:    c.Price,
	}
}

