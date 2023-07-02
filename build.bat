SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
SET GOPROXY=https://goproxy.cn
go build -o arkose_token cmd/main.go
