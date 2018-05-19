package lua

func Lua2Core(value lua.LValue) state.Value {
	var v state.Value
	switch value.(type) {
	case lua.LNumber:
		vl := value.(*lua.LNumber)
		v = state.MakeVFloat(float64(*vl))
	case lua.LString:
		v = state.MakeVString(value.String())
	case lua.LBool:
		if value == lua.LTrue {
			v = state.VTrue
		} else {
			v = state.VFalse
		}
	}
	return v
}


