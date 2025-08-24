package addtwonumbers

/*
addtwonumbers_basic.go
說明 (ZH):
  這是一個更基礎版本的 Add Two Numbers 解法，特別為剛學 Golang 的新手設計。
  程式中刻意使用直觀的變數名稱和大量中英雙語註解，說明每一步的意圖與邏輯。

Description (EN):
  A more basic version of the Add Two Numbers solution, aimed at beginners learning Go.
  The code uses clear variable names and many bilingual comments (Traditional Chinese + English)
  to explain the algorithm step by step.

注意：
  - 假定本 package 中已有 `ListNode` 的定義（與專案內其他檔案共用），因此這裡不會重新定義 `ListNode`。
  - 本檔案不使用任何外部模組（only standard Go constructs）。
*/

// AddTwoNumbersBasic 將兩個以反向數位表示的非負整數（單向鏈表）相加，回傳和的鏈表表示（教學版）。
// AddTwoNumbersBasic adds two non-negative integers represented by linked lists in reverse order and returns the sum as a linked list (educational version).
//
// 輸入契約 / Input contract:
//   - l1, l2: 各為指向 ListNode 的指標，代表整數的每一位；最低位（units）在鏈表的頭部。
//   - 範例: 數字 342 會被表示為 [2 -> 4 -> 3]
//
// 輸出 / Output:
//   - 回傳一個新的鏈表，表示兩數之和（相同的低位在頭格式）。
//
// 實作要點 / Key ideas:
//  1. 使用一個 "dummy" （哨兵）節點來簡化頭節點的處理（避免處理空結果時的特殊情況）。
//  2. 使用一個 `carry` 變數儲存進位（可能為 0 或 1，因為每位相加最大為 9+9+1=19）。
//  3. 逐位取出 l1 與 l2 的值（若節點為 nil，視為 0），相加後建立新節點放到結果鏈表尾端。
//  4. 當兩個輸入鏈表都走完且沒有進位時結束。
//
// 時間複雜度：O(max(n,m))
// 空間複雜度：O(max(n,m))  (輸出鏈表大小)
func AddTwoNumbersBasic(l1 *ListNode, l2 *ListNode) *ListNode {
	// 先建立一個哨兵 (dummy) 節點，方便最後回傳結果的第一個有效節點。
	// Create a dummy node to simplify appending and returning the head later.
	dummy := &ListNode{}

	// tail 會指向結果鏈表目前的最後一個節點（開始時指向 dummy）。
	// tail points to the last node in the result list (initially the dummy).
	tail := dummy

	// carry 儲存進位，初始為 0。
	// carry stores the carry value between digit additions.
	carry := 0

	// p, q 分別用來遍歷 l1 與 l2。若某一邊先走完，我們會把它當作 0 處理。
	// p and q iterate over l1 and l2 respectively.
	p, q := l1, l2

	// 當有任一個節點還沒處理或仍有進位時，繼續處理。
	// Continue while either list has nodes or there is a carry.
	for p != nil || q != nil || carry != 0 {
		// 先把目前的進位放到 sum 裡面。
		// Start sum with the current carry.
		sum := carry

		// 如果 p 還有節點，就把它的值加進 sum，然後前進 p。
		// If p is not nil, add its value and move to next.
		if p != nil {
			sum += p.Val
			p = p.Next
		}

		// 同理，處理 q。
		// Same for q.
		if q != nil {
			sum += q.Val
			q = q.Next
		}

		// sum 目前是兩位數相加 + 可能的進位，新的進位為 sum / 10（整數除法）。
		// The new carry is sum / 10.
		carry = sum / 10

		// 當前位的數值是 sum % 10。
		// The digit for this position is sum % 10.
		digit := sum % 10

		// 建立一個新節點，值為 digit，並把它接到結果鏈表的尾端。
		// Create a new node with the digit and append it to the result.
		tail.Next = &ListNode{Val: digit}
		// 把 tail 移到剛剛新增的節點，保持尾端指標正確。
		// Move tail to the new last node.
		tail = tail.Next
	}

	// 迴圈結束後，dummy.Next 就是結果鏈表的真正頭節點。
	// Return the real head which follows dummy.
	return dummy.Next
}

// 注意：
// - 這個函式的寫法刻意保持簡單且逐步說明，適合教學用途。若要追求更短的實作，
//   可將一些步驟合併（例如直接在建立節點時使用 sum%10 而不另存 digit）。
// - 本檔案不依賴任何外部套件；只使用 Go 基礎語法與資料結構。
//
// Edge cases (邊界情況) to keep in mind:
// - l1 或 l2 為 nil（代表 0）
// - l1 與 l2 長度不同
// - 最後有一個進位需要新增節點（例如 5 + 5 = 10，結果 [0 -> 1]）
