# Go

## Script

```sh
# go.mod初始化
go mod init mygo
# 套件:單一安裝
go get github.com/gin-gonic/gin
go get github.com/jinzhu/gorm
# 套件:全部安裝
go mod download
# 套件:移除未使用
go mod tidy
# 看環境變數
go env
# 測試
go test ./...
```

## Features

- restful/mvc/repository+service
- graphql
- jwt
- test
