// +build tools

package tools

import (
	_ "github.com/axw/gocov/gocov"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/t-yuki/gocov-xml"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/twitchtv/twirp/protoc-gen-twirp"
)
