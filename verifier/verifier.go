package verifier

import "github.com/ethereum/go-ethereum/core/vm"

const (
	MaxBlockGas uint64 = 1000000
)

//go:generate gencode go -schema=structs.schema -package=verifier

// 底层verifier，用来组织vm，不要直接使用
type Verifier struct {
	Pool state.Pool
	vmMonitor
}

func (v *Verifier) Verify(contract vm.Contract) (state.Pool, uint64, error) {
	_, pool, gas, err := v.Call(v.Pool, contract.Info().Prefix, "main")
	return pool, gas, err
}

func (v *Verifier) SetPool(pool state.Pool) {
	v.Pool = pool
}

// 验证新tx的工具类
type CacheVerifier struct {
	Verifier
}
