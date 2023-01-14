#!/usr/bin/make -f

qface-gen:
	@echo "Generate qface"
	@mkdir -p gen
	@docker build -t qface_gen -f Dockerfile .
	@docker run --user $(shell id -u):$(shell id -g) --rm -v $(shell pwd):/workspace --workdir /workspace qface_gen sh ./generate.sh
	@gofmt -w ./gen/*
	@python3 format_cpp.py
	@./format.sh
.PHONY: proto-gen
