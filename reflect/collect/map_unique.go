package collect

import "reflect"

// MapUnique method returns all of the unique items in the map as new slice
func MapUnique(value interface{}) interface{} {
	refEmptyStruct := reflect.ValueOf(struct{}{})
	refMap := reflect.ValueOf(value)
	uMap := reflect.MakeMap(
		reflect.MapOf(
			refMap.Type().Elem(),
			refEmptyStruct.Type(),
		),
	)

	for _, k := range refMap.MapKeys() {
		uMap.SetMapIndex(refMap.MapIndex(k), refEmptyStruct)
	}

	slice := reflect.New(
		reflect.SliceOf(
			refMap.Type().Elem(),
		),
	).Elem()

	for _, k := range uMap.MapKeys() {
		slice = reflect.Append(slice, k)
	}

	return slice.Interface()
}
