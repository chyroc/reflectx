package reflectx

import (
	"math"
	"reflect"
	"regexp"
	"testing"

	"github.com/chyroc/go-ptr"
)

func Test_SetInt(t *testing.T) {
	tests := []struct {
		name    string
		args    reflect.Value
		n       int64
		wantErr bool
		errReg  *regexp.Regexp
		want    reflect.Value
	}{
		{
			name:    "invalid",
			args:    reflect.Value{},
			n:       1,
			wantErr: true,
			errReg:  regexp.MustCompile(`invalid unsupport set int`),
		},
		{
			name:    "int-1",
			args:    reflect.ValueOf(int(0)),
			n:       1,
			wantErr: true,
			errReg:  regexp.MustCompile(`int unsupport set int`),
		},
		{
			name: "int-1",
			args: reflect.ValueOf(ptr.Int(int(0))),
			n:    1,
			want: reflect.ValueOf(int(1)),
		},
		{
			name:    "overflow-int32",
			args:    reflect.ValueOf(ptr.Int32(int32(0))),
			n:       math.MaxInt64,
			wantErr: true,
			errReg:  regexp.MustCompile(`9223372036854775807 overflow for int32`),
		},
		{
			name:    "unsupport-type-str",
			args:    reflect.ValueOf(ptr.String("")),
			n:       1,
			wantErr: true,
			errReg:  regexp.MustCompile(`ptr of string unsupport set int`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetInt(tt.args, tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("set, error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errReg != nil && !tt.errReg.MatchString(err.Error()) {
				t.Errorf("set, error want match %s, but got %s", tt.errReg, err)
				return
			}
			if tt.wantErr {
				return
			}
			if tt.args.Elem().Int() != tt.want.Int() {
				t.Errorf("set, got = %v, want %v", tt.args.Elem().Int(), tt.want.Int())
			}
		})
	}
}

func Test_SetUint(t *testing.T) {
	tests := []struct {
		name    string
		args    reflect.Value
		n       uint64
		wantErr bool
		errReg  *regexp.Regexp
		want    reflect.Value
	}{
		{
			name:    "invalid",
			args:    reflect.Value{},
			n:       1,
			wantErr: true,
			errReg:  regexp.MustCompile(`invalid unsupport set uint`),
		},
		{
			name:    "uint-1",
			args:    reflect.ValueOf(uint(0)),
			n:       1,
			wantErr: true,
			errReg:  regexp.MustCompile(`int unsupport set uint`),
		},
		{
			name: "uint-1",
			args: reflect.ValueOf(ptr.UInt(uint(0))),
			n:    1,
			want: reflect.ValueOf(uint(1)),
		},
		{
			name:    "overflow-uint32",
			args:    reflect.ValueOf(ptr.UInt32(uint32(0))),
			n:       math.MaxUint64,
			wantErr: true,
			errReg:  regexp.MustCompile(`18446744073709551615 overflow for uint32`),
		},
		{
			name:    "unsupport-type-str",
			args:    reflect.ValueOf(ptr.String("")),
			n:       1,
			wantErr: true,
			errReg:  regexp.MustCompile(`ptr of string unsupport set uint`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetUint(tt.args, tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("set, error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errReg != nil && !tt.errReg.MatchString(err.Error()) {
				t.Errorf("set, error want match %s, but got %s", tt.errReg, err)
				return
			}
			if tt.wantErr {
				return
			}
			if tt.args.Elem().Uint() != tt.want.Uint() {
				t.Errorf("set, got = %v, want %v", tt.args.Elem().Uint(), tt.want.Uint())
			}
		})
	}
}

func Test_SetFloat64(t *testing.T) {
	tests := []struct {
		name    string
		args    reflect.Value
		n       float64
		wantErr bool
		errReg  *regexp.Regexp
		want    reflect.Value
	}{
		{
			name:    "invalid",
			args:    reflect.Value{},
			n:       1,
			wantErr: true,
			errReg:  regexp.MustCompile(`invalid unsupport set float`),
		},
		{
			name:    "float64-1",
			args:    reflect.ValueOf(float64(0)),
			n:       1.1,
			wantErr: true,
			errReg:  regexp.MustCompile(`float64 unsupport set float`),
		},
		{
			name: "float64-1",
			args: reflect.ValueOf(ptr.Float64(float64(0))),
			n:    1.1,
			want: reflect.ValueOf(float64(1.1)),
		},
		{
			name:    "overflow-float32",
			args:    reflect.ValueOf(ptr.Float32(float32(0))),
			n:       math.MaxFloat64,
			wantErr: true,
			errReg:  regexp.MustCompile(`179769313486231570814527423731704356798070567525844996598917476803157260780028538760589558632766878171540458953514382464234321326889464182768467546703537516986049910576551282076245490090389328944075868508455133942304583236903222948165808559332123348274797826204144723168738177180919299881250404026184124858368.00 overflow for float32`),
		},
		{
			name:    "unsupport-type-str",
			args:    reflect.ValueOf(ptr.String("")),
			n:       1,
			wantErr: true,
			errReg:  regexp.MustCompile(`ptr of string unsupport set float`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetFloat(tt.args, tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("set, error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errReg != nil && !tt.errReg.MatchString(err.Error()) {
				t.Errorf("set, error want match %s, but got %s", tt.errReg, err)
				return
			}
			if tt.wantErr {
				return
			}
			if tt.args.Elem().Float() != tt.want.Float() {
				t.Errorf("set, got = %v, want %v", tt.args.Elem().Float(), tt.want.Float())
			}
		})
	}
}
