BIN=gol

build:
	GOARCH=amd64 GOOS=linux go build -o $(BIN)-linux main.go

run:
	./$(BIN)

build_and_run: build run

clean:
	go clean
	rm $(BIN)-linux
