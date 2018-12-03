BASE_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
SUBDIRS := $(wildcard $(BASE_DIR)/*/.)
SUBDIRS := $(BASE_DIR)/services/. $(BASE_DIR)/gateway/. $(BASE_DIR)/common/. $(BASE_DIR)/utilities/.

TARGETS := all test
SUBDIRS_TARGETS := $(foreach target,$(TARGETS),$(addsuffix $(target),$(SUBDIRS)))
DEPLOY_GATEWAY_TARGETS := gateway
DEPLOY_SERVICE_TARGETS := auth user registration decision rsvp checkin upload mail event stat notifications
DEPLOY_TARGETS := $(DEPLOY_GATEWAY_TARGETS) $(DEPLOY_SERVICE_TARGETS)

.PHONY: $(TARGETS) $(SUBDIRS_TARGETS)

$(TARGETS): % : $(addsuffix %,$(SUBDIRS))
	@echo 'Finished running target "$*"'

$(SUBDIRS_TARGETS):
	$(MAKE) -C $(@D) $(@F:.%=%)

.PHONY: run
run:
	@$(BASE_DIR)/scripts/run.sh

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: deploy
deploy:
	mkdir -p $(BASE_DIR)/build/
	$(foreach target,$(DEPLOY_TARGETS),mkdir -p $(BASE_DIR)/build/$(target)/;cp $(GOPATH)/bin/hackillinois-api-$(target) $(BASE_DIR)/build/$(target)/hackillinois-api-$(target);)
	$(foreach target,$(DEPLOY_GATEWAY_TARGETS),cp $(BASE_DIR)/$(target)/Dockerfile $(BASE_DIR)/build/$(target)/Dockerfile;)
	$(foreach target,$(DEPLOY_SERVICE_TARGETS),cp $(BASE_DIR)/services/$(target)/Dockerfile $(BASE_DIR)/build/$(target)/Dockerfile;)
	$(foreach target,$(DEPLOY_GATEWAY_TARGETS),cp $(BASE_DIR)/$(target)/buildspec.yml $(BASE_DIR)/build/$(target)/buildspec.yml;)
	$(foreach target,$(DEPLOY_SERVICE_TARGETS),cp $(BASE_DIR)/services/$(target)/buildspec.yml $(BASE_DIR)/build/$(target)/buildspec.yml;)
	mkdir -p $(BASE_DIR)/deploy/
	$(foreach target,$(DEPLOY_TARGETS),mkdir -p $(BASE_DIR)/deploy/api-$(target)/;cd $(BASE_DIR)/build/$(target);zip -r $(BASE_DIR)/deploy/api-$(target)/api-$(target).zip *; cd $(CURDIR);)
	rm -rf build/
	cp $(BASE_DIR)/config/production_config.json $(BASE_DIR)/deploy/config.json

.PHONY: release
release: all
	$(foreach target,$(DEPLOY_TARGETS),cp $(GOPATH)/bin/hackillinois-api-$(target) $(BASE_DIR)/release/hackillinois-api-$(target);)
	cd $(BASE_DIR)/release/ && docker build -t hackillinois-api:release . && cd $(CURDIR)
	docker save -o $(BASE_DIR)/release/hackillinois-api-image.tar hackillinois-api:release
	docker image rm -f hackillinois-api:release
	$(foreach target,$(DEPLOY_TARGETS),rm $(BASE_DIR)/release/hackillinois-api-$(target);)

.PHONY: setup
setup: all
	@echo 'Generating API admin token'
	hackillinois-utility-accountgen
	hackillinois-utility-tokengen

.PHONY: docs
docs:
	$(MAKE) -C $(BASE_DIR)/documentation build
