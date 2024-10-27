BINARY=./bin/ImageResizer

run: build
	@$(BINARY)


build:
	@mkdir -p bin
	@go build -o $(BINARY) cmd/api/main.go

clean:
	@rm -f $(BINARY)

