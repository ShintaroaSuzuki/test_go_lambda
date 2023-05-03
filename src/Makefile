SOURCES := $(wildcard _lambdas/*/main.go)
TARGETS := $(patsubst _lambdas/%/main.go, ../bin/lambdas/%, $(SOURCES))

../bin/lambdas/%: _lambdas/%/main.go
	GOARCH=amd64 GOOS=linux go build -o $@ $<

.PHONY: deps
deps:
	go mod download
	go mod tidy

.PHONY: clean
clean:
	rm -rf $(TARGETS)

.PHONY: bulid
build: deps $(TARGETS)

.PHONY: test
test:
	go test -v ./handlers/... ./services/... ./repositories/... -cover -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
