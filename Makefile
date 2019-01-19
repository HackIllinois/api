BASE_PACKAGE := github.com/HackIllinois/api
REPO_ROOT := $(shell git rev-parse --show-toplevel)

SERVICES := auth user registration decision rsvp checkin upload mail event stat notifications
GATEWAYS := gateway common

UTILITIES := accountgen tokengen

.PHONY: all
all: api utilities

.PHONY: api
api:
	@echo 'Building api'
	@mkdir -p $(REPO_ROOT)/bin
	@go build -o $(REPO_ROOT)/bin/hackillinois-api $(BASE_PACKAGE)

.PHONY: test
test:
	@echo 'Testing services'
	@$(foreach service,$(SERVICES),HI_CONFIG=file://$(REPO_ROOT)/config/test_config.json go test $(BASE_PACKAGE)/services/$(service)/tests;)
	@echo 'Testing gateway'
	@$(foreach gateway,$(GATEWAYS),HI_CONFIG=file://$(REPO_ROOT)/config/test_config.json go test $(BASE_PACKAGE)/$(gateway)/tests;)

.PHONY: utilities
utilities:
	@echo 'Building utilities'
	@mkdir -p $(REPO_ROOT)/bin
	@$(foreach utility,$(UTILITIES),go build -o $(REPO_ROOT)/bin/hackillinois-utility-$(utility) $(BASE_PACKAGE)/utilities/$(utility);)

.PHONY: setup
setup: all
	@echo 'Generating API admin account'
	@export HI_CONFIG=file://$(REPO_ROOT)/config/dev_config.json; \
	$(REPO_ROOT)/bin/hackillinois-utility-accountgen
	@echo 'Generating token for admin account'
	@export HI_CONFIG=file://$(REPO_ROOT)/config/dev_config.json; \
	$(REPO_ROOT)/bin/hackillinois-utility-tokengen

.PHONY: run
run:
	@$(REPO_ROOT)/scripts/run.sh

.PHONY: fmt
fmt:
	@gofmt -s -w -l .

.PHONY: container
container: api
	@echo 'Builing API container'
	@mkdir -p $(REPO_ROOT)/build
	@cp $(REPO_ROOT)/bin/hackillinois-api $(REPO_ROOT)/build/hackillinois-api
	@cp $(REPO_ROOT)/container/Dockerfile $(REPO_ROOT)/build/Dockerfile
	@docker build -t hackillinois-api:latest $(REPO_ROOT)/build
	@rm -rf $(REPO_ROOT)/build

.PHONY: release
release: api
	@echo 'Building release container'
	@mkdir -p $(REPO_ROOT)/build
	@cp $(REPO_ROOT)/bin/hackillinois-api $(REPO_ROOT)/build/hackillinois-api
	@cp $(REPO_ROOT)/release/Dockerfile $(REPO_ROOT)/build/Dockerfile
	@cp $(REPO_ROOT)/release/start.sh $(REPO_ROOT)/build/start.sh
	@docker build -t hackillinois-api:release $(REPO_ROOT)/build
	@docker save -o $(REPO_ROOT)/bin/hackillinois-api-image.tar hackillinois-api:release
	@docker image rm -f hackillinois-api:release
	@rm -rf $(REPO_ROOT)/build

.PHONY: docs
docs:
	$(MAKE) -C $(BASE_DIR)/documentation build
