package compiler

type Compiler interface {
	Compile(code string) vm.Contract
}
