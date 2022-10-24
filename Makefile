.PHONY: build test test-integration

build:
	go build -o conduit-connector-connectorName cmd/connector/main.go

test:
	go test $(GOTEST_FLAGS) -v -race ./...

test-integration:
	# run required docker containers, execute integration tests, stop containers after tests
	docker compose -f test/docker-compose.yml up -d
	go test $(GOTEST_FLAGS) -v -race ./...; ret=$$?; \
		docker compose -f test/docker-compose.yml down; \
		exit $$ret
