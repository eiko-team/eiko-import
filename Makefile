BIN = eiko-import

all: clean
all: build
all: exec

build:
	go build -o $(BIN)

exec:
	./$(BIN)

clean:
	rm -fr $(BIN)

test:
	go test $(ARGS) ./...
