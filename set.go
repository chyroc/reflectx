package reflectx

import (
	"fmt"
	"reflect"
)

func SetInt(rv reflect.Value, n int64) error {
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("%s unsupport set int", rv.Kind())
	}
	rv = rv.Elem()
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if rv.OverflowInt(n) {
			return fmt.Errorf("%d overflow for %s", n, rv.Kind())
		}
		rv.SetInt(n)
	default:
		return fmt.Errorf("ptr of %s unsupport set int", rv.Kind())
	}
	return nil
}

func SetUint(rv reflect.Value, n uint64) error {
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("%s unsupport set uint", rv.Kind())
	}
	rv = rv.Elem()
	switch rv.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if rv.OverflowUint(n) {
			return fmt.Errorf("%d overflow for %s", n, rv.Kind())
		}
		rv.SetUint(n)
	default:
		return fmt.Errorf("ptr of %s unsupport set uint", rv.Kind())
	}
	return nil
}

func SetFloat(rv reflect.Value, n float64) error {
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("%s unsupport set float", rv.Kind())
	}
	rv = rv.Elem()
	switch rv.Kind() {
	case reflect.Float32, reflect.Float64:
		if rv.OverflowFloat(n) {
			return fmt.Errorf("%.2f overflow for %s", n, rv.Kind())
		}
		rv.SetFloat(n)
	default:
		return fmt.Errorf("ptr of %s unsupport set float", rv.Kind())
	}
	return nil
}
