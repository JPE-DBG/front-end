package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()

	router.Get("/foo", handlerFoo)

	listnerAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listnerAddr", listnerAddr)
	http.ListenAndServe(listnerAddr, router)
}

func handlerFoo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))
}
