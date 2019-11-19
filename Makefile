BIN = eiko-import
CONFIG = $(PWD)/toto.json

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
