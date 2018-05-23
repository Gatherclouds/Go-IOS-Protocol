package verifier

import "github.com/ethereum/go-ethereum/core/vm"

type vmHolder struct {
	vm.VM
	contract vm.Contract
}

type vmMonitor struct {
	vms map[string]vmHolder
}

func newVMMonitor() vmMonitor {
	return vmMonitor{
		vms: make(map[string]vmHolder),
	}
}

func (m *vmMonitor) StartVM(contract vm.Contract) vm.VM {
	if _, ok := m.vms[contract.Info().Prefix]; ok {
		return nil
	}

	switch contract.(type) {
	case *lua.Contract:
		var lvm lua.VM
		err := lvm.Prepare(contract.(*lua.Contract), m)
		if err != nil {
			panic(err)
		}
		err = lvm.Start()
		if err != nil {
			panic(err)
		}
		m.vms[contract.Info().Prefix] = vmHolder{&lvm, contract}
		return &lvm
	}
	return nil
}

func (m *vmMonitor) RestartVM(contract vm.Contract) vm.VM {
	if _, ok := m.vms[contract.Info().Prefix]; ok {
		m.StopVM(contract)
	}
	return m.StartVM(contract)
}

func (m *vmMonitor) StopVM(contract vm.Contract) {
	m.vms[contract.Info().Prefix].Stop()
	delete(m.vms, string(contract.Hash()))
}

