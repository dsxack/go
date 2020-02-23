package collect

import "reflect"

func Concat(slices ...Some) Slice {
	refArgs := reflect.ValueOf(slices)
	if refArgs.Len() == 1 {
		refArgs = refArgs.Index(0).Elem()
	}

	refSlices := make([]reflect.Value, refArgs.Len())
	for i := 0; i < refArgs.Len(); i++ {
		refSlices[i] = indirectInterface(refArgs.Index(i))
	}

	var refResultSlice reflect.Value
	if sameTypes(refSlices...) {
		refResultSlice = reflect.MakeSlice(
			refSlices[0].Type(),
			0,
			refSlices[0].Cap(),
		)
	} else {
		var is []interface{}
		refResultSlice = reflect.MakeSlice(
			reflect.TypeOf(is),
			0,
			refSlices[0].Len(),
		)
	}

	for _, refSlice := range refSlices {
		values := make([]reflect.Value, refSlice.Len())
		for i := 0; i < refSlice.Len(); i++ {
			values[i] = refSlice.Index(i)
		}
		refResultSlice = reflect.Append(refResultSlice, values...)
	}

	return refResultSlice.Interface()
}
