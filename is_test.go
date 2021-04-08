package reflectx

import (
	"reflect"
	"testing"
)

type stru struct{}

func TestIsBool(t *testing.T) {
	tests := []struct {
		name string
		args reflect.Value
		want bool
	}{
		{
			name: "int-1",
			args: reflect.ValueOf(1),
			want: false,
		},
		{
			name: "int-0",
			args: reflect.ValueOf(0),
			want: false,
		},
		{
			name: "int--1",
			args: reflect.ValueOf(-1),
			want: false,
		},
		{
			name: "str-",
			args: reflect.ValueOf(""),
			want: false,
		},
		{
			name: "str-x",
			args: reflect.ValueOf("x"),
			want: false,
		},
		{
			name: "float-1.0",
			args: reflect.ValueOf(1.0),
			want: false,
		},
		{
			name: "struct",
			args: reflect.ValueOf(stru{}),
			want: false,
		},
		{
			name: "struct-ptr",
			args: reflect.ValueOf(&stru{}),
			want: false,
		},
		{
			name: "bool-true",
			args: reflect.ValueOf(true),
			want: true,
		},
		{
			name: "bool-false",
			args: reflect.ValueOf(false),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBool(tt.args); got != tt.want {
				t.Errorf("IsBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInt(t *testing.T) {
	tests := []struct {
		name string
		args reflect.Value
		want bool
	}{
		{
			name: "int-1",
			args: reflect.ValueOf(1),
			want: true,
		},
		{
			name: "int-0",
			args: reflect.ValueOf(0),
			want: true,
		},
		{
			name: "int64-1",
			args: reflect.ValueOf(int64(1)),
			want: true,
		},
		{
			name: "int64-0",
			args: reflect.ValueOf(int64(0)),
			want: true,
		},
		{
			name: "uint32-1",
			args: reflect.ValueOf(uint32(1)),
			want: true,
		},
		{
			name: "uint32-0",
			args: reflect.ValueOf(uint32(0)),
			want: true,
		},
		{
			name: "uint64-1",
			args: reflect.ValueOf(uint64(1)),
			want: true,
		},
		{
			name: "uint64-0",
			args: reflect.ValueOf(uint64(0)),
			want: true,
		},
		{
			name: "int--1",
			args: reflect.ValueOf(-1),
			want: true,
		},
		{
			name: "str-",
			args: reflect.ValueOf(""),
			want: false,
		},
		{
			name: "str-x",
			args: reflect.ValueOf("x"),
			want: false,
		},
		{
			name: "float-1.0",
			args: reflect.ValueOf(1.0),
			want: false,
		},
		{
			name: "struct",
			args: reflect.ValueOf(stru{}),
			want: false,
		},
		{
			name: "struct-ptr",
			args: reflect.ValueOf(&stru{}),
			want: false,
		},
		{
			name: "bool-true",
			args: reflect.ValueOf(true),
			want: false,
		},
		{
			name: "bool-false",
			args: reflect.ValueOf(false),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInt(tt.args); got != tt.want {
				t.Errorf("IsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSignedInt(t *testing.T) {
	tests := []struct {
		name string
		args reflect.Value
		want bool
	}{
		{
			name: "int-1",
			args: reflect.ValueOf(1),
			want: true,
		},
		{
			name: "int-0",
			args: reflect.ValueOf(0),
			want: true,
		},
		{
			name: "int64-1",
			args: reflect.ValueOf(int64(1)),
			want: true,
		},
		{
			name: "int64-0",
			args: reflect.ValueOf(int64(0)),
			want: true,
		},
		{
			name: "uint32-1",
			args: reflect.ValueOf(uint32(1)),
			want: false,
		},
		{
			name: "uint32-0",
			args: reflect.ValueOf(uint32(0)),
			want: false,
		},
		{
			name: "uint64-1",
			args: reflect.ValueOf(uint64(1)),
			want: false,
		},
		{
			name: "uint64-0",
			args: reflect.ValueOf(uint64(0)),
			want: false,
		},
		{
			name: "int--1",
			args: reflect.ValueOf(-1),
			want: true,
		},
		{
			name: "str-",
			args: reflect.ValueOf(""),
			want: false,
		},
		{
			name: "str-x",
			args: reflect.ValueOf("x"),
			want: false,
		},
		{
			name: "float-1.0",
			args: reflect.ValueOf(1.0),
			want: false,
		},
		{
			name: "struct",
			args: reflect.ValueOf(stru{}),
			want: false,
		},
		{
			name: "struct-ptr",
			args: reflect.ValueOf(&stru{}),
			want: false,
		},
		{
			name: "bool-true",
			args: reflect.ValueOf(true),
			want: false,
		},
		{
			name: "bool-false",
			args: reflect.ValueOf(false),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSignedInt(tt.args); got != tt.want {
				t.Errorf("IsSignedInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUnsignedInt(t *testing.T) {
	tests := []struct {
		name string
		args reflect.Value
		want bool
	}{
		{
			name: "int-1",
			args: reflect.ValueOf(1),
			want: false,
		},
		{
			name: "int-0",
			args: reflect.ValueOf(0),
			want: false,
		},
		{
			name: "int64-1",
			args: reflect.ValueOf(int64(1)),
			want: false,
		},
		{
			name: "int64-0",
			args: reflect.ValueOf(int64(0)),
			want: false,
		},
		{
			name: "uint32-1",
			args: reflect.ValueOf(uint32(1)),
			want: true,
		},
		{
			name: "uint32-0",
			args: reflect.ValueOf(uint32(0)),
			want: true,
		},
		{
			name: "uint64-1",
			args: reflect.ValueOf(uint64(1)),
			want: true,
		},
		{
			name: "uint64-0",
			args: reflect.ValueOf(uint64(0)),
			want: true,
		},
		{
			name: "int--1",
			args: reflect.ValueOf(-1),
			want: false,
		},
		{
			name: "str-",
			args: reflect.ValueOf(""),
			want: false,
		},
		{
			name: "str-x",
			args: reflect.ValueOf("x"),
			want: false,
		},
		{
			name: "float-1.0",
			args: reflect.ValueOf(1.0),
			want: false,
		},
		{
			name: "struct",
			args: reflect.ValueOf(stru{}),
			want: false,
		},
		{
			name: "struct-ptr",
			args: reflect.ValueOf(&stru{}),
			want: false,
		},
		{
			name: "bool-true",
			args: reflect.ValueOf(true),
			want: false,
		},
		{
			name: "bool-false",
			args: reflect.ValueOf(false),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUnsignedInt(tt.args); got != tt.want {
				t.Errorf("IsUnsignedInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
