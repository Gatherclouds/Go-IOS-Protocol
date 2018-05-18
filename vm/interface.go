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

