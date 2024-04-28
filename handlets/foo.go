package handlers

import (
	"gothstarter/views/foo"
	"net/http"
)

func HandlerFoo(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, foo.Index())
}
