GOFILES=$(shell find . -name '*.go' | egrep -v '^\./vendor')

all: build

build: $(GOFILES)
	go build -o bin/apitraining
