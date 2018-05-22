package verifier

import "github.com/ethereum/go-ethereum/core/vm"

type vmHolder struct {
	vm.VM
	contract vm.Contract
}

type vmMonitor struct {
	vms map[string]vmHolder
}

