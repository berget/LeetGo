package romantointeger

import (
	"fmt"
	"testing"
)

// TestRomanToInt table-driven tests
// 測試使用 table-driven pattern，包含 Happy-path 範例
func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{"III -> 3", "III", 3},
		{"LVIII -> 58", "LVIII", 58},
		{"MCMXCIV -> 1994", "MCMXCIV", 1994},
		{"IV -> 4", "IV", 4},
		{"IX -> 9", "IX", 9},
		{"XL -> 40", "XL", 40},
		{"CM -> 900", "CM", 900},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RomanToInt(tt.in)
			if got != tt.want {
				t.Fatalf("RomanToInt(%q) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

// BenchmarkRomanToInt 基準測試 / benchmark
func BenchmarkRomanToInt(b *testing.B) {
	s := "MCMXCIV"
	for i := 0; i < b.N; i++ {
		_ = RomanToInt(s)
	}
}

// ExampleRomanToInt 範例輸出 / example
func ExampleRomanToInt() {
	fmt.Println(RomanToInt("MCMXCIV"))
	// Output: 1994
}
