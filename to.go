package reflectx

import (
	"fmt"
	"reflect"
)

// ToInt64 convert reflect.Value to int64 value, if invalid, return error
func ToInt64(v reflect.Value) (int64, error) {
	if !v.IsValid() {
		return 0, fmt.Errorf("reflect value is not valid")
	}

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return int64(v.Uint()), nil
	}
	return 0, fmt.Errorf("value(%v) of kind(%s) can not convert to int64", v.Interface(), v.Kind())
}

// ToBool convert reflect.Value to bool value, if invalid, return error
func ToBool(v reflect.Value) (bool, error) {
	if IsBool(v) {
		return v.Bool(), nil
	}
	return false, fmt.Errorf("value(%v) of kind(%s) can not convert to bool", v.Interface(), v.Kind())
}

// ToPtr convert reflect.Value to pointer of v value
func ToPtr(v reflect.Value) reflect.Value {
	pt := reflect.PtrTo(v.Type())
	pv := reflect.New(pt.Elem())
	pv.Elem().Set(v)
	return pv
}
