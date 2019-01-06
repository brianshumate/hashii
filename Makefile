BINARY:=hashii
PACKAGESRC:=main.go
SOURCEDIR:=.
SOURCES:=$(shell find $(SOURCEDIR) -name '*.go')
VERSION:=$(shell cat version.txt)
BUILD_TIME:=$(shell date +%FT%T%z)
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_BUILD_RACE=$(GO_CMD) build -race
GO_TEST=$(GO_CMD) test
GO_TEST_VERBOSE=$(GO_CMD) test -v
GO_INSTALL=$(GO_CMD) install -v
GO_CLEAN=$(GO_CMD) clean
GO_DEPS=$(GO_CMD) get -d -v
GO_DEPS_UPDATE=$(GO_CMD) get -d -v -u
GO_VET=$(GO_CMD) vet
GO_FMT=$(GO_CMD) fmt
GO_LINT=golint
MAN_DIR=$(SOURCEDIR)/doc
MAN_SRC=hashii.1.ronn
MAN_SYS_LOCAL=/usr/local/share/man/man1
MAN_PAGE=hashii.1
WEB_PAGE=hashii.1.html
SHARE_DIR=./share
RONN=ronn
LDFLAGS=-ldflags "-X main.version=${VERSION} -X github.com/brianshumate/hashii/core.Version=${VERSION} -X github.com/brianshumate/hashii/core.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	$(GO_BUILD) ${LDFLAGS} -o ${BINARY} ${PACKAGESRC}

all: build

build: fmt lint vet
	@echo "==> Build $(PACKAGESRC) ..."; \
	$(GO_BUILD) ${LDFLAGS} -o ${BINARY} $(SOURCEDIR)/$(PACKAGESRC) || exit 1; \

install: clean all doc
	@echo "==> Install $(BINARY) ..."; \
	$(GO_INSTALL) || exit 1; \

clean:
	@echo "==> Clean $(PACKAGESRC) ..."; \
	$(GO_CLEAN) $(SOURCEDIR)/$(PACKAGESRC); \

doc:
	@echo "==> Generate documentation ..."; \
	$(RONN) $(MAN_DIR)/$(MAN_SRC) > /dev/null 2>&1 || exit 1; \
	mv $(MAN_DIR)/$(MAN_PAGE) $(SHARE_DIR)/; \
	mv $(MAN_DIR)/$(WEB_PAGE) $(SHARE_DIR)/; \
	cp $(SHARE_DIR)/$(MAN_PAGE) $(MAN_SYS_LOCAL)/; \

fmt:
	@echo "==> Format $(PACKAGESRC) ..."; \
	$(GO_FMT) $(SOURCEDIR)/$(PACKAGESRC) || exit 1; \

vet:
	@echo "==> Vet $(PACKAGESRC) ..."; \
	$(GO_VET) $(SOURCEDIR)/$(PACKAGESRC); \

lint:
	@echo "==> Lint $(PACKAGESRC) ..."; \
	$(GO_LINT) $(SOURCEDIR)/$(PACKAGESRC); \
