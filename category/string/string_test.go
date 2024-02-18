package string

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "example1", args: args{[]byte{'h', 'e', 'l', 'l', 'o'}}, want: []byte{'o', 'l', 'l', 'e', 'h'}},
		{name: "example2", args: args{[]byte{'e', 'l', 'l', 'o'}}, want: []byte{'o', 'l', 'l', 'e'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if reverseString(tt.args.s); !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("fourSum() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}
