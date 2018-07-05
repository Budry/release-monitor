BUILD_DIR := ./dist
APP_NAME := release-monitor
SOURCE_DIR := ./src
ENTRYPOINT := main.go

all: prepare test build clean

prepare:
	go get -v ./src

run:
	go run $(SOURCE_DIR)/main.go ${ARG}

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) $(SOURCE_DIR)/main.go

test:
	go test ${SOURCE_DIR}/...

clean:
	go clean