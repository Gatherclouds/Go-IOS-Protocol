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

func Core2Lua(value state.Value) lua.LValue {
	var v lua.LValue
	switch value.(type) {
	case *state.VNilType:
		return lua.LNil
	case *state.VFloat:
		vl := value.(*state.VFloat)
		v = lua.LNumber(vl.ToFloat64())
		return v
	case *state.VString:
		v = lua.LString([]rune(value.EncodeString())[1:])
		return v
	case *state.VBool:
		if value == state.VTrue {
			v = lua.LTrue
		} else {
			v = lua.LFalse
		}
		return v
	}
	panic(fmt.Errorf("not support convertion: %v", reflect.TypeOf(value).String()))
}
