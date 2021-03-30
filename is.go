package reflectx

import (
	"reflect"
)

func IsBool(rv reflect.Value) bool {
	k := rv.Kind()
	return k == reflect.Bool
}

func IsSignedInt(rv reflect.Value) bool {
	k := rv.Kind()
	return k == reflect.Int ||
		k == reflect.Int8 ||
		k == reflect.Int16 ||
		k == reflect.Int32 ||
		k == reflect.Int64
}

func IsUnsignedInt(rv reflect.Value) bool {
	k := rv.Kind()
	return k == reflect.Uint ||
		k == reflect.Uint8 ||
		k == reflect.Uint16 ||
		k == reflect.Uint32 ||
		k == reflect.Uint64
}

func IsInt(rv reflect.Value) bool {
	return IsSignedInt(rv) || IsUnsignedInt(rv)
}
