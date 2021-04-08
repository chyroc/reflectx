package reflectx

import (
	"reflect"
)

// IsBool check reflect.Value is bool type, example: true and false
func IsBool(rv reflect.Value) bool {
	k := rv.Kind()
	return k == reflect.Bool
}

// IsSignedInt check reflect.Value is singed-int type, example: int(1)
func IsSignedInt(rv reflect.Value) bool {
	k := rv.Kind()
	return k == reflect.Int ||
		k == reflect.Int8 ||
		k == reflect.Int16 ||
		k == reflect.Int32 ||
		k == reflect.Int64
}

// IsUnsignedInt check reflect.Value is unsinged-int type, example: uint(1)
func IsUnsignedInt(rv reflect.Value) bool {
	k := rv.Kind()
	return k == reflect.Uint ||
		k == reflect.Uint8 ||
		k == reflect.Uint16 ||
		k == reflect.Uint32 ||
		k == reflect.Uint64
}

// IsInt check reflect.Value is int type, example: int(1) and uint(1)
func IsInt(rv reflect.Value) bool {
	return IsSignedInt(rv) || IsUnsignedInt(rv)
}
