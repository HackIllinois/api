SOURCES = mkdocs.yml $(wildcard */*.md) 

.PHONY: install-deps serve build copy-docs-addendum

install-deps:
	pip3 install mkdocs mkdocs-material

# Starts a live-reloading server that serves the API docs
serve: $(SOURCES) install-deps copy-docs-addendum
	mkdocs serve

# Generates the static website
build: $(SOURCES) install-deps copy-docs-addendum
	mkdocs build

# Propagates contents of repository-level README.MD to developer guide
copy-docs-addendum: $(DOCS_ADDENDUM)
	cp -f ../README.md docs/getting-started/developer-guide.md
