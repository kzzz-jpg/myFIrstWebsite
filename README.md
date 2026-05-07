# 我第一個自己寫的網站！！

## 功能

可以新增貼文，編輯貼文，刪除貼文，就醬。

## 使用方式

確保你電腦上有安裝 golang、npm、postgreSQL。
部署前端：
```bash
cd frontend/memeboard/
npm run dev
```

部署後端：
```bash
cd backend/
go mod tidy
go run main.go
```

網站前端預設會運行在 localhost:5173，訪問此 URL 即可使用此網站。
網站後端預設會運行在 localhost:8080。

## 廢話
第一次碰網站全端開發，啊我前後端沒一個會，摸索了很久終於有辦法自己寫出一個網站，雖然後端邏輯很簡單，前端 UIUX 也不怎麼好，但我還是可以碰到超多我不知道怎麼解決的問題，寫不到十行程式就得問一次 AI，不過我沒有叫 AI 幫我寫，這是我自己寫的，我很棒：）