ENVFLAGS = GO111MODULE=on CGO_ENABLED=0 GOOS=$(shell go env GOOS) GOARCH=$(shell go env GOARCH)

help:
	@echo "usage:"
	@echo "make test   -- run tests"
	@echo "make lint   -- run go linter in all code"
	@echo "make build  -- build the API"
	@echo "make run    -- run the API"


test:
	@echo "Running tests"
	go test ./... 

lint:
	go vet ./...

build:
	echo "Building the API..." ; \
	if [ -z "$(TAG)" ]; then \
		$(ENVFLAGS) go build -a -tags netgo -ldflags="-w -s -extldflags \"-static\" -X github.com/hosting-engine/envoy-proxy-manager-api/pkg/utils.version=dev-build-$$(date +%s)" -o bin/epm-api ./cmd/api ; \
	else \
		$(ENVFLAGS) go build -a -tags netgo -ldflags="-w -s -extldflags \"-static\" -X github.com/hosting-engine/envoy-proxy-manager-api/pkg/utils.version=$$TAG" -o bin/epm-api ./cmd/api ; \
	fi ; \
	echo "API is built in bin/epm-api"

run:
	@echo "Running API"
	./bin/epm-api
	@echo "API is not running anymore"
