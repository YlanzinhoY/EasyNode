build:
	GOOS=windows GOARCH=amd64 go build -o EasyNode/easynode.exe
	GOOS=linux GOARCH=amd64 go build -o EasyNode/easynode
