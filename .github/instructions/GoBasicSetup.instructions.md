---
applyTo: '**'
---

# Go 語言學習與專案設置指示 / Go Learning & Project Setup Instructions

## 目的 / Purpose:
// 本檔案用來指示 Agent 在建立或輔助建立 Go 專案時，必須遵循的規範與範例。
// This file instructs the Agent which rules and templates to follow when creating or assisting Go projects.

## 核心要求 / Core requirements (請務必全部遵守 / Must satisfy all):
1) 主要重點是教學 Golang 語法特性與實務應用 / Focus on teaching Go syntax features and practical usage
2) 當建立初始化檔案時，須包含必要的檔案內容範本 / When creating initial files, include file templates
3) 必須包含 `Makefile` 與測試項目（單元、benchmark、example）/ Must include Makefile and test items (unit, benchmark, example)
4) 所有註解與說明採繁體中文與英文對照 / All comments and explanations must be Traditional Chinese + English
5) 明確告訴 Agent 怎麼做（步驟/行動清單）/ Tell the Agent exactly how to act (steps/action list)

## 使用情境說明 / Usage guidance:
若用戶要求「建立新專案」或「初始化範本」，Agent 應自動生成以下檔案並回覆說明：
If the user requests "create project" or "initialize templates", the Agent should auto-generate the files below and reply with explanations:

## 預期至少產生的檔案 / Minimum files to generate:
- cmd/main.go            : 主程式範本含註解 / main program template with comments
- go.mod                 : module + minimal dependencies
- Makefile               : 建構、測試、自動化指令 / build/test automation
- tests/*_test.go        : 單元測試、benchmark、example 範本 / unit tests, benchmarks, examples
- README.md              : 專案說明與如何執行 / README with run/test instructions
- .gitignore, .env.example

## 行動步驟 / Agent action steps (實作須依序執行 / implement in order):
1. 先簡短說明（繁中/EN）要示範的語法特性與目的 / First explain (ZH/TW & EN) the syntax features and goals
2. 在工作區建立標準目錄結構 / Create the standard folder structure in workspace
3. 產生 `cmd/main.go` 範本，程式內含雙語註解示範語法（變數、常數、struct、interface、goroutine、channel、error）
4. 產生 `go.mod` 並 `go mod tidy` 建議的依賴（如 testify, gorilla/mux 可選）
5. 產生可用的 `Makefile`，包含：build, run, test, test-cover, test-verbose, fmt, vet, lint, deps-install, deps-update, clean, learn
6. 產生測試範本：單元測試、benchmark、example，並以 table-driven pattern 示範（所有註解雙語）
7. 提供如何執行與測試的步驟（繁中 / EN），並回傳示範輸出或測試結果範例
8. 若專案已有程式碼，先讀取現有檔案以避免覆寫，並將新增檔案放在正確位置

## 範例說明要點 / Explanation checklist for each generated file (Each file must include these):
- 檔案頂端：簡短說明（繁中 / EN）目的與如何使用 / Top-of-file short description (ZH/TW & EN)
- 註解：每個主要程式語法區塊皆提供繁中/EN 對照註解 / Comments: bilingual for each major code block
- 測試：至少包含一個 Happy-path 單元測試及一個 benchmark / Tests: at least one happy-path unit test and one benchmark
- Makefile：包含學習專用目標（例如 make learn）/ Makefile includes learning targets (e.g., make learn)

## 範本簡短範例（Agent 在生成檔案時可參考此格式）/ Minimal example templates (Agent may use as reference):

cmd/main.go 頂端註解示例 / cmd/main.go top comment example
"""
main.go
說明：示範 Go 基礎語法與一個簡單 HTTP 伺服器 / Demonstrates Go basics and a simple HTTP server
使用：go run cmd/main.go
"""

Makefile 目標說明示例 / Makefile target description example
help: 列出可用指令 / list available commands
build: 編譯專案 / build the project
test: 執行測試 / run tests

## 測試檔案注記示例 / test file note example
_test.go 檔案須包含 table-driven 測試，並以中文/英文描述每個案例 / _test.go must contain table-driven tests with bilingual descriptions

## 風險與限制 / Risks & constraints:
- Agent 不應在未經使用者同意下覆寫使用者已有重要檔案，除非使用者明確要求覆寫 / Do not overwrite user's important files without consent
- 若需外部網路動作（例如安裝工具或下載依賴），先在回覆中提示並提供相應命令 / If external network actions required, inform user and provide commands

## 回覆格式 / Agent reply format after generation:
1) 簡短說明（繁中 / EN）剛完成哪些檔案與要點 / Short summary of files created and highlights
2) 列出可執行的指令（copy/paste）與預期輸出範例 / Provide runnable commands and expected sample output
3) 提供下一步建議（例如新增功能或教學單元）/ Next steps suggestions

## 範例快速檢查清單（Agent 在完成後應回報）/ Quick checklist Agent reports back after completion:
- [ ] 已建立 `cmd/main.go`（含雙語註解）/ created cmd/main.go with bilingual comments
- [ ] 已建立 `go.mod` / created go.mod
- [ ] 已建立 `Makefile`（含 learn 目標）/ created Makefile including learn target
- [ ] 已建立至少一個 `_test.go` 範本 / created at least one _test.go template
- [ ] 已更新 README.md 說明執行/測試方式 / updated README.md with run/test instructions