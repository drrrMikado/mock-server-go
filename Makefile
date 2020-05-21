# Go parameters
GO_CMD=go
CURRENT_DIR = $(shell pwd)
GO_BUILD=$(GO_CMD) build
GO_TEST=$(GO_CMD) test
GO_RUN=$(GO_CMD) run
MAIN_FILE=cmd/main.go
# added .gitignore
BINARY_NAME=cmd/mock-server
LOG_FILE=log/run.log

all: test build
build:
	@$(GO_BUILD) -o $(BINARY_NAME) -v $(MAIN_FILE)
test:
	@WORK_DIR=$(CURRENT_DIR) $(GO_TEST) -v ./...
run:
	@$(GO_RUN) $(MAIN_FILE)
run-bg: build
	$(BINARY_NAME=cmd/mock-server) &>>$(LOG_FILE) &