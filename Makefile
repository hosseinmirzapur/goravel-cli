docs:
	go run . --gen-docs

build: docs
	go build -o bin/goravel-cli

run: build
	./bin/goravel-cli --gen-docs

install: build
	go install