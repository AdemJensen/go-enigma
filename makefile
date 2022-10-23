# use mingw32-make.exe on windows please.
OUTPUT_DIR = ./bin

all: fmt build

run:
	go run ./cmd/test_cli

gen:
	go run ./cmd/gen

build:
	go build -o $(OUTPUT_DIR)/enigma ./cmd/test_cli
	go build -o $(OUTPUT_DIR)/gen ./cmd/gen

fmt:
	goimports -w ./
	go mod tidy

install_imports:
	go install golang.org/x/tools/cmd/goimports@latest