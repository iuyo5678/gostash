GOPATH := $(shell pwd)
.PHONY: clean test

all:
	@GOPATH=$(GOPATH) go install gostash

clean:
	@rm -fr bin pkg

test:
	@GOPATH=$(GOPATH) go test gostash
