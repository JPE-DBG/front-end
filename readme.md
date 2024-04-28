## About

Based on 
- youtube https://www.youtube.com/watch?v=1dqp1s72Z8E&ab_channel=AnthonyGG
- git example source https://github.com/anthdm/gothstarter/tree/master

## Setup

- add `make` to windows git bash 
    - full guide https://gist.github.com/evanwill/0207876c3243bbb6863e65ec5dc3f058#make    
    - goto https://sourceforge.net/projects/ezwinports/files/
    - Download `make-4.1-2-without-guile-w32-bin.zip` (get the version without guile).
    - Extract zip.
    - Copy the contents to your `Git\mingw64\` merging the folders, but do NOT overwrite/replace any existing files.

- get `chi`
```
go get -u github.com/go-chi/chi/v5
```
- get `godotenv`
```
go get github.com/joho/godotenv
```
- get `templ`
    - guide https://templ.guide/quick-start/installation
```
For install
go install github.com/a-h/templ/cmd/templ@latest

Add to project
go get github.com/a-h/templ
```

- get `air` for auto-reload
    - https://github.com/cosmtrek/air
```
Install go 1.22 and higher
go install github.com/cosmtrek/air@latest

run with command
air
```

To sync templ changes on top of air
```
templ generate --watch --proxy=http://localhost:3000
```