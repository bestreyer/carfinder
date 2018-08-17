GOCMD				= go
GOBUILD				= $(GOCMD) build
GOCLEAN				= $(GOCMD) clean
GOINSTALL			= $(GOCMD) install
GOTEST				= $(GOCMD) test
GOGET				= $(GOCMD) get
WATCH				= realize start
GOBIN      			= $(shell go env GOPATH)/bin
CARFINDER_CMD_SRC	= cmd/carfinder/main.go
CARFINDER_CMD_BIN	= build/carfinder


build:
	$(GOBUILD) -o $(CARFINDER_CMD_BIN) $(CARFINDER_CMD_SRC)

watch:
	$(WATCH)

clean: $(CARFINDER_CMD_BIN)
	rm $(CARFINDER_CMD_BIN)

.PHONY: build watch clean