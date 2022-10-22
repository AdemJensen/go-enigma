# use mingw32-make.exe on windows please.
OUTPUT_DIR = ./bin

all: fmt build

run:
	go run ./

gen:
	go run ./gen

build:
	go build -o $(OUTPUT_DIR)/enigma
	go build -o $(OUTPUT_DIR)/gen ./gen

fmt:
	goimports -w ./
	go mod tidy

install_imports:
	go install golang.org/x/tools/cmd/goimports@latest