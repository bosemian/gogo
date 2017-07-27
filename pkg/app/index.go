package app

import (
	"net/http"

	"github.com/bosemian/gogo/pkg/model"
	"github.com/bosemian/gogo/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	list := model.ListNews()
	view.Index(w, &view.IndexData{
		List: list,
	})
}
