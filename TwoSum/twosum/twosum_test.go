package twosum

import "testing"

func eq(a [2]int, b [2]int) bool {
	return (a[0] == b[0] && a[1] == b[1]) || (a[0] == b[1] && a[1] == b[0])
}

func TestTwoSumExamples(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   [2]int
	}{
		{[]int{2, 7, 11, 15}, 9, [2]int{0, 1}},
		{[]int{3, 2, 4}, 6, [2]int{1, 2}},
		{[]int{3, 3}, 6, [2]int{0, 1}},
	}

	for _, tt := range tests {
		got := TwoSum(tt.nums, tt.target)
		if !eq(got, tt.want) {
			t.Fatalf("TwoSum(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}
