SOURCES = mkdocs.yml $(wildcard */*.md) 
DOCS_ADDENDUM = ../README.md

.PHONY: serve build copy-docs-addendum

# Starts a live-reloading server that serves the API docs
serve: $(SOURCES) copy-docs-addendum
	mkdocs serve

# Generates the static website
build: $(SOURCES) copy-docs-addendum
	mkdocs build

# Propagates contents of repository-level README.MD to developer guide
copy-docs-addendum: $(DOCS_ADDENDUM)
	cp -f ../README.md docs/getting-started/developer-guide.md
