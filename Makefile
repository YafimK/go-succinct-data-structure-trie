# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)


test:
	@go test -v

run:
	@go run example/usage.go

fmt:
	@go fmt *.go
	@go fmt example/*.go

help:
	@go help

protoc:
	@rm -f proto/tree.pb.go
	@protoc --go_out=. --go_opt=paths=source_relative tree_proto/tree.proto