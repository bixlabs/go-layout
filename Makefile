all: deps lint

.PHONY: test clean format lint coverage coverage-html build build-for-mac build-for-windows

deps:
		./deps.sh

test:
		go test -v -covermode=count -coverprofile=coverage.out ./...

coverage:
		go test -covermode=count -coverprofile=coverage.out ./...

coverage-html:
		make coverage && go tool cover -html=coverage.out

format:
		go vet ./... && go fmt ./...

build:
		make api-docs && make format && go build -o ./tmp/auth-server ./api/main.go

build-for-mac:
		GOOS=darwin GOARCH=amd64 make build

build-for-windows:
		GOOS=windows GOARCH=386 make api-docs && make format && go build -o ./tmp/auth-server.exe ./api/main.go

clean:
		rm -r -f ./tmp

lint:
		golangci-lint run

run-dev:
		make format && air -c .air.config

run:
		make api-docs && make format && go run api/main.go

run-cli:
		make format && go run cmd/cli/main.go

api-docs:
		swag init -g api/main.go

ci:
		make all build
