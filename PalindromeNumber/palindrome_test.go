package palindromenumber

import (
	"fmt"
	"testing"
)

/*
palindrome_test.go
說明: table-driven 單元測試、benchmark 與 example 範例
*/

func TestIsPalindrome(t *testing.T) {
	// table-driven 測試範例 (Table-driven tests)
	tests := []struct {
		name string
		in   int
		want bool
	}{
		{"positive palindrome", 121, true},
		{"negative number", -121, false},
		{"ends with zero", 10, false},
		{"single digit", 7, true},
		{"even digits palindrome", 1221, true},
		{"large non-palindrome", 123456, false},
		{"large palindrome odd", 12321, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsPalindrome(tc.in)
			if got != tc.want {
				t.Fatalf("IsPalindrome(%d) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsPalindrome(123454321)
	}
}

// ExampleIsPalindrome demonstrates usage and will be used by `go test` as an example.
func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome(121))
	// Output: true
}
