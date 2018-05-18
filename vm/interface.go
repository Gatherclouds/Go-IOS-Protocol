/*
Package vm  define vm of smart contract. Use verifier/ to verify txs and blocks
*/
package vm

// Privilege 设定智能合约的接口权限
type Privilege int

type IOSTAccount string

const (
	Private Privilege = iota
	Protected
	Public
)

type Code string

// VM 虚拟机interface，定义了虚拟机的接口
//
// 调用流程为prepare - start - call - stop
type VM interface {
	Prepare(contract Contract, monitor Monitor) error
	Start() error
	Stop()
	Call(pool state.Pool, methodName string, args ...state.Value) ([]state.Value, state.Pool, error)
	PC() uint64
}