build:
	@go build -o bin/ecommercewebsite cmd/main.go

test:
	@go test -v ./..

run: build
	@./bin/ecommercewebsite
