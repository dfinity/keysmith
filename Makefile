MAJOR ?= 1
MINOR ?= 2
PATCH ?= 0
BUILD ?= release

LDFLAGS := -X=main.MAJOR=$(MAJOR)
LDFLAGS += -X=main.MINOR=$(MINOR)
LDFLAGS += -X=main.PATCH=$(PATCH)
LDFLAGS += -X=main.BUILD=$(BUILD)

GOFLAGS := -ldflags="$(LDFLAGS)"

PRIVATE_KEY ?= $(HOME)/.openssl/private.key

.PHONY: all
all: release

.PHONY: release
release:
	mkdir -p release
	GOOS=darwin GOARCH=amd64 go build $(GOFLAGS)
	tar -c -f keysmith-darwin-amd64.tar.gz -z keysmith
	openssl dgst -sha256 keysmith-darwin-amd64.tar.gz > SHA256.SUM
	mv keysmith-darwin-amd64.tar.gz release
	GOOS=linux GOARCH=amd64 go build $(GOFLAGS)
	tar -c -f keysmith-linux-amd64.tar.gz -z keysmith
	openssl dgst -sha256 keysmith-linux-amd64.tar.gz >> SHA256.SUM
	mv keysmith-linux-amd64.tar.gz release
	GOOS=linux GOARCH=arm64 go build $(GOFLAGS)
	tar -c -f keysmith-linux-arm64.tar.gz -z keysmith
	openssl dgst -sha256 keysmith-linux-arm64.tar.gz >> SHA256.SUM
	mv keysmith-linux-arm64.tar.gz release
	openssl dgst -out SHA256.SIG -sha256 -sign $(PRIVATE_KEY) SHA256.SUM
	mv SHA256.SUM SHA256.SIG release

.PHONY: clean
clean:
	rm -f -r release
