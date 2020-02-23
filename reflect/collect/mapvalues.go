package collect

import (
	"fmt"
	"reflect"
)

func MapValues(value Collection, mapFunc Fn) Collection {
	refValue := reflect.ValueOf(value)
	refFunc := reflect.ValueOf(mapFunc)
	switch refValue.Kind() {
	case reflect.Map:
		return mapMapValues(refValue, refFunc)
	case reflect.Slice:
		return sliceMapValues(refValue, refFunc)
	default:
		panic(fmt.Errorf("unsupported kind: %v", refValue.Kind().String()))
	}
}

func FnMapValues(mapFunc Fn) func(Collection) Collection {
	return func(value Collection) Collection {
		return MapValues(value, mapFunc)
	}
}

// mapMapValues func iterates through the map and passes each value to the given mapFunc.
// The mapFunc is free to modify the item and return it, thus forming a new slice of modified items
func mapMapValues(refMap, refFunc reflect.Value) Map {
	refFuncType := refFunc.Type()
	resMap := reflect.MakeMap(
		reflect.MapOf(
			refMap.Type().Key(),
			refFuncType.Out(0),
		),
	)
	for _, k := range refMap.MapKeys() {
		v := fnKVCall(refFunc, refFuncType, k, refMap.MapIndex(k))[0]
		resMap.SetMapIndex(k, v)
	}
	return resMap.Interface()
}

// sliceMapValues func iterates through the slice and passes each value to the given mapFunc.
// The mapFunc is free to modify the item and return it, thus forming a new slice of modified items
func sliceMapValues(refSlice, refFunc reflect.Value) Slice {
	refFuncType := refFunc.Type()
	resSlice := reflect.MakeSlice(
		reflect.SliceOf(
			refFuncType.Out(0),
		),
		refSlice.Len(),
		refSlice.Len(),
	)
	for i := 0; i < refSlice.Len(); i += 1 {
		v := fnKVCall(refFunc, refFuncType, reflect.ValueOf(i), refSlice.Index(i))[0]
		resSlice.Index(i).Set(v)
	}
	return resSlice.Interface()
}
