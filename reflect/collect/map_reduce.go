package collect

import "reflect"

// MapReduce method reduces the map to a single value, passing the result of each iteration
// into the subsequent iteration:
func MapReduce(mapValue interface{}, initialValue interface{}, reduceFunc interface{}) interface{} {
	refFunc := reflect.ValueOf(reduceFunc)
	refMap := reflect.ValueOf(mapValue)

	refInitialValue := reflect.ValueOf(initialValue)
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
