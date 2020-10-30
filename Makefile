SHELL=/bin/bash
BINARY_NAME:="devops-terraformenv-fix"
GOPATH:="${HOME}/go"

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@go get $(get)

.PHONY: install
install: ## Install the binary
	@go build -o ${GOPATH}/bin/${BINARY_NAME}