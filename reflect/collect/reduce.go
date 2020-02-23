package collect

import (
	"fmt"
	"reflect"
)

func Reduce(value Collection, initialValue Some, reduceFunc Fn) Some {
	refValue := reflect.ValueOf(value)
	refInitialValue := reflect.ValueOf(initialValue)
	refFunc := reflect.ValueOf(reduceFunc)

	switch refValue.Kind() {
	case reflect.Map:
		return mapReduce(refValue, refInitialValue, refFunc)
	case reflect.Slice:
		return sliceReduce(refValue, refInitialValue, refFunc)
	default:
		panic(fmt.Errorf("unsupported kind: %v", refValue.Kind().String()))
	}
}

func FnReduce(initialValue Some, reduceFunc Fn) func(Collection) Some {
	return func(value Collection) Some {
		return Reduce(value, initialValue, reduceFunc)
	}
}

// mapReduce method reduces the map to a single value, passing the result of each iteration
// into the subsequent iteration:
func mapReduce(refMap, refInitialValue, refFunc reflect.Value) Some {
	refValue := reflect.New(refInitialValue.Type())
	refValue.Elem().Set(refInitialValue)
	for _, k := range refMap.MapKeys() {
		v := refFunc.Call(
			[]reflect.Value{refValue.Elem(), k, refMap.MapIndex(k)},
		)[0]

		refValue.Elem().Set(v)
	}
	return refValue.Elem().Interface()
}

// sliceReduce method reduces the slice to a single value, passing the result of each iteration
// into the subsequent iteration:
func sliceReduce(refSlice, refInitialValue, refFunc reflect.Value) interface{} {
	refValue := reflect.New(refInitialValue.Type())
	refValue.Elem().Set(refInitialValue)
	for i := 0; i < refSlice.Len(); i += 1 {
		v := refFunc.Call(
			[]reflect.Value{refValue.Elem(), reflect.ValueOf(i), refSlice.Index(i)},
		)[0]
		refValue.Elem().Set(v)
	}
	return refValue.Elem().Interface()
}
