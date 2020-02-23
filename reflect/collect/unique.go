package collect

import (
	"fmt"
	"reflect"
)

func Unique(value Collection) Slice {
	refValue := reflect.ValueOf(value)

	switch refValue.Kind() {
	case reflect.Map:
		return mapUnique(refValue)
	case reflect.Slice:
		return sliceUnique(refValue)
	default:
		panic(fmt.Errorf("unsupported kind: %v", refValue.Kind().String()))
	}
}

// mapUnique method returns all of the unique items in the map as new slice
func mapUnique(refMap reflect.Value) Slice {
	refEmptyStruct := reflect.ValueOf(struct{}{})
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

// sliceUnique method returns all of the unique items in the slice as new slice
func sliceUnique(refSlice reflect.Value) Slice {
	refEmptyStruct := reflect.ValueOf(struct{}{})
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
