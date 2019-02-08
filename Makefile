.PHONY: all test build

all: clean test build

test:
		go test -v -covermode=count -coverprofile=coverage.out ./...

coverage:
		go test -covermode=count -coverprofile=coverage.out ./...

coverage-html:
		make coverage && go tool cover -html=coverage.out

build:
		go build -o ./tmp/web-server ./cmd/api/main.go

clean:
		rm -r -f ./tmp

lint:
		golangci-lint run

run-dev:
		~/.air -c .air.config

run:
		$ go run cmd/api/main.go

deps:
		sh ./scripts/install_dep.sh
		sh ./scripts/install_air.sh
		sh ./scripts/install_golangci_lint.sh
		dep ensure
