SOURCE_FILES?=./...
TEST_PATTERN?=.
TEST_OPTIONS?=-v

export PATH := ./bin:$(PATH)
export GO111MODULE := on

.PHONY: setup
setup: ## Install all the build and lint dependencies
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh
	go mod download

.PHONY: test
test: ## Run all the tests
	go test $(TEST_OPTIONS) -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.txt $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=2m

.PHONY: cover
cover: test ## Run all the tests and opens the coverage report
	go tool cover -html=coverage.txt

.PHONY: fmt
fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done
	# find . -name '*.md' -not -wholename './vendor/*' | xargs prettier --write

.PHONY: lint
lint: ## Run all the linters
	# TODO: fix tests and lll issues
	./bin/golangci-lint run ./...
	# find . -name '*.md' -not -wholename './vendor/*' | xargs prettier -l

.PHONY: ci
ci: build test lint ## Run all the tests and code checks

.PHONY: build
build: ## Build a beta version of rambler
	go build

.PHONY: todo
todo: ## Show to-do items per file.
	@grep \
		--exclude-dir=vendor \
		--exclude-dir=node_modules \
		--exclude=Makefile \
		--text \
		--color \
		-nRo -E ' TODO:.*|SkipNow' .

.PHONY: help
help: ## Get help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build