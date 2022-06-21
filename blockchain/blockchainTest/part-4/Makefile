BINARY := blockchain.exe

all: build test

build: deps
	@echo "====> Go build"
	@go build -o $(BINARY)

deps:
	@go get -u github.com/boltdb/bolt

test:
	./$(BINARY) createblockchain -address "Ivan"
	./$(BINARY) printchain
	./$(BINARY) send -from Ivan -to Pedro -amount 6
	./$(BINARY) send -from Ivan -to TLL -amount 2
	./$(BINARY) printchain

.PHONY: build deps test
