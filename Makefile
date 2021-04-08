MAJOR ?= 1
MINOR ?= 2
PATCH ?= 0
BUILD ?= release

LDFLAGS := -X=main.MAJOR=$(MAJOR)
LDFLAGS += -X=main.MINOR=$(MINOR)
LDFLAGS += -X=main.PATCH=$(PATCH)
LDFLAGS += -X=main.BUILD=$(BUILD)

GOFLAGS := -ldflags="$(LDFLAGS)"

.PHONY: all
all: release

.PHONY: release
release:
	mkdir -p release
	GOOS=darwin GOARCH=amd64 go build $(GOFLAGS)
	tar czvf keysmith-darwin-amd64.tar.gz keysmith
	mv keysmith-darwin-amd64.tar.gz release
	GOOS=linux GOARCH=amd64 go build $(GOFLAGS)
	tar czvf keysmith-linux-amd64.tar.gz keysmith
	mv keysmith-linux-amd64.tar.gz release
	GOOS=linux GOARCH=arm64 go build $(GOFLAGS)
	tar czvf keysmith-linux-arm64.tar.gz keysmith
	mv keysmith-linux-arm64.tar.gz release

.PHONY: clean
clean:
	rm -fr release
