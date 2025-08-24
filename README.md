# Leetcode - Go 練習專案

這個專案用來透過 LeetCode 題目練習與學習 Golang 的語法特性與應用。每個題目放在獨立的資料夾中，包含題目的實作、測試檔和簡短的說明。

## 專案目標

- 練習 Go 語法（types、slices、maps、interfaces、goroutines 等）。
- 熟悉 Go 的專案結構與模組管理（go.mod）。
- 使用 table-driven tests、benchmark 和小型範例來驗證與理解演算法。

## 主要目錄結構

- `AddTwoNumbers/` - 題目：Add Two Numbers。
  - `addtwonumbers.go` - 解題主要程式。
  - `addtwonumbers_test.go` - 單元測試。
  - `listnode.go` - 共用的 ListNode 定義。
  - `README.md` - 題目說明（題目級別的 README）。
- `TwoSum/` - 題目：Two Sum（包含 sub-package 範例）。
  - `twosum/` - 實作與測試。
  - `cmd/` - 範例 main (如果有需要手動執行範例)

（每個題目資料夾通常會是獨立的 go module 或在同一 module 下的 package，視情況而定。）

## 開發與執行測試

需要在 macOS 或 Linux 環境上安裝 Go（建議使用 Go 1.20+）。

在專案根目錄下執行全部套件的測試：

```zsh
cd /Users/james/www/Go_Project/Leetcode
go test ./...
```

單獨執行某個題目的測試（範例）：

```zsh
cd AddTwoNumbers
go test -v
```

若某個子資料夾是獨立 module（有自己的 `go.mod`），請在該子資料夾裡執行 `go test`。

## 程式風格與慣例

- 使用 table-driven tests（對於多個輸入/輸出案例）。
- 優先使用標準函式庫（`testing`, `fmt`, `errors` 等）。
- 命名盡量直觀：函式與變數以駝峰式（camelCase）或符合 Go 慣例。
- 保持小型、單一職責的函式，並為邊界情況寫測試。

## 練習建議

- 每題先寫出能通過測試的簡單實作，再嘗試優化複雜度（時間/空間）。
- 為常見邊界情況撰寫額外測試（空輸入、極端值、nil）。
- 使用 `go vet` 與 `golangci-lint`（選用）來提升程式品質。

## 貢獻

歡迎新增題目或優化既有解法：

1. 在專案根目錄新增新的資料夾（例如 `TwoPointers/Add`）。
2. 加入 `go.mod`（如果你想讓它成為獨立 module），或在現有 module 中新增 package。 
3. 實作 `.go` 檔與對應的 `_test.go`。
4. 確保 `go test ./...` 全部通過後再開 Pull Request。

## 其他資源

- 官方 Go 文件：[https://golang.org/doc/](https://golang.org/doc/)
- Effective Go、Go Proverbs 與各種博文可助於寫出更慣用的 Go 風格程式碼。
