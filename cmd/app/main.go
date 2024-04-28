package main

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewMux()
	router.Get("/foo", handlerFoo)

	listnerAddr := ":3000"
	slog.Info("HTTP server started", "listnerAddr", listnerAddr)
	http.ListenAndServe(listnerAddr, router)
}

func handlerFoo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))
}
