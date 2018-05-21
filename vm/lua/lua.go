package lua

//go:generate gencode go -schema=structs.schema -package=lua

type api struct {
	name     string
	function func(L *lua.LState) int
}

type VM struct {
	APIs []api
	L    *lua.LState

	cachePool state.Pool

	monitor vm.Monitor

	Contract *Contract

	callerPC uint64
}

func (l *VM) Start() error {
	for _, api := range l.APIs {
		l.L.SetGlobal(api.name, l.L.NewFunction(api.function))
	}

	if err := l.L.DoString(l.Contract.code); err != nil {
		return err
	}

	return nil
}
func (l *VM) Stop() {
	l.L.Close()
}

