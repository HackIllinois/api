ifndef GOPATH
$(error GOPATH not set, aborting build)
endif

ifndef BASE_DIR
$(error BASE_DIR not set, aborting build)
endif

BUILD_PACKAGE:=$(subst $(GOPATH)/src/,,$(BASE_DIR))
TARGET_NAME:=$(notdir $(BUILD_PACKAGE:%/=%))

.PHONY: all
all:
	go build -i -o $(GOPATH)/bin/hackillinois-utility-$(TARGET_NAME) $(BUILD_PACKAGE) 
