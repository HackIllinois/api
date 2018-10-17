SOURCES = mkdocs.yml $(wildcard */*.md) 

.PHONY: install-deps serve build

install-deps:
	pip3 install mkdocs mkdocs-material
	
serve: $(SOURCES)
	mkdocs serve

build: $(SOURCES)
	mkdocs build 