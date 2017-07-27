package app

import (
	"net/http"

	"github.com/bosemian/gogo/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	view.Index(w, nil)
}
