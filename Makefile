CONTAINER_NAME := mss_api

run:
	go run ./server.go

cl:
	clear

rund: build
	@./bin/api

build:
# Faz o build primeiro
	@go build -o bin/api

test:
	@go test -v ./...

dc:
	@docker build -t api:mss .

cr:
	@docker run --rm -p 8001:8001 --name go-mss-app api:mss

tr: 
	@docker build -t api:mss .  &&  @docker run --rm -p 8001:8001 api:mss 