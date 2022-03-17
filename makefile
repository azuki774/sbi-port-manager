.PHONY: build
build:
	go build -o ./build/sbi-port-manager .

.PHONY: test
test:
	gofmt -l -w .
	go test ./... -v -cover

