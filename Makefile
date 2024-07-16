build:
	go build -ldflags="-s -w" -o ./bin/app main.go

run:
	go run main.go

lint:
	golangci-lint run --fast

fix:
	golangci-lint run --fix

install-hooks:
	@pre-commit install
