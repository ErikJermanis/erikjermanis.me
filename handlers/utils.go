package handlers

import (
	"log/slog"
	"time"

	"github.com/ErikJermanis/erikjermanis.me/views/blog"
)

type ByDate []blog.BlogPostMetadata

func (a ByDate) Len() int {
	return len(a)
}

func (a ByDate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByDate) Less(i, j int) bool {
	dateFormat := "2006-01-02"
	dateI, errI := time.Parse(dateFormat, a[i].Date)
	dateJ, errJ := time.Parse(dateFormat, a[j].Date)
	if errI != nil || errJ != nil {
		slog.Error("error parsing date", "errI", errI, "errJ", errJ)
		return a[i].Date > a[j].Date
	}
	return dateI.After(dateJ)
}
