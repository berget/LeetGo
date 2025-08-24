package twosum

// TwoSum finds two distinct indices i and j such that nums[i] + nums[j] == target.
// TwoSum 尋找兩個不同的索引 i 和 j，使得 nums[i] + nums[j] == target。
//
// Algorithm (single-pass hash):
// - Iterate the slice once. For each element v at index i compute need = target - v.
// - If need has already been seen (stored in the map), we found the pair (seen[need], i).
// - Otherwise record the current value -> index in the map and continue.
//
// 演算法（單次掃描 + 哈希表）：
// - 對陣列做單次遍歷。對於每個位於索引 i 的元素 v，計算 need = target - v。
// - 若之前已經見過 need（存在哈希表中），則找到了匹配的索引對 (seen[need], i)。
// - 否則把當前值與索引記錄到哈希表中，繼續向後處理。
//
// Complexity:
// - Time: O(n) single pass over the input.
// - Space: O(n) to store seen values in the hash map.
//
// 複雜度：
// - 時間：O(n)，只需對輸入遍歷一次。
// - 空間：O(n)，用於在哈希表中保存已見過的值與索引。
//
// Assumptions / notes:
//   - The problem guarantees exactly one valid answer exists, so the function returns
//     as soon as a match is found. If no match is found the function returns [-1, -1].
//   - The check is performed before inserting the current value into the map, so
//     the algorithm never uses the same element twice.
//   - The returned indices may be in any order.
//
// 假設 / 注意事項：
//   - 題目保證恰好有一個解，因此在找到匹配時函式會立即返回。若沒有找到（不符合題意的輸入），
//     函式會回傳 [-1, -1] 作為明確的哨兵值。
//   - 在把目前元素插入表之前就檢查其互補值，能確保不會使用到同一個元素兩次。
//   - 回傳的索引順序可以是任意的。
func TwoSum(nums []int, target int) [2]int {
	// 'seen' maps a value -> index for values we've already iterated.
	// Using the map lets us check complements in O(1) average time.
	//
	// 'seen' 將數值映射到其索引（value -> index），代表我們已經遍歷過的元素。
	// 使用哈希表可以在平均 O(1) 時間內檢查互補值是否存在。
	seen := make(map[int]int, len(nums))

	// Walk the slice once, keeping the current index and value.
	// 單次遍歷陣列，同時追蹤目前的索引與數值。

	// Example walkthrough (nums = [2,7,11,15], target = 9):
	// 範例步驟（nums = [2,7,11,15], target = 9）：
	// Iteration 0:
	//  i=0, v=2, need = 9 - 2 = 7
	//  seen = {} -> 7 not in seen -> record seen[2] = 0
	//  (no return; continue)
	//  迴圈 0：
	//   i=0, v=2, need = 9 - 2 = 7
	//   seen = {} -> 7 不在 seen -> 記錄 seen[2] = 0
	//   (不返回，繼續)
	//
	// Iteration 1:
	//  i=1, v=7, need = 9 - 7 = 2
	//  seen = {2:0} -> 2 in seen with index 0 -> return [0,1]
	//  (function returns, algorithm stops)
	//  迴圈 1：
	//   i=1, v=7, need = 9 - 7 = 2
	//   seen = {2:0} -> 2 在 seen，對應索引為 0 -> 回傳 [0,1]
	//   (函式返回，演算法結束)
	for i, v := range nums {
		// 'need' is the complementary value required to reach target: need + v == target
		// 'need' 是要達成 target 所需的互補值：need + v == target
		need := target - v

		// If we've already seen 'need', its index is a valid pair with current i.
		// We check the map before inserting the current value to avoid using the same element twice.
		//
		// 如果先前已經見過 need，則其索引與目前索引 i 組成一對合法解。
		// 在插入目前值之前先檢查，能避免使用到相同元素兩次。
		if j, ok := seen[need]; ok {
			return [2]int{j, i}
		}

		// Record the current value and its index for future complements to find.
		// 將目前的值與索引記錄到哈希表，供後續元素檢查互補值時使用。
		seen[v] = i
	}

	// Per problem statement this should not happen (exactly one solution exists).
	// Return an explicit sentinel in case of unexpected input.
	// 根據題目此分支不應發生（保證存在且只有一個解）；若發生非預期輸入，回傳哨兵值。
	return [2]int{-1, -1}
}
