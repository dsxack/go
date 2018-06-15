package collect

import "reflect"

// The MapKeys func returns all of the map keys as slice
func MapKeys(value interface{}) interface{} {
	refMap := reflect.ValueOf(value)
	resSlice := reflect.MakeSlice(
		reflect.SliceOf(
			refMap.Type().Key(),
		),
		refMap.Len(),
		refMap.Len(),
	)

	for i, k := range refMap.MapKeys() {
		resSlice.Index(i).Set(k)
	}

	return resSlice.Interface()
}
