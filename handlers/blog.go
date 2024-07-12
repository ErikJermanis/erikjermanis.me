package handlers

import (
	"encoding/json"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/ErikJermanis/erikjermanis.me/views/blog"
	"github.com/ErikJermanis/erikjermanis.me/views/notfound"
)

func Blog(w http.ResponseWriter, r *http.Request) {
	blog.Index().Render(r.Context(), w)
}

func BlogPost(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	content, err := template.ParseFiles("blogposts/" + slug + ".html")
	if err != nil {
		slog.Error("error reading blog post content", "err", err)
		notfound.Index().Render(r.Context(), w)
		return
	}

	metadataFile, err := os.ReadFile("blogposts/" + slug + ".json")
	if err != nil {
		slog.Error("error reading blog post metadata", "err", err)
		notfound.Index().Render(r.Context(), w)
		return
	}

	var metadata blog.BlogPostMetadata
	if err := json.Unmarshal(metadataFile, &metadata); err != nil {
		slog.Error("error parsing blog post metadata", "err", err)
		notfound.Index().Render(r.Context(), w)
		return
	}

	blog.BlogPost(metadata.Title, content, &metadata).Render(r.Context(), w)
}
