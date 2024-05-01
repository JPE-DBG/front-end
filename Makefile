run: build
	@./bin/app.exe

build:
	@go build -o bin/app.exe .

css:
	tailwindcss -i views/css/app.css -o public/styles.css --watch