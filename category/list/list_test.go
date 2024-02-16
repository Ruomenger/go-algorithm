package list

import (
	"reflect"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "example0", args: args{head: nil, val: 0}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}

	head1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	head2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	head3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	head4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	head1.Next = head2
	head2.Next = head3
	head3.Next = head4
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "example0", args: args{head: head1}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
