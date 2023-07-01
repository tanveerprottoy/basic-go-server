run:
	go run cmd/basicserver/main.go
build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/app cmd/basicserver/main.go