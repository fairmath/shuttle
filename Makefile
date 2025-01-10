.PHONY: view test lint generate clean all release
.DEFAULT_GOAL := all

SWAGGER_API = internal/client/tendermint/api/internal/txs \
	internal/client/tendermint/api/internal/blocks \
	internal/client/tendermint/api/internal/bank

.PHONY: view
view:
	go tool cover -html=coverage.txt 

.PHONY: test
test: 
	go test -v -coverprofile coverage.txt ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags="-extldflags=-static" -o ./bin/shuttle ./cmd/

force: ;

internal/client/tendermint/api/%: force
	swagger generate client -f $@/api.swagger.yaml -t $@

mocks: 
	go generate ./...

.PHONY: generate
generate: $(SWAGGER_API) mocks

.PHONY: clean
clean:
	git clean -fd

.PHONY: all
all: generate lint test build

.PHONY: release 
release: $(IN_DOCKER_TARGET)
	docker build -f ./Dockerfile.release -t yashalabinc/shuttle:local .

fail:
	@echo "This target should be run outside of the container"
