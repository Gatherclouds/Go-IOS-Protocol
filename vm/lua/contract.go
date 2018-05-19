package lua

import "fmt"

type Contract struct {
	info vm.ContractInfo
	code string
	main Method
	apis map[string]Method
}

func (c *Contract) Info() vm.ContractInfo {
	return c.info
}

func (c *Contract) SetPrefix(prefix string) {
	c.info.Prefix = prefix
}
func (c *Contract) SetSender(sender vm.IOSTAccount) {
	c.info.Sender = sender
}
func (c *Contract) AddSigner(signer vm.IOSTAccount) {
	c.info.Signers = append(c.info.Signers, signer)
}

func (c *Contract) Api(apiName string) (vm.Method, error) {
	if apiName == "main" {
		return &c.main, nil
	}
	rtn, ok := c.apis[apiName]
	if !ok {
		return nil, fmt.Errorf("api %v : not found", apiName)
	}
	return &rtn, nil
}
func (c *Contract) Encode() []byte {
	cr := contractRaw{
		info: c.info.Encode(),
		code: []byte(c.code),
	}
	b, err := cr.Marshal(nil)
	if err != nil {
		panic(err)
		return nil
	}
	return b
}