package other

import (
	"reflect"
	"testing"
)

func Test_maximum69Number(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{9669}, 9969},
		{"test2", args{9999}, 9999},
		{"test3", args{9996}, 9999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum69Number(tt.args.num); got != tt.want {
				t.Errorf("maximum69Number() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printVertically(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test1", args{"HOW ARE YOU"}, []string{"HAY", "ORO", "WEU"}},
		{"test2", args{"TO BE OR NOT TO BE"}, []string{"TBONTB", "OEROOE", "   T"}},
		{"test3", args{"CONTEST IS COMING"}, []string{"CIC", "OSO", "N M", "T I", "E N", "S G", "T"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printVertically(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printVertically() = %v, want %v", got, tt.want)
			}
		})
	}
}
