dep:
	go mod download

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/keto-granola .

run:
	go run main.go

lint: lint/install lint/run

lint/install:
	curl -sSfL https://golangci-lint.run/install.sh | sh -s -- v2.12.2

lint/run:
	bin/golangci-lint run --config .golangci.yml

lint/fix:
	bin/golangci-lint run --config .golangci.yml --fix

test: # runs all tests
	go test -shuffle=on -tags=e2e ./...

test/unit: # unit tests only
	go test -shuffle=on ./...

test/e2e: # e2e tests only
	go test -shuffle=on -tags=e2e ./...

sqlc:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.31.1 generate

mocks:
	go install github.com/matryer/moq@latest && \
	go generate ./...

migrate/create:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make migrate/create name=<migration_name>"; \
		exit 1; \
	fi
	migrate create -ext sql -dir internal/store/db/migrations -seq $(name)