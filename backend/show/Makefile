build:
	go build -o bin/show

run: build
	./bin/show

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/*.proto

.PHONY: proto