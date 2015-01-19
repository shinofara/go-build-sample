

default: all
all: build
build: 
	go build -ldflags "-X main.version $(git describe)"
