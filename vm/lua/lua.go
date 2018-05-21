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

