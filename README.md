# 我第一個自己寫的網站！！

## 功能

可以新增貼文，編輯貼文，刪除貼文，就醬。

## 使用方式

### 環境需求
- golang
- npm
- PostgreSQL

### 資料庫設定

1. 開啟 PostgreSQL 並建立資料庫和使用者：
```sql
CREATE USER memeboard_user WITH PASSWORD 'your_password_here';
CREATE DATABASE memeboard OWNER memeboard_user;
```

2. 在 `backend/` 目錄建立 `.env` 檔案（可參考 `.env.example`）：
```bash
cp .env.example .env
# 編輯 .env，填入你的 PostgreSQL 密碼和其他設定
```

### 部署前端
```bash
cd frontend/memeboard/
npm install
npm run dev
```

### 部署後端
```bash
cd backend/
go mod tidy
go run main.go
```

網站前端預設會運行在 localhost:5173，訪問此 URL 即可使用此網站。
網站後端預設會運行在 localhost:8080。

## 廢話
第一次碰網站全端開發，啊我前後端沒一個會，摸索了很久終於有辦法自己寫出一個網站，雖然後端邏輯很簡單，前端 UIUX 也不怎麼好，但我還是可以碰到超多我不知道怎麼解決的問題，寫不到十行程式就得問一次 AI，不過我沒有叫 AI 幫我寫，這是我自己寫的，我很棒：）
阿對了，我是先叫 AI 寫一個網站（和這個不同的網站），讓我觀摩一下整個專案的架構大概長怎樣，vue 和 golang 和資料庫分別該怎麼寫這樣，這個專案則是我請 AI 給我一個目標，讓我自己想辦法把網站搓出來，我覺得做完之後我對網站開發有非常多理解上的進步，啊不過我現在還是很多東西不會就是了。
~~還有上面的readme是AI寫的，我甚至不知道一台空白的電腦要怎麼部署這個網站~~