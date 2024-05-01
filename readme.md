## About

Based on 
- youtube https://www.youtube.com/watch?v=1dqp1s72Z8E&ab_channel=AnthonyGG
- git example source https://github.com/anthdm/gothstarter/tree/master

## Setup

### add `make` to windows git bash 
    - full guide https://gist.github.com/evanwill/0207876c3243bbb6863e65ec5dc3f058#make    
    - goto https://sourceforge.net/projects/ezwinports/files/
    - Download `make-4.1-2-without-guile-w32-bin.zip` (get the version without guile).
    - Extract zip.
    - Copy the contents to your `Git\mingw64\` merging the folders, but do NOT overwrite/replace any existing files.

### get `chi`
```
go get -u github.com/go-chi/chi/v5
```
### get `godotenv`
```
go get github.com/joho/godotenv
```
### get `templ`
- Guide https://templ.guide/quick-start/installation

For install
```
go install github.com/a-h/templ/cmd/templ@latest
```
Add to project
```
go get github.com/a-h/templ
```
To sync templ changes on top of air
```
templ generate --watch --proxy=http://localhost:3000
```

### get `air` for auto-reload
- Guide https://github.com/cosmtrek/air

Install go 1.22 and higher
```
go install github.com/cosmtrek/air@latest
```
run with command
```
air
```

### get `tailwindcss`
https://tailwindcss.com/docs/installation

1. Download stand alone executable for you platform
2. Rename it and put it somwhere to your path

```
# Create a tailwind.config.js file
./tailwindcss init

# Start a watcher example
./tailwindcss -i input.css -o output.css --watch

# Compile and minify your CSS for production example
./tailwindcss -i input.css -o output.css --minify
```

3. adjust `tailwind.config.js` created by `tailwindcss init` with

```
/** @type {import('tailwindcss').Config} */
module.exports = {
 	content: [ "./**/*.html", "./**/*.templ", "./**/*.go", ],
	safelist: [],
}
```

4. Add `app.css` to project

For example to `/views/css/app.css`
```
@tailwind base;
@tailwind components;
@tailwind utilities;
```

5. adjust `Makefile` add

```
css:
    tailwindcss -i views/css/app.css -o public/styles.css --watch
```

Path `public/styles.css` must link to what is set in `base.templ`

```
    <link rel="stylesheet" href="/public/styles.css"/>
```

### Opensources components

#### Franken
https://github.com/sveltecult/franken-ui

Franken UI is an HTML-first, open-source library of UI components based on the utility-first Tailwind CSS with UIkit 3 compatibility. The design is based on shadcn/ui ported to be framework-agnostic.

#### Alpine
https://devdojo.com/pines/docs/how-to-use

