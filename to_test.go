package reflectx

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/chyroc/go-ptr"
)

func TestToInt64(t *testing.T) {
	tests := []struct {
		name    string
		args    reflect.Value
		want    int64
		wantErr bool
		errReg  *regexp.Regexp
	}{
		{
			name:    "invalid",
			args:    reflect.Value{},
			wantErr: true,
			errReg:  regexp.MustCompile(`reflect value is not valid`),
		},
		{
			name: "int-1",
			args: reflect.ValueOf(1),
			want: 1,
		},
		{
			name: "int-0",
			args: reflect.ValueOf(0),
			want: 0,
		},
		{
			name: "int64-1",
			args: reflect.ValueOf(int64(1)),
			want: 1,
		},
		{
			name: "int64-0",
			args: reflect.ValueOf(int64(0)),
			want: 0,
		},
		{
			name: "uint32-1",
			args: reflect.ValueOf(uint32(1)),
			want: 1,
		},
		{
			name: "uint32-0",
			args: reflect.ValueOf(uint32(0)),
			want: 0,
		},
		{
			name:    "uint64-1",
			args:    reflect.ValueOf(uint64(1)),
			wantErr: true,
		},
		{
			name:    "uint64-0",
			args:    reflect.ValueOf(uint64(0)),
			wantErr: true,
		},
		{
			name:    "str-",
			args:    reflect.ValueOf(""),
			wantErr: true,
		},
		{
			name:    "str-x",
			args:    reflect.ValueOf("x"),
			wantErr: true,
		},
		{
			name:    "float-1.0",
			args:    reflect.ValueOf(1.0),
			wantErr: true,
		},
		{
			name:    "struct",
			args:    reflect.ValueOf(stru{}),
			wantErr: true,
		},
		{
			name:    "struct-ptr",
			args:    reflect.ValueOf(&stru{}),
			wantErr: true,
		},
		{
			name:    "bool-true",
			args:    reflect.ValueOf(true),
			wantErr: true,
		},
		{
			name:    "bool-false",
			args:    reflect.ValueOf(false),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt64(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("convert error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errReg != nil && !tt.errReg.MatchString(err.Error()) {
				t.Errorf("convert error want match %s, but got %s", tt.errReg, err)
				return
			}
			if got != tt.want {
				t.Errorf("convert got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToPtr(t *testing.T) {
	s := &stru{}
	tests := []struct {
		name string
		args reflect.Value
		want reflect.Value
	}{
		{
			name: "int-1",
			args: reflect.ValueOf(1),
			want: reflect.ValueOf(ptr.Int(1)),
		},
		{
			name: "int-0",
			args: reflect.ValueOf(0),
			want: reflect.ValueOf(ptr.Int(0)),
		},
		{
			name: "int64-1",
			args: reflect.ValueOf(int64(1)),
			want: reflect.ValueOf(ptr.Int64(1)),
		},
		{
			name: "int64-0",
			args: reflect.ValueOf(int64(0)),
			want: reflect.ValueOf(ptr.Int64(0)),
		},
		{
			name: "uint32-1",
			args: reflect.ValueOf(uint32(1)),
			want: reflect.ValueOf(ptr.UInt32(1)),
		},
		{
			name: "uint32-0",
			args: reflect.ValueOf(uint32(0)),
			want: reflect.ValueOf(ptr.UInt32(0)),
		},
		{
			name: "uint64-1",
			args: reflect.ValueOf(uint64(1)),
			want: reflect.ValueOf(ptr.UInt64(1)),
		},
		{
			name: "uint64-0",
			args: reflect.ValueOf(uint64(0)),
			want: reflect.ValueOf(ptr.UInt64(0)),
		},
		{
			name: "str-",
			args: reflect.ValueOf(""),
			want: reflect.ValueOf(ptr.String("")),
		},
		{
			name: "str-x",
			args: reflect.ValueOf("x"),
			want: reflect.ValueOf(ptr.String("x")),
		},
		{
			name: "float-1.0",
			args: reflect.ValueOf(float32(1.0)),
			want: reflect.ValueOf(ptr.Float32(1.0)),
		},
		{
			name: "struct",
			args: reflect.ValueOf(stru{}),
			want: reflect.ValueOf(&stru{}),
		},
		{
			name: "struct-ptr",
			args: reflect.ValueOf(&stru{}),
			want: reflect.ValueOf(&s),
		},
		{
			name: "bool-true",
			args: reflect.ValueOf(true),
			want: reflect.ValueOf(ptr.Bool(true)),
		},
		{
			name: "bool-false",
			args: reflect.ValueOf(false),
			want: reflect.ValueOf(ptr.Bool(false)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPtr(tt.args); !reflect.DeepEqual(got.Interface(), tt.want.Interface()) {
				t.Errorf("ToPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBool(t *testing.T) {
	tests := []struct {
		name    string
		args    reflect.Value
		want    bool
		wantErr bool
	}{
		{
			name:    "int-1",
			args:    reflect.ValueOf(1),
			wantErr: true,
		},
		{
			name:    "int-0",
			args:    reflect.ValueOf(0),
			wantErr: true,
		},
		{
			name:    "int64-1",
			args:    reflect.ValueOf(int64(1)),
			wantErr: true,
		},
		{
			name:    "int64-0",
			args:    reflect.ValueOf(int64(0)),
			wantErr: true,
		},
		{
			name:    "uint32-1",
			args:    reflect.ValueOf(uint32(1)),
			wantErr: true,
		},
		{
			name:    "uint32-0",
			args:    reflect.ValueOf(uint32(0)),
			wantErr: true,
		},
		{
			name:    "uint64-1",
			args:    reflect.ValueOf(uint64(1)),
			wantErr: true,
		},
		{
			name:    "uint64-0",
			args:    reflect.ValueOf(uint64(0)),
			wantErr: true,
		},
		{
			name:    "str-",
			args:    reflect.ValueOf(""),
			wantErr: true,
		},
		{
			name:    "str-x",
			args:    reflect.ValueOf("x"),
			wantErr: true,
		},
		{
			name:    "float-1.0",
			args:    reflect.ValueOf(1.0),
			wantErr: true,
		},
		{
			name:    "struct",
			args:    reflect.ValueOf(stru{}),
			wantErr: true,
		},
		{
			name:    "struct-ptr",
			args:    reflect.ValueOf(&stru{}),
			wantErr: true,
		},
		{
			name: "bool-true",
			args: reflect.ValueOf(true),
			want: true,
		},
		{
			name: "bool-false",
			args: reflect.ValueOf(false),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBool(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}
