package lua

//go:generate gencode go -schema=structs.schema -package=lua

type api struct {
	name     string
	function func(L *lua.LState) int
}

// VM lua 虚拟机的实现
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

func (l *VM) Call(pool state.Pool, methodName string, args ...state.Value) ([]state.Value, state.Pool, error) {
	if pool != nil {
		l.cachePool = pool.Copy()
	}

	method0, err := l.Contract.API(methodName)
	if err != nil {
		return nil, nil, err
	}

	method := method0.(*Method)
	
	if len(args) == 0 {
    		err = l.L.CallByParam(lua.P{
    			Fn:      l.L.GetGlobal(method.name),
    			NRet:    method.outputCount,
    			Protect: true,
    		})
    	} else {
    		largs := make([]lua.LValue, 0)
    		for _, arg := range args {
    			largs = append(largs, Core2Lua(arg))
    		}
    		err = l.L.CallByParam(lua.P{
    			Fn:      l.L.GetGlobal(method.name),
    			NRet:    method.outputCount,
    			Protect: true,
    		}, largs...)
    	}

	if err != nil {
		return nil, nil, err
	}

	rtnValue := make([]state.Value, 0, method.outputCount)
	for i := 0; i < method.outputCount; i++ {
		ret := l.L.Get(-1) // returned value
		l.L.Pop(1)
		rtnValue = append(rtnValue, Lua2Core(ret))
	}

	return rtnValue, l.cachePool, nil
}