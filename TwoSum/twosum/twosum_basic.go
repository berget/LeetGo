package twosum

// twosum_basic.go
// 說明 (ZH): 這是一個非常基礎的新手版本 Two Sum 解法，使用雙重迴圈 (暴力法) 來尋找兩個數字之和等於目標值。
// Description (EN): A beginner-friendly Two Sum implementation using a simple double-loop (brute force).
//
// 目的 / Purpose:
//  - 展示最直觀的作法，幫助剛開始學 Go 的人了解索引存取、巢狀迴圈與回傳值的概念。
//  - Demonstrate direct array indexing, nested loops, and returning results.
//
// 使用方法 / Usage:
//  call: TwoSumBasic([]int{2,7,11,15}, 9) -> returns [0,1]
//
// 複雜度 / Complexity:
//  - 時間: O(n^2) (雙重迴圈，對每一個元素再掃描其後方所有元素)
//  - 空間: O(1) (僅使用固定大小的輸出陣列)
//
// 注意 (Notes):
//  - 這個版本沒有使用任何額外的資料結構（如 map），故較直觀但非最佳效能。
//  - The function assumes indices of a valid pair exist; if none found, it returns [-1, -1] as a sentinel.

// TwoSumBasic finds two distinct indices i and j such that nums[i] + nums[j] == target.
// TwoSumBasic 尋找兩個不同的索引 i 和 j，使得 nums[i] + nums[j] == target。
func TwoSumBasic(nums []int, target int) [2]int {
	// 我們在此使用最簡單、也最直觀的暴力法 (brute force):
	// 對每一個元素 i，再檢查它之後的每一個元素 j，看是否有和等於 target。
	// We use the simplest brute-force approach: for each index i, check every following index j.

	// 1) 先檢查輸入長度，避免不必要的運算或 panic。
	// 1) Quick guard: if there are fewer than 2 elements, 沒有解。
	if len(nums) < 2 {
		// 沒有足夠的元素可以配對 -> 回傳哨兵值
		// not enough elements to form a pair -> return sentinel
		return [2]int{-1, -1}
	}

	// 2) 使用兩層 for 迴圈遍歷所有 i < j 的組合
	//  Outer loop picks the first index i
	for i := 0; i < len(nums); i++ {
		// 內層從 i+1 開始，避免使用到同一個元素兩次 (也避免重複配對)
		// Inner loop starts at i+1 so we never pair an element with itself and avoid duplicate checks
		for j := i + 1; j < len(nums); j++ {
			// 說明：我們現在在比較 nums[i] 與 nums[j]
			// If the sum equals the target, 我們找到答案並立即回傳
			if nums[i]+nums[j] == target {
				// 回傳配對索引 (i, j)
				// return the pair of indices (i, j)
				return [2]int{i, j}
			}
			// 若不相等，內層迴圈繼續到下一個 j
			// otherwise continue checking next j
		}
		// 當前 i 的所有配對都檢查完，繼續下一個 i
	}

	// 若整個迴圈結束都沒有找到，回傳明確的哨兵 [-1, -1]
	// If no pair is found (unexpected for LeetCode constraints), return sentinel
	return [2]int{-1, -1}
}
