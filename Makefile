.PHONY: default
default: build

.PHONY: generate
generate:
	SEED=$(shell go run internal/generate/seed/main.go) go generate ./...

.PHONY: build
build: generate
	go build ./cmd/ecsu
	go build ./cmd/ecsu-keygen

.PHONY: clean
clean:
	rm ./ecsu
	rm ./ecsu-keygen
	rm ./cmd/ecsu/ecsu_key.go
	rm ./cmd/ecsu-keygen/keygen_key.go
