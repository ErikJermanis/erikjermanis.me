package handlers

import (
	"net/http"

	"github.com/ErikJermanis/erikjermanis.me/views/index"
)

func Index(w http.ResponseWriter, r *http.Request) {
	index.Index().Render(r.Context(), w)
}