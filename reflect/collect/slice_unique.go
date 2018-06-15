package collect

import "reflect"

// The SliceUnique method returns all of the unique items in the slice as new slice
func SliceUnique(value interface{}) interface{} {
	refEmptyStruct := reflect.ValueOf(struct{}{})
	refSlice := reflect.ValueOf(value)
	uMap := reflect.MakeMap(
		reflect.MapOf(
			refSlice.Type().Elem(),
			refEmptyStruct.Type(),
		),
	)

	for i := 0; i < refSlice.Len(); i += 1 {
		uMap.SetMapIndex(refSlice.Index(i), refEmptyStruct)
	}

	slice := reflect.New(
		reflect.SliceOf(
			refSlice.Type().Elem(),
		),
	).Elem()

	for _, k := range uMap.MapKeys() {
		slice = reflect.Append(slice, k)
	}

	return slice.Interface()
}
