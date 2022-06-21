BINARY := blockchain

all: build test

build: deps
	@echo "====> Go build"
	@go build -o $(BINARY)

deps:
	@go get -u github.com/boltdb/bolt

test:
#	./$(BINARY) createwallet
#	./$(BINARY) createwallet
#	./$(BINARY) createwallet
	./$(BINARY) listaddresses
#	./$(BINARY) createblockchain -address 16PgKKB8keAqgjS9f168GDPHWZbTRaZE8c

	./$(BINARY) printchain
	./$(BINARY) getbalance -address 16PgKKB8keAqgjS9f168GDPHWZbTRaZE8c
	./$(BINARY) send -from 16PgKKB8keAqgjS9f168GDPHWZbTRaZE8c -to 1ARe2doGnG8ycMHaQx1Snh8wnTj9zj4kes -amount 2
	./$(BINARY) getbalance -address 1ARe2doGnG8ycMHaQx1Snh8wnTj9zj4kes


#	./$(BINARY) printUsage



#	./$(BINARY) printchain

.PHONY: build deps test




