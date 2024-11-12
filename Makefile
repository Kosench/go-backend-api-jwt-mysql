APP_NAME=go-backend-api
BIN_DIR=bin
SRC_DIR=cmd/main.go
MIGRATION_DIR=cmd/migrate/migrations

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

test:
	@go test -v ./...

migration:
	@migrate create -ext sql -dir $(MIGRATION_DIR) $(filter-out $@,$(MAKECMDGOALS))

# Применение миграций (up)
migrate-up:
	@go run cmd/migrate/main.go up -path $(MIGRATION_DIR)

# Откат миграций (down)
migrate-down:
	@go run cmd/migrate/main.go down -path $(MIGRATION_DIR)