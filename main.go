package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/ErikJermanis/erikjermanis.me/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("ENV")
	if env != "local" {
		if err := godotenv.Load(); err != nil {
			slog.Error("error loading .env", "err", err)
		}
	}

	r := chi.NewRouter()

	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	r.Get("/", handlers.Index)

	port := os.Getenv("HTTP_PORT")

	slog.Info(fmt.Sprintf("Server is running on port %s", port))
	http.ListenAndServe(port, r)
}