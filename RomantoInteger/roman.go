package romantointeger

// roman.go
// 說明 (ZH/TW):
//   將羅馬數字字串轉換為整數。示範 map、迴圈與減法規則處理。
// Description (EN):
//   Convert a Roman numeral string to an integer. Demonstrates using maps,
//   looping and handling subtraction cases.

// RomanToInt 轉換羅馬數字為整數 / Convert Roman numeral to integer
func RomanToInt(s string) int {
	// mapping of Roman numerals to values
	vals := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	n := len(s)
	if n == 0 {
		return 0
	}

	total := 0
	// iterate over characters; when a smaller value precedes a larger one,
	// subtract the smaller instead of adding.
	for i := 0; i < n; i++ {
		v := vals[s[i]]
		if i+1 < n {
			next := vals[s[i+1]]
			if v < next {
				total -= v
				continue
			}
		}
		total += v
	}
	return total
}
