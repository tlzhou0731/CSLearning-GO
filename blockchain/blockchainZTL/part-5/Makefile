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
	./$(BINARY) getbalance -address 1P1aPQZtYZJSqA7xpchkiHoYYgzrv6KVYT
	./$(BINARY) getbalance -address 14JysPEsDAKgTtyAX2xzj1pr6zXrdwXRY4
	./$(BINARY) getbalance -address 16mJyscx2dKoAe4zFiPLTAELhysehjk189
	./$(BINARY) getbalance -address 16Dv82iCU3AvGc3HCQVm9Zn8auLzEaAhBT
	./$(BINARY) getbalance -address 1C2QVXHEHstJ85YQLQVfVQn1dUdKf6tbaj
	./$(BINARY) getbalance -address 1AYuUQCNvJqPpyn3U4JcYXZiUvF7Mn2XVK
	./$(BINARY) getbalance -address 13DPc3uiLAQD9nsLDvMpwqmyKj8qi5RHU3
	./$(BINARY) getbalance -address 1MoVD4ioR7D5sFxj2TAGoFFEbnLLVZFA7s
	./$(BINARY) getbalance -address 1QHUeBaGekjEisfRzCaj9JLDzFzkTTNQmS
	./$(BINARY) listaddresses
	./$(BINARY) printchain
	./$(BINARY) send -from 14JysPEsDAKgTtyAX2xzj1pr6zXrdwXRY4 -to 1P1aPQZtYZJSqA7xpchkiHoYYgzrv6KVYT -amount 5
	./$(BINARY) printchain
	./$(BINARY) getbalance -address 14JysPEsDAKgTtyAX2xzj1pr6zXrdwXRY4
	./$(BINARY) getbalance -address 1P1aPQZtYZJSqA7xpchkiHoYYgzrv6KVYT
#	./$(BINARY) printchain

.PHONY: build deps test

#./blockchain createwallet
#Your new address: 13DPc3uiLAQD9nsLDvMpwqmyKj8qi5RHU3
#./blockchain createwallet
#Your new address: 1P1aPQZtYZJSqA7xpchkiHoYYgzrv6KVYT
#./blockchain createwallet
#Your new address: 16mJyscx2dKoAe4zFiPLTAELhysehjk189