package addtwonumbers

// AddTwoNumbers 將兩個以反向數位表示的非負整數（單向鏈表）相加，返回和的鏈表表示。
//
// 說明（中文）：
//   - 輸入：l1, l2，分別表示兩個整數，最低位在鏈表頭部（例如數字 342 表示為 [2,4,3]）。
//   - 輸出：一個新的鏈表，表示兩數之和，同樣為低位在頭的形式。
//   - 演算法：使用一個哨兵節點（dummy）和當前尾指針 cur。逐位相加 l1 與 l2 的節點值以及進位 carry，
//     每次計算 (sum = digit1 + digit2 + carry)，將 sum%10 作為新節點的值，carry = sum/10 保存到下一位。
//   - 迴圈條件：當任一鏈表仍有節點或 carry 不為 0 時繼續（處理不同長度與最後的進位）。
//   - 時間複雜度：O(max(n,m))，n 與 m 分別為 l1、l2 長度。
//   - 空間複雜度：O(max(n,m))（結果鏈表的節點數，忽略輸入空間）。
//   - 邊界情況：任一輸入為單節點 0；長度不同；最後有一個額外進位（例如 5+5 -> [0,1]）。
//
// 輸出契約：返回一個新的鏈表（不改變原來的 l1 或 l2 結構）。
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// dummy 為哨兵節點，方便返回頭節點並簡化邏輯
	dummy := &ListNode{}
	// cur 始終指向結果鏈表的尾節點，便於 append
	cur := dummy
	// carry 保存當前位的進位值（可能為 0 或更大的整數，但由於節點值在 0-9，實際為 0 或 1）
	carry := 0

	p, q := l1, l2
	// 當任一鏈表還有節點或還有進位時，繼續逐位計算
	for p != nil || q != nil || carry != 0 {
		// 從進位開始累加
		sum := carry
		if p != nil {
			// 若 p 有當前位，加入該位並前進
			sum += p.Val
			p = p.Next
		}
		if q != nil {
			// 若 q 有當前位，加入該位並前進
			sum += q.Val
			q = q.Next
		}
		// 更新進位（整除 10 得到進位，例如 15 -> carry = 1）
		carry = sum / 10
		// 當前位為 sum 的個位數，建立新節點並接到結果鏈表
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	// 返回哨兵節點之後的實際頭節點
	return dummy.Next
}
