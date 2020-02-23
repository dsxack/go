package collect

import (
	"fmt"
	"reflect"
)

func MapWithKeys(value Collection, mapFunc Fn) Map {
	refValue := reflect.ValueOf(value)
	refFunc := reflect.ValueOf(mapFunc)
	switch refValue.Kind() {
	case reflect.Map:
		return mapMapWithKeys(refValue, refFunc)
	case reflect.Slice:
		return sliceMapWithKeys(refValue, refFunc)
	default:
		panic(fmt.Errorf("unsupported kind: %v", refValue.Kind().String()))
	}
}

func FnMapWithKeys(mapFunc Fn) func(Collection) Map {
	return func(value Collection) Map {
		return MapWithKeys(value, mapFunc)
	}
}

// mapMapWithKeys func iterates through the map and passes each key and value to the given mapFunc.
// The mapFunc should return a single key / value pair
func mapMapWithKeys(refMap, refFunc reflect.Value) Map {
	resMap := reflect.MakeMap(
		reflect.MapOf(
			refFunc.Type().Out(0),
			refFunc.Type().Out(1),
		),
	)
	for _, key := range refMap.MapKeys() {
		value := refMap.MapIndex(key)
		refResults := refFunc.Call([]reflect.Value{key, value})
		resMap.SetMapIndex(refResults[0], refResults[1])
	}
	return resMap.Interface()
}

// SliceMapWithKeys func iterates through the slice and passes each key and value to the given mapFunc.
// The mapFunc should return a single key / value pair
func sliceMapWithKeys(refSlice, refFunc reflect.Value) Map {
	resMap := reflect.MakeMap(
		reflect.MapOf(
			refFunc.Type().Out(0),
			refFunc.Type().Out(1),
		),
	)
	for i := 0; i < refSlice.Len(); i += 1 {
		value := refSlice.Index(i)
		refResults := refFunc.Call([]reflect.Value{reflect.ValueOf(i), value})
		resMap.SetMapIndex(refResults[0], refResults[1])
	}
	return resMap.Interface()
}
