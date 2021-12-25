package manipulations

import "reflect"

func InterfaceToInterfaceSlice(in interface{}) (out []interface{}) {
	rv := reflect.ValueOf(in)
	if rv.Kind() == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			out = append(out, rv.Index(i).Interface())
		}
	}

	return
}
