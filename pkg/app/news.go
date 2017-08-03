package app

import (
	"net/http"

	"github.com/bosemian/gogo/pkg/model"

	"github.com/bosemian/gogo/pkg/view"
)

func newsView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	new, err := model.GetNews(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.News(w, new)
}
