package palindromenumber

/*
palindrome_annotated.go
說明: 極為詳盡註解的初學者版回文檢查（純數值實作，不使用任何 import）
Description: Heavily annotated beginner version palindrome check (numeric-only, no imports)
Usage: go test ./... or go test

此檔案的重點：
- 每一行主要程式碼皆有繁中與英文註解說明意圖與原因，方便初學者閱讀。
- 不使用 strconv 或其他標準套件，僅用純算術（% 和 /）與整數運算來實作。
- 範例與複雜度說明也包含在註解中。

注意 / Note:
- 直接反轉整個整數（如本檔）在某些語言或情況可能會發生整數溢位（overflow），
  因此此實作主要為教學用途；在 production 可考慮採用只反轉一半的方法或字串方法。
  Reversing the whole integer can overflow for very large values; this is for learning/demo.
*/

// IsPalindromeAnnotated 用最直觀的方式：反轉整個整數，並與原始值比較，
// 每一個步驟都有繁中/EN 註解說明。
// IsPalindromeAnnotated reverses the whole integer and compares to original with step-by-step comments.
//
// 算法複雜度 / Complexity:
// - 時間：O(d)，d 為數字的位數（每位處理一次）
// - 空間：O(1)，只使用常數個額外變數
func IsPalindromeAnnotated(x int) bool {
	// 1) 處理負數：任何帶有負號的數字都不是回文（因為 '-' 與數字不對稱）
	// 1) Negative numbers: not palindromes because of the leading '-' sign
	if x < 0 {
		return false
	}

	// 2) 保存原始值以便最後比較。
	// 2) Save the original value so we can compare after reversing.
	original := x

	// 3) 準備一個變數來累積反轉後的數字（從 0 開始）。
	// 3) Prepare a variable to accumulate the reversed number (start at 0).
	rev := 0

	// 4) 逐位反轉：
	//    - 取出最後一位 digit = x % 10
	//    - 把 rev 左移一位（rev*10）並加入 digit（rev = rev*10 + digit）
	//    - 把原數 x 去掉最後一位（x = x / 10）
	// 4) Reverse digits one by one:
	//    - digit = x % 10 (pop last digit)
	//    - rev = rev*10 + digit (push digit onto rev)
	//    - x = x / 10 (remove last digit from x)
	for x != 0 {
		// 取尾數 / get the last digit
		digit := x % 10

		// 將目前 rev 左移一位並加入 digit
		// shift rev left (multiply by 10) and add the digit
		rev = rev*10 + digit

		// 移除 x 的最後一位（整數除法）
		// remove the last digit from x via integer division
		x /= 10
	}

	// 5) 當迴圈結束，rev 為原始數字的反轉。如果 rev 與 original 相同，則為回文。
	// 5) After the loop, rev is the reversed number. If rev == original, it's a palindrome.
	// 注意：如果輸入非常大，rev*10 + digit 的運算可能造成整數溢位—這裡不做 overflow 檢查以保留範例簡潔性。
	// Note: For very large inputs rev = rev*10 + digit may overflow; overflow checks omitted here for clarity.
	return rev == original
}
