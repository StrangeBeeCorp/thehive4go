GO := go
GO_IMAGE := golang:1.24.7-alpine
OPENAPI_GENERATOR_IMAGE := openapitools/openapi-generator-cli:v7.14.0
BGreen="\033[1;32m"       # Green
Color_Off="\033[0m"       # Text Reset

.PHONY: all
all: fmt security test generate ## Format, run security checks, test, and generate client

.PHONY: help
help: ## Display this help screen
	@echo $(BGreen)--------------$(Color_Off)
	@echo $(BGreen)-- Help :-\) --$(Color_Off)
	@echo $(BGreen)--------------$(Color_Off)
	@grep -E '^[0-9a-z.A-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1;36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: fmt
fmt: ## Format the code
	@echo $(BGreen)-------------$(Color_Off)
	@echo $(BGreen)--- Format --$(Color_Off)
	@echo $(BGreen)-------------$(Color_Off)
	docker run -i --rm -v $(CURDIR):/app -w /app $(GO_IMAGE) go fmt ./...
	@echo "Code formatted"

.PHONY: security
security: vulncheck vetlint ## Run security checks

.PHONY: test
test: ## Run tests with coverage
	@echo $(BGreen)-----------------------$(Color_Off)
	@echo $(BGreen)-- Running UnitTests --$(Color_Off)
	@echo $(BGreen)-----------------------$(Color_Off)
	docker run -i --rm -v $(CURDIR):/app -w /app $(GO_IMAGE) go test -v ./...

.PHONY: vulncheck
vulncheck: ## Check for vulnerabilities
	@echo $(BGreen)------------------------------$(Color_Off)
	@echo $(BGreen)-- Security Vulnerability  --$(Color_Off)
	@echo $(BGreen)------------------------------$(Color_Off)
	docker run -i --rm -v $(CURDIR):/app -w /app $(GO_IMAGE) sh -c 'go install golang.org/x/vuln/cmd/govulncheck@latest && govulncheck ./...'

.PHONY: vetlint
vetlint: ## Run linter checks
	@echo $(BGreen)-----------------------------$(Color_Off)
	@echo $(BGreen)-- Linter Checks --$(Color_Off)
	@echo $(BGreen)-----------------------------$(Color_Off)
	docker run -v $(CURDIR):/app -w /app -i --rm golangci/golangci-lint:latest golangci-lint run -v ./...

.PHONY: updatedep
updatedep: ## Update dependencies
	@echo $(BGreen)-----------------------$(Color_Off)
	@echo $(BGreen)-- Update Dependencies --$(Color_Off)
	@echo $(BGreen)-----------------------$(Color_Off)
	docker run -i --rm -v $(CURDIR):/app -w /app $(GO_IMAGE) sh -c 'go get -u ./... && go mod tidy'

.PHONY: generate
generate: ## Generate the client code from OpenAPI spec
	@echo $(BGreen)-----------------------$(Color_Off)
	@echo $(BGreen)-- Generating Client  --$(Color_Off)
	@echo $(BGreen)-----------------------$(Color_Off)
	@docker build -f Dockerfile.generator -t thehive4go-generator .
	@docker run --rm -v $(CURDIR):/workspace -u $(shell id -u):$(shell id -g) thehive4go-generator

.PHONY: clean
clean: ## Remove build artifacts and Docker images
	@echo $(BGreen)-------------------$(Color_Off)
	@echo $(BGreen)-- Cleaning up   --$(Color_Off)
	@echo $(BGreen)-------------------$(Color_Off)
	@rm -rf tmp/
	@docker rmi thehive4go-generator 2>/dev/null || true

.PHONY: integration-test
integration-test: ## Run full integration tests with TheHive stack
	@echo $(BGreen)---------------------------$(Color_Off)
	@echo $(BGreen)-- Running Integration  --$(Color_Off)
	@echo $(BGreen)-- Tests with Full Stack--$(Color_Off)
	@echo $(BGreen)---------------------------$(Color_Off)
	cd integration && docker-compose up --abort-on-container-exit integration-tests
	cd integration && docker-compose down
