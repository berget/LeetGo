package palindromenumber

/*
palindrome.go
說明: 判斷整數是否為回文 (Palindrome check for integers)
Usage: go test ./... or go test
*/

// IsPalindrome 檢查一個整數是否為回文。
// IsPalindrome checks whether an integer is a palindrome.
//
// 詳細說明 / Detailed explanation:
// - 為什麼不直接把整數轉成字串再比對？
//
//  1. 教學與效率：使用數學方法可示範數字處理技巧，並避免額外的字串分配。
//
//  2. 可避免在某些語言或平台上將整數整個反轉時可能造成的溢位（overflow）問題。
//     Why not convert to string:
//
//  1. Shows numeric manipulation and avoids allocation.
//
//  2. Avoids risk of integer overflow if reversing the whole number.
//
//     - 算法主旨：只反轉數字的一半（reverse half）再比較。
//     原理：當你逐步把原數字的尾數取出疊到 rev，並把原數字整除10到 n，
//     當 rev >= n 時，表示我們已經處理至少一半的位數。這時候只要比較 rev 與 n（或 rev/10 與 n，
//     若位數為奇數的情況需要丟棄中間那個位數）即可。
//     Key idea: reverse only half of the digits. When rev >= n we've processed half (or more).
//     For even-length numbers compare n==rev. For odd-length numbers the middle digit is in rev, so compare n==rev/10.
//
// - 早期排除條件 / Early exits:
//
//   - 負數 (x < 0)：含負號不是回文。
//
//   - 若末位為0但 x != 0：回文數不會以0開頭（leading zero），因此排除。
//
//   - 時間/空間複雜度 / Complexity:
//     Time: O(d) where d is number of digits (~log10(x)).
//     Space: O(1) (constant extra space).
//
//   - 範例 / Example walkthrough:
//     1221: n=1221, rev=0 -> rev=1,n=122 -> rev=12,n=12 -> stop (n==rev) => palindrome.
//     12321: n=12321,rev=0 -> rev=1,n=1232 -> rev=12,n=123 -> rev=123,n=12 -> stop (rev>n)
//     compare n==rev/10 -> 12 == 123/10 (12) => palindrome.
//
// 輸入/Input: x int
// 輸出/Output: bool
func IsPalindrome(x int) bool {
	// If x < 0 then it's not a palindrome because of the minus sign.
	if x < 0 {
		return false
	}
	// If last digit is 0 and x != 0 then can't be palindrome (leading zero would be required).
	if x%10 == 0 && x != 0 {
		return false
	}

	rev := 0
	n := x
	// Reverse half of the number
	for n > rev {
		rev = rev*10 + n%10
		n /= 10
	}

	// For odd digits, discard the middle digit before comparison
	return n == rev || n == rev/10
}
