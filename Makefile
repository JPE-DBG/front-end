run: build
	@./bin/app.exe

build:
	@go build -o bin/app.exe cmd/app/main.go