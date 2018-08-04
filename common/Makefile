ifndef GOPATH
$(error GOPATH not set, aborting build)
endif

BASE_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
BUILD_PACKAGE:=$(subst $(GOPATH)/src/,,$(BASE_DIR))

.PHONY: all
all:
	@:

.PHONY: test
test:
	go test $(BUILD_PACKAGE)/tests
