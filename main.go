package main

import (
	handlers "gothstarter/handlets"
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

	router.Handle("/*", public())
	router.Get("/foo2", handlers.Make(handlers.HandlerFoo))

	listnerAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listnerAddr", listnerAddr)
	http.ListenAndServe(listnerAddr, router)
}
