APP_NAME=go-backend-api
BIN_DIR=bin
SRC_DIR=cmd/main.go

.PHONY: all build clean run

all: build

build: $(BIN_DIR)/$(APP_NAME)

$(BIN_DIR)/$(APP_NAME): $(SRC_DIR)
	mkdir -p $(BIN_DIR)
	go build -o $@ $<

run: build
	./$(BIN_DIR)/$(APP_NAME)

clean:
	rm -rf $(BIN_DIR)/*