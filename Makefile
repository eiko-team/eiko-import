BIN = eiko-import
CONFIG = $(PWD)/config.json

all: clean
all: build
all: exec

build:
	go build -o $(BIN)

exec:
	CONFIG=$(CONFIG) ./$(BIN)

clean:
	rm -fr $(BIN)

test:
	go test $(ARGS) ./...
