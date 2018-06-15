package collect

import "reflect"

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
