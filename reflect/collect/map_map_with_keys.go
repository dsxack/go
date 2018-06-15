package collect

import "reflect"

// MapMapWithKeys func iterates through the map and passes each key and value to the given mapFunc.
// The mapFunc should return a single key / value pair
func MapMapWithKeys(value interface{}, mapFunc interface{}) interface{} {
	refFunc := reflect.ValueOf(mapFunc)
	refMap := reflect.ValueOf(value)
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
