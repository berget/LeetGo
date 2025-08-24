package addtwonumbers

import (
	"reflect"
	"testing"
)

func makeList(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func listToSlice(l *ListNode) []int {
	var out []int
	for l != nil {
		out = append(out, l.Val)
		l = l.Next
	}
	return out
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		a, b []int
		want []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for _, tc := range tests {
		l1 := makeList(tc.a)
		l2 := makeList(tc.b)
		got := AddTwoNumbers(l1, l2)
		if !reflect.DeepEqual(listToSlice(got), tc.want) {
			t.Fatalf("AddTwoNumbers(%v, %v) = %v, want %v", tc.a, tc.b, listToSlice(got), tc.want)
		}
	}
}
