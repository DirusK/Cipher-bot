.SILENT:
.EXPORT_ALL_VARIABLES:
.PHONY: all test build run run-build clean rebuild

BUILD_NAME=racer-bot

all: run

lint:
	golangci-lint run -c golangci-lint.yml

test:
	go clean -testcache ./...
	go test -v -p 4 ./...

check: test lint

gen:
	go generate ./...

build: test
	go build -o ./build/$(BUILD_NAME) ./cmd/.

run:
	go run cmd/main.go

run-build: build
	cd cmd; ./$(BUILD_NAME)

run-docker: build
	docker build -t $(BUILD_NAME) . && docker run -it $(BUILD_NAME)

clean:
	go clean ./...
	rm -f ./cmd/$(BUILD_NAME)

rebuild: clean build

vendor:
	go mod vendor

mod:
	go mod tidy

# ubuntu
# install sudo snap install goimports-reviser --devmode

# mac
# brew tap incu6us/homebrew-tap
# brew install incu6us/homebrew-tap/goimports-reviser

imports:
	$(shell find . -name \*.go -and -not -name '**mock*.go' -and -not -name 'validate.go' -exec sh -c 'goimports -w {}; goimports-reviser -project-name ${BUILD_NAME} -file-path {}' \;)
