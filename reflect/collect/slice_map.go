package collect

import "reflect"

func SliceMap(value interface{}, mapFunc interface{}) interface{} {
	refFunc := reflect.ValueOf(mapFunc)
	refSlice := reflect.ValueOf(value)
	resSlice := reflect.MakeSlice(
		reflect.SliceOf(
			refFunc.Type().Out(0),
		),
		refSlice.Len(),
		refSlice.Len(),
	)

	for i := 0; i < refSlice.Len(); i += 1 {
		v := refFunc.Call([]reflect.Value{refSlice.Index(i)})[0]
		resSlice.Index(i).Set(v)
	}

	return resSlice.Interface()
}
