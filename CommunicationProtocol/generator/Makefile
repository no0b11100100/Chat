#!/usr/bin/make -f

qface-gen:
	@echo "Generate qface"
	@ docker build -t qface_gen -f Dockerfile .
	@ docker run --user $(shell id -u):$(shell id -g) --rm -v $(shell pwd):/workspace --workdir /workspace qface_gen sh ./generate.sh
.PHONY: proto-gen
