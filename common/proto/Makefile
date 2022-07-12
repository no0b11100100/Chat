#!/usr/bin/make -f

proto-gen:
	@echo "Generate proto"
	@ docker build -t code_gen -f tools/Dockerfile.gen_proto .
	@ docker run --user $(shell id -u):$(shell id -g) --rm -v $(shell pwd):/workspace --workdir /workspace code_gen sh ./tools/protogen.sh
.PHONY: proto-gen
