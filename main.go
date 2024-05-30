package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

    "github.com/Jake-Andrews/go-web-dev/handler"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()
    router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))

    port := os.Getenv("LISTEN_ADDR")
	slog.Info("Running on: ", "port", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}

}

func initEverything() error {
	return godotenv.Load()
}
