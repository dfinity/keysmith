MAJOR ?= 1
MINOR ?= 6
PATCH ?= 0

LDFLAGS := -X=main.MAJOR=$(MAJOR)
LDFLAGS += -X=main.MINOR=$(MINOR)
LDFLAGS += -X=main.PATCH=$(PATCH)

PRIVATE_KEY ?= $(HOME)/.openssl/private.key

.PHONY: all
all: build

.PHONY: build
build:
	$(eval BUILD ?= $(shell git describe --always --dirty))
	$(eval LDFLAGS += -X=main.BUILD=$(BUILD))
	$(eval GOFLAGS := -ldflags="$(LDFLAGS)")
	go build $(GOFLAGS)

.PHONY: release
release:
	$(eval BUILD ?= release)
	$(eval LDFLAGS += -X=main.BUILD=$(BUILD))
	$(eval GOFLAGS := -ldflags="$(LDFLAGS)")

	# Create release directory.
	mkdir -p release

	# Create checksum file.
	> SHA256.SUM

	# Create completions script.
	openssl dgst -sha256 keysmith-completions.bash >> SHA256.SUM
	cp keysmith-completions.bash release

	# Create release tarball for darwin/amd64.
	GOOS=darwin GOARCH=amd64 go build $(GOFLAGS)
	tar -c -f keysmith-darwin-amd64.tar.gz -z keysmith
	openssl dgst -sha256 keysmith-darwin-amd64.tar.gz >> SHA256.SUM
	mv keysmith-darwin-amd64.tar.gz release

	# Create release tarball for darwin/arm64.
	GOOS=darwin GOARCH=arm64 go build $(GOFLAGS)
	tar -c -f keysmith-darwin-arm64.tar.gz -z keysmith
	openssl dgst -sha256 keysmith-darwin-arm64.tar.gz >> SHA256.SUM
	mv keysmith-darwin-arm64.tar.gz release

	# Create release tarball for linux/amd64.
	GOOS=linux GOARCH=amd64 go build $(GOFLAGS)
	tar -c -f keysmith-linux-amd64.tar.gz -z keysmith
	openssl dgst -sha256 keysmith-linux-amd64.tar.gz >> SHA256.SUM
	mv keysmith-linux-amd64.tar.gz release

	# Create release tarball for linux/arm32.
	GOOS=linux GOARCH=arm go build $(GOFLAGS)
	tar -c -f keysmith-linux-arm32.tar.gz -z keysmith
	openssl dgst -sha256 keysmith-linux-arm32.tar.gz >> SHA256.SUM
	mv keysmith-linux-arm32.tar.gz release

	# Create release tarball for linux/arm64.
	GOOS=linux GOARCH=arm64 go build $(GOFLAGS)
	tar -c -f keysmith-linux-arm64.tar.gz -z keysmith
	openssl dgst -sha256 keysmith-linux-arm64.tar.gz >> SHA256.SUM
	mv keysmith-linux-arm64.tar.gz release

	# Create release tarball for windows/amd64.
	GOOS=windows GOARCH=amd64 go build $(GOFLAGS)
	tar -c -f keysmith-windows-amd64.tar.gz -z keysmith.exe
	openssl dgst -sha256 keysmith-windows-amd64.tar.gz >> SHA256.SUM
	mv keysmith-windows-amd64.tar.gz release

	# Sign release tarballs.
	openssl dgst -out SHA256.SIG -sha256 -sign $(PRIVATE_KEY) SHA256.SUM
	mv SHA256.SIG SHA256.SUM release

.PHONY: clean
clean:
	rm -f -r keysmith keysmith.exe release
