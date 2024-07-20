package handlers

import (
	"encoding/json"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/ErikJermanis/erikjermanis.me/views/blog"
	"github.com/ErikJermanis/erikjermanis.me/views/notfound"
)

func Blog(w http.ResponseWriter, r *http.Request) {
  files, err := os.ReadDir("blogposts")
  if err != nil {
    slog.Error("error reading blog posts directory", "err", err)
    notfound.Index().Render(r.Context(), w)
    return
  }

  tag := strings.ReplaceAll(r.URL.Query().Get("tag"), "-", " ")

  var posts []blog.BlogPostMetadata
  var tags []string

  for _, file := range files {
    if strings.HasSuffix(file.Name(), ".json") {
      metaFile, err := os.ReadFile("blogposts/" + file.Name())
      if err != nil {
        slog.Error("error reading blog post metadata", "err", err)
        continue
      }

      var metadata blog.BlogPostMetadata
      if err := json.Unmarshal(metaFile, &metadata); err != nil {
        slog.Error("error parsing blog post metadata", "err", err)
        continue
      }

      if tag != "" && !slices.Contains(metadata.Tags, tag) {
        continue
      }

      for _, tag := range metadata.Tags {
        if !slices.Contains(tags, tag) {
          tags = append(tags, tag)
        }
      }

      posts = append(posts, metadata)
    }
  }

  sort.Sort(ByDate(posts))

  blog.Index(&posts, tags, tag).Render(r.Context(), w)
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

  var documentTitle = metadata.Title
  if len(documentTitle) > 40 {
    documentTitle = documentTitle[:37] + "..."
  }

  blog.BlogPost(documentTitle, content, &metadata).Render(r.Context(), w)
}
