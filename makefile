.PHONY: build
build:
	go build -o ./build/sbi-port-manager .

.PHONY: test
test:
	gofmt -l -w .
	go test ./... -v -cover

.PHONY: push
push:
	docker build -t ghcr.io/azuki774/sbi-port-manager:nightly -f build/Dockerfile .
	docker push ghcr.io/azuki774/sbi-port-manager:nightly
