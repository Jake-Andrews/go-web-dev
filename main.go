package main

import (
	//"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Jake-Andrews/go-web-dev/handler"
	"github.com/Jake-Andrews/go-web-dev/pkg/supabase"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()
	router.Handle("/*", Publicfs())
	router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))
	router.Get("/signin", handler.MakeHandler(handler.HandleSigninIndex))

	port := os.Getenv("LISTEN_ADDR")
	slog.Info("Running on: ", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func initEverything() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return supabase.Init()
}
