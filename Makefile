.EXPORT_ALL_VARIABLES:
GO111MODULE=on

deps:
	go get ./...

test:
	go test -v -race ./...
