#export PATH:=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/go/bin


test:
	@go test -v

run:
	@go run example/usage.go

fmt:
	@go fmt *.go
	@go fmt example/*.go

help:
	@go help

build:
	go build -o ./bin/writer ./domain_tree/cmd/

protoc:
	@rm -f proto/tree.pb.go
	@protoc --go_out=. --go_opt=paths=source_relative tree_proto/tree.proto