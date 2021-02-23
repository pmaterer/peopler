cover:
	go test ./... -coverprofile cover.out 
	go tool cover -func cover.out

lint:
	golangci-lint run

fmt:
	go fmt ./...

test:
	go test -v ./...

run: fmt test
	go run cmd/peopler.go

build:
	go build cmd/peopler.go