BASE_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

SOURCES = mkdocs.yml $(wildcard */*.md) 
DOCS_ADDENDUM = $(BASE_DIR)/../README.md
REQUIREMENTS = requirements.txt

.PHONY: install-deps serve build copy-docs-addendum

# Install dependencies
install-deps: $(REQUIREMENTS)
	pip3 install --user -r $(REQUIREMENTS)

# Starts a live-reloading server that serves the API docs
serve: $(SOURCES) install-deps copy-docs-addendum
	mkdocs serve

# Generates the static website
build: $(SOURCES) install-deps copy-docs-addendum
	mkdocs build

# Propagates contents of repository-level README.MD to developer guide
copy-docs-addendum: $(DOCS_ADDENDUM)
	mkdir -p $(BASE_DIR)/docs/getting-started/
	cp -f $(DOCS_ADDENDUM) $(BASE_DIR)/docs/getting-started/developer-guide.md
