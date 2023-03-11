SHELL := /bin/bash
BINARY=aragog
export PATH := $(CURDIR)/_tools/bin:$(PATH)
USER ?= `whoami`

.PHONY: all
.DEFAULT: all

all: mod test

tidy:
	go mod tidy

mod:
	go mod vendor

docs:
	# will create docs

test:
	go test -vv internal/...

format: mod
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run --fix

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run -v --timeout 5m --out-format checkstyle > golangci-report.xml
.PHONY: lint

gocov:
	mkdir -p ./coverage
	CGO_ENABLED=1 go run github.com/axw/gocov/gocov test --race ./internal/... | go run github.com/t-yuki/gocov-xml > coverage/coverage.xml
.PHONY: gocov

verify: lint gocov
.PHONY: verify

build:
	# linter
	go build -o _tools/bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

install:
	go get -d \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \

PROTOC-exists:
	@which protoc > /dev/null

proto: PROTOC-exists
	protoc -I ./rpc \
       --go_out ./rpc/ --go_opt paths=source_relative \
       --go-grpc_out ./rpc/ --go-grpc_opt paths=source_relative \
       --python_out=./rpc/ \
       --js_out=./rpc/ \
       ./rpc/aragog/aragog.proto


coverage:
	@rm -f coverage.txt
	CGO_ENABLED=1 go test -race -coverprofile=coverage.txt -covermode=atomic ./internal/...
	go tool cover -html=coverage.txt
.PHONY: coverage

local:
	docker build -t ${BINARY} . && docker run -p 8080:8080 ${BINARY}


$(BINARY):
	go build ${LDFLAGS} -o ${BINARY} ./cmd/aragog
