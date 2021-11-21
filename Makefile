test:
	go test -mod vendor -test.v -cover ./...
dep:
	go mod init; go mod tidy && go mod vendor

gen:
	go generate ./...