package lua

import (
	"fmt"
	"reflect"
)

func Lua2Core(value lua.LValue) state.Value {
	var v state.Value
	switch value.(type) {
	case *lua.LNilType:
		return state.VNil
	case lua.LNumber:
		vl := value.(lua.LNumber)
		v = state.MakeVFloat(float64(vl))
		return v
	case lua.LString:
		v = state.MakeVString(value.String())
		return v
	case lua.LBool:
		if value == lua.LTrue {
			v = state.VTrue
		} else {
			v = state.VFalse
		}
		return v
	}
	panic(fmt.Errorf("not support convertion: %v", reflect.TypeOf(value).String()))

}