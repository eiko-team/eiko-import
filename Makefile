BIN = eiko-off

all: clean
all: build
all: exec

build:
	go build -o $(BIN)

exec:
	./$(BIN)

clean:
	rm -fr $(BIN)