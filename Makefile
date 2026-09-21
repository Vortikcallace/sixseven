APP_NAME := sixseven
BIN_DIR  := bin

.PHONY: build run clean

build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) .

run: build
	./$(BIN_DIR)/$(APP_NAME)

clean:
	rm -rf $(BIN_DIR)
