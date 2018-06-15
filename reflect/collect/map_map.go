package collect

import "reflect"

// The MapMap func iterates through the map and passes each value to the given mapFunc.
// The mapFunc is free to modify the item and return it, thus forming a new slice of modified items
func MapMap(value interface{}, mapFunc interface{}) interface{} {
	refFunc := reflect.ValueOf(mapFunc)
	refMap := reflect.ValueOf(value)
	resMap := reflect.MakeMap(
		reflect.MapOf(
			refMap.Type().Key(),
			refFunc.Type().Out(0),
		),
	)

	for _, k := range refMap.MapKeys() {
		v := refFunc.Call([]reflect.Value{refMap.MapIndex(k)})[0]
		resMap.SetMapIndex(k, v)
	}

	return resMap.Interface()
}
