include .env
clean:
	@go mod tidy
run:
	@go run main.go
checkrace:
	@go run -race main.go
	
build:
	@go build -o swapi cmd/main.go

# GO111MODULE=on go get mvdan.cc/gofumpt
lint:
	@gofumpt -l -w .
test:
	@go test -v -count=1 ./...