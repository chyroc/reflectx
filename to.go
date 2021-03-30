package reflectx

import (
	"fmt"
	"reflect"
)

func ToInt64(v reflect.Value) (int64, error) {
	if !v.IsValid() {
		return 0, fmt.Errorf("reflect value is not valid")
	}

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil
	}
	return 0, fmt.Errorf("value(%v) of kind(%s) can not convert to int64", v.Interface(), v.Kind())
}

func ToBool(rv reflect.Value) bool {
	var v uint64
	if IsUnsignedInt(rv) {
		v = rv.Uint()
	} else {
		v = uint64(rv.Int())
	}

	return v > 0
}
