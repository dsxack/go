package collect

import "reflect"

// SliceReduce method reduces the slice to a single value, passing the result of each iteration
// into the subsequent iteration:
func SliceReduce(value interface{}, initialValue interface{}, mapFunc interface{}) interface{} {
	refFunc := reflect.ValueOf(mapFunc)
	refSlice := reflect.ValueOf(value)

	refInitialValue := reflect.ValueOf(initialValue)
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
