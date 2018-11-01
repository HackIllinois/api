SOURCES = mkdocs.yml $(wildcard */*.md) 

.PHONY: install-deps serve build

install-deps:
	pip3 install mkdocs mkdocs-material
	
serve: $(SOURCES) install-deps
	mkdocs serve

build: $(SOURCES) install-deps
	cp -f ../README.md docs/getting-started/developer-guide.md
	cat docs/getting-started/docs_addendum.md >> docs/getting-started/developer-guide.md

	# Propagates contents of repository-level README.MD to developer guide
	mkdocs build
