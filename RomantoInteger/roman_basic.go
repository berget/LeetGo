package romantointeger

// roman_simple.go
// 說明 (ZH/TW):
//   一個更基礎的羅馬數字轉換實作，使用最簡單的語法與大量註解，適合剛學 Go 的新手。
// Description (EN):
//   A beginner-friendly implementation to convert Roman numerals to integers.
//   Uses only basic Go constructs and extensive comments to explain intent.

// RomanToIntBasic 是一個容易理解的版本，會把傳入的羅馬字串轉為整數
// RomanToIntBasic is an easy-to-read version that converts a Roman numeral string to an integer
func RomanToIntBasic(s string) int {
	// 我們不使用任何外部套件，只有內建語法：迴圈、條件與 switch。
	// We don't use external packages — only loops, conditionals and switch.

	sLen := len(s)

	// 長度為零的字串直接回傳 0（依需求也可以改為回傳錯誤，但這裡保持簡單）。
	// If input is empty return 0. A real API might return an error instead.
	if sLen == 0 {
		return 0
	}

	// total 用來累加最終結果
	// total accumulates the final integer value
	total := 0

	// 我們一步一步走過字串的每一個字元（byte），因為羅馬字母都是 ASCII
	// We iterate bytes (ASCII letters) since Roman numerals use ASCII characters
	for i := 0; i < sLen; i++ {
		// 當前字元的數值
		// value of the current Roman character
		v := romanCharValue(s[i])

		// 預設把當前值加入 total，但若下一個字元存在且比當前值大，
		// 代表這是一個減法情況（例如 IV => I before V means 4），
		// 這時我們應該減去當前字元的值。
		// By default add the current value. If the next value is larger, this
		// is a subtraction case (e.g. IV), so subtract the current value.
		if i+1 < sLen {
			next := romanCharValue(s[i+1])
			if v < next {
				// 減法情況
				total -= v
				continue
			}
		}

		// 正常情況，直接加上當前值
		total += v
	}

	// 回傳累加結果
	// return the accumulated total
	return total
}

// romanCharValue 回傳單一羅馬字元對應的整數值
// romanCharValue returns the integer value of a single Roman character
func romanCharValue(c byte) int {
	// 我們使用最基本的 switch 來做對應，對初學者來說比 map 更直觀
	// Use a simple switch instead of a map for clarity to beginners
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		// 如果遇到未知字元，回傳 0（可視情況改為回傳錯誤）
		// Unknown characters return 0 here; a stricter implementation would return an error
		return 0
	}
}

/*
使用說明 / Notes:
- 複雜度 (Complexity): O(n) 時間，因為我們只掃描一次字串；O(1) 額外空間。
- 減法規則 (Subtraction rule): 當一個較小的符號放在較大的符號左邊時，
  應該用減法處理（例如 'I' 在 'V' 前 -> 4）。
- 邊界情況 (Edge cases): 空字串會回傳 0；未知字元會被視為 0。

範例 / Example:
  RomanToIntBasic("III") -> 3
  RomanToIntBasic("IV")  -> 4
  RomanToIntBasic("MCMXCIV") -> 1994
*/
