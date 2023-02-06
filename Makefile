# This makefile handles build, testing, and deploying all services in the HackIllinois API
# It is designed to be simple to understand and easy to add a new service by following the instructions in the comments

# BASE_PACKAGE should be the name of the go module
# REPO_ROOT will be used to build absolute paths during build or testing stages
BASE_PACKAGE := github.com/HackIllinois/api
REPO_ROOT := $(shell git rev-parse --show-toplevel)

# SERVICES is the list services to test, services are located at $(REPO_ROOT)/service/<service_name>
# GATEWAYS is the list of top level directories to test, gateways are located at $(REPO_ROOT)/<gateway_name>
# Add new services or top level directories to test here
SERVICES := auth user registration decision rsvp checkin upload mail event stat notifications project profile
GATEWAYS := gateway common

# UTILITIES is the list of utilities to build, utilities are located at $(REPO_ROOT)/utilities/<utility_name>
# Add new utilities to build here
UTILITIES := accountgen tokengen

# TAG is used to tag the docker containers being built
TAG := latest
ifeq ($(strip $(TAG)),)
	override TAG := latest
endif

# Builds the API binary and all utilities
.PHONY: all
all: api utilities

# Builds the API binary
.PHONY: api
api:
	@echo 'Building api'
	@mkdir -p $(REPO_ROOT)/bin
	@go build -o $(REPO_ROOT)/bin/hackillinois-api $(BASE_PACKAGE)

# Tests all services and gateways
.PHONY: test
test:
	@echo 'Testing services'
	@$(foreach service,$(SERVICES),HI_CONFIG=file://$(REPO_ROOT)/config/test_config.json go test $(BASE_PACKAGE)/services/$(service)/tests || exit 1;)
	@echo 'Testing gateway'
	@$(foreach gateway,$(GATEWAYS),HI_CONFIG=file://$(REPO_ROOT)/config/test_config.json go test $(BASE_PACKAGE)/$(gateway)/tests || exit 1;)

# Builds all utilities
.PHONY: utilities
utilities:
	@echo 'Building utilities'
	@mkdir -p $(REPO_ROOT)/bin
	@$(foreach utility,$(UTILITIES),go build -o $(REPO_ROOT)/bin/hackillinois-utility-$(utility) $(BASE_PACKAGE)/utilities/$(utility);)

# Builds the API binary, all utilities, and then sets up an admin account
.PHONY: setup
setup: all
	@echo 'Generating API admin account'
	@export HI_CONFIG=file://$(REPO_ROOT)/config/dev_config.json; \
	$(REPO_ROOT)/bin/hackillinois-utility-accountgen
	@echo 'Generating token for admin account'
	@export HI_CONFIG=file://$(REPO_ROOT)/config/dev_config.json; \
	$(REPO_ROOT)/bin/hackillinois-utility-tokengen

# Runs the API with each service in a seperate process
.PHONY: run
run:
	@$(REPO_ROOT)/scripts/run.sh

# Runs the API with each service in a separate process using the test config
.PHONY: run-test
run-test:
	@$(REPO_ROOT)/scripts/run-test.sh

# Runs the API with all services in a single process
.PHONY: run-single
run-single:
	@$(REPO_ROOT)/scripts/run-single.sh

# Formats the repo
.PHONY: fmt
fmt:
	@gofmt -s -w -l .

# Builds a docker container with the API binary
.PHONY: container
container:
	@echo 'Builing API container'
	@mkdir -p $(REPO_ROOT)/build
	@cp $(REPO_ROOT)/bin/hackillinois-api $(REPO_ROOT)/build/hackillinois-api
	@cp $(REPO_ROOT)/container/Dockerfile $(REPO_ROOT)/build/Dockerfile
	@docker build -t hackillinois-api:$(TAG) $(REPO_ROOT)/build
	@rm -rf $(REPO_ROOT)/build

# Generates a tar archive of the API container
.PHONY: release
release: container
	@echo 'Building API container release'
	@docker save -o $(REPO_ROOT)/bin/hackillinois-api-image.tar hackillinois-api:$(TAG)
	@rm -rf $(REPO_ROOT)/build

# Pushes the API container to DockerHub
.PHONY: container-push
container-push:
	@echo 'Pushing container to DockerHub'
	@echo "$(DOCKER_PASSWORD)" | docker login -u "$(DOCKER_USERNAME)" --password-stdin
	@docker tag hackillinois-api:$(TAG) hackillinois/api:$(TAG)
	@docker push hackillinois/api:$(TAG)

# Builds static html documentation for the API
.PHONY: docs
docs:
	$(MAKE) -C $(REPO_ROOT)/documentation build

# Starts up the API on the test config and runs the E2E integration tests
.PHONY: integration-test
integration-test:
	@$(REPO_ROOT)/scripts/run-integration.sh "$(API_OUTPUT)" "$(TEST_DIR)" "$(RUN)" || exit 1;
