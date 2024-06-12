package main

import (
	"fmt"
	"net/http"
	"os"
)

func Publicfs() http.Handler {
	fmt.Println("building static files for development")
	return http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public")))
}
