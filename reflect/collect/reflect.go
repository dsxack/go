package collect

import "reflect"

func fnKVCall(refFunc reflect.Value, refFuncType reflect.Type, key, value reflect.Value) []reflect.Value {
	switch refFuncType.NumIn() {
	case 1:
		return refFunc.Call([]reflect.Value{value})
	case 2:
		return refFunc.Call([]reflect.Value{key, value})
	default:
		panic("function should have 1 or 2 input arguments")
	}
}

func sameTypes(values ...reflect.Value) bool {
	v := values[0]
	for i := 1; i < len(values); i++ {
		if v.Type() != values[i].Type() {
			return false
		}
	}

	return true
}

func indirectInterface(value reflect.Value) reflect.Value {
	for {
		if value.Kind() != reflect.Interface {
			return value
		}
		value = value.Elem()
	}
}
