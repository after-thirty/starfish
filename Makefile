# Go parameters
GO           = go
GO_PATH      = $(shell $(GO) env GOPATH)
GO_OS 		 = $(shell $(GO) env GOOS)
ifeq ($(GO_OS), darwin)
    GO_OS = mac
endif

TIMEOUT_UNIT = 5m

# License environment
GO_LICENSE_CHECKER_DIR = license-header-checker-$(GO_OS)
GO_LICENSE_CHECKER = $(GO_PATH)/bin/license-header-checker
LICENSE_DIR = /tmp/tools/license


.PHONY: all
all: build-mac

.PHONY: build
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO) build -o starfish -ldflags "-X github.com/transaction-mesh/starfish/common/version.Version=1.0.0" -v ./cmd/...
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -o starfish -ldflags "-X github.com/transaction-mesh/starfish/common/version.Version=1.0.0" -v ./cmd/...
build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO) build -o starfish -ldflags "-X github.com/transaction-mesh/starfish/common/version.Version=1.0.0" -v ./cmd/...

.PHONY: test
test:
	$(GO) test -timeout $(TIMEOUT_UNIT) -v ./test/...

.PHONY: clean
clean:
	$(GO) clean
	@rm -rf test/tests.* test/coverage.*
	@rm -rf starfish

prepareLic:
	$(GO_LICENSE_CHECKER) -version || (wget https://github.com/lsm-dev/license-header-checker/releases/download/v1.2.0/$(GO_LICENSE_CHECKER_DIR).zip -O $(GO_LICENSE_CHECKER_DIR).zip && unzip -o $(GO_LICENSE_CHECKER_DIR).zip && mkdir -p $(GO_PATH)/bin/ && cp $(GO_LICENSE_CHECKER_DIR)/64bit/license-header-checker $(GO_PATH)/bin/)
	ls /tmp/tools/license/license.txt || wget -P $(LICENSE_DIR) https://github.com/dubbogo/resources/raw/master/tools/license/license.txt

.PHONY: license
license: prepareLic
	$(GO_LICENSE_CHECKER) -v -a -r -i vendor $(LICENSE_DIR)/license.txt . go