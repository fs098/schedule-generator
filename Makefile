run: build
	@./bin/schedule-generator

build:
	@go build -o bin/schedule-generator

clean:
	@rm -rf bin