build:
	go build -o bin/goravel-cli

run: build
	./bin/goravel-cli

install: build
	go install