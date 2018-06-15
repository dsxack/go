package collect

import "reflect"

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
