package handlers

import (
	"gothstarter/views/home"
	"net/http"
)

func HandlerHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}
