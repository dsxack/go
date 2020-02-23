package collect

import (
	"reflect"
)

func Compose(f1, f2 Fn, funcs ...Fn) func(...Some) Some {
	return func(values ...Some) Some {
		refValues := MapValues(values, func(_, v Some) reflect.Value {
			return reflect.ValueOf(v)
		}).([]reflect.Value)

		funcs = append([]Fn{f1, f2}, funcs...)
		for i := len(funcs) - 1; i >= 0; i-- {
			f := funcs[i]

			refFunc := reflect.ValueOf(f)

			if !refFunc.Type().IsVariadic() {
				for i, v := range refValues {
					refValues[i] = v.Convert(refFunc.Type().In(i))
				}
			}

			refValues = refFunc.Call(refValues)

			for i, v := range refValues {
				refValues[i] = reflect.ValueOf(v.Interface())
			}
		}
		return refValues[0].Interface()
	}
}
