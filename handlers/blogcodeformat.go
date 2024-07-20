package handlers

import (
	"bytes"
	"net/http"
	"os"
	"strings"

	"github.com/ErikJermanis/erikjermanis.me/views/blogcodeformat"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/quick"
)

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != os.Getenv("HTTP_USER") || password != os.Getenv("HTTP_PASS") {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func BlogCodeFormat(w http.ResponseWriter, r *http.Request) {

	lexers := lexers.Names(false)

	blogcodeformat.Index("", lexers).Render(r.Context(), w)
}

func HandleFormatCode(w http.ResponseWriter, r *http.Request) {
	lexer := r.FormValue("lexer")
	code := r.FormValue("code")

	buf := new(bytes.Buffer)
	err := quick.Highlight(buf, code, lexer, "html", "github-dark")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	html := buf.String()
	startIdx := strings.Index(html, "<pre")
	if startIdx == -1 {
		startIdx = 0
	}
	endIdx := strings.Index(html, "</body>")
	if endIdx == -1 {
		endIdx = len(html)
	}
	html = html[startIdx:endIdx]

	lexers := lexers.Names(false)

	blogcodeformat.Index(html, lexers).Render(r.Context(), w)
}
