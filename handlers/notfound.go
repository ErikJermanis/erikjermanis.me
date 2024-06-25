package handlers

import (
	"net/http"

	"github.com/ErikJermanis/erikjermanis.me/views/notfound"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	notfound.Index().Render(r.Context(), w)
}