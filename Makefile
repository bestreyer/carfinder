GOCMD						= go
GOBUILD						= $(GOCMD) build
GOCLEAN						= $(GOCMD) clean
GOINSTALL					= $(GOCMD) install
GOTEST						= $(GOCMD) test
GOGET						= $(GOCMD) get
WATCH						= realize start
GOBIN      					= $(shell go env GOPATH)/bin
DOCKER_COMPOSE				= docker-compose
DOCKER_COMPOSE_DEVELOPMENT 	= $(DOCKER_COMPOSE) -f build/docker-compose/development/docker-compose.yml
DOCKER_COMPOSE_LOADTEST		= $(DOCKER_COMPOSE) -f docker-compose-loadtest.yml

CARFINDER_CMD_SRC	= cmd/carfinder/main.go
CARFINDER_CMD_BIN	= build/carfinder

build:
	$(GOBUILD) -o $(CARFINDER_CMD_BIN) $(CARFINDER_CMD_SRC)

build_production:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -o $(CARFINDER_CMD_BIN) $(CARFINDER_CMD_SRC)

watch:
	$(WATCH)

development_up:
	$(DOCKER_COMPOSE_DEVELOPMENT) build
	$(DOCKER_COMPOSE_DEVELOPMENT) up -d
	docker cp scripts/init_tables.sh $(shell $(DOCKER_COMPOSE_DEVELOPMENT) ps -q postgre):/init_tables.sh
	$(DOCKER_COMPOSE_DEVELOPMENT) exec postgre bash /init_tables.sh

development_tests: development_up
	$(DOCKER_COMPOSE_DEVELOPMENT) exec application $(GOTEST) -v ./pkg/*
	$(DOCKER_COMPOSE_DEVELOPMENT) exec application $(GOTEST) -v ./integration

development_down:
	$(DOCKER_COMPOSE_DEVELOPMENT) down

loadtest:
	$(DOCKER_COMPOSE_LOADTEST) up -d

clean: $(CARFINDER_CMD_BIN)
	rm $(CARFINDER_CMD_BIN)

.PHONY: build watch clean development_up build_production development_down development_tests