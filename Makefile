# Go parameters
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
PROGRAM_NAME=gLC3VM

all: test build
build:
				$(GO_BUILD) -o $(PROGRAM_NAME) -v
test:
				$(GO_TEST) -v ./...
clean:
				$(GO_CLEAN)
				rm -f $(PROGRAM_NAME)

run:
				$(GO_BUILD) -o $(PROGRAM_NAME) -v ./...
				./$(PROGRAM_NAME) demo/2048.obj
