.PHONY: build

build:
	CGO_ENABLED=0 go build -o application main.go
test:
	go test -v ./... -cover
deps:
	go mod tidy
run:
	go run main.go server
migrate:
#	go run main.go migrate
generate:
	protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative proto/*.proto