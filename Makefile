run: build
	@./bin/app.exe

build:
	@go build -o bin/app.exe .