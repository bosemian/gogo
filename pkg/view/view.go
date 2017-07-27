package view

import "net/http"
import "github.com/bosemian/gogo/pkg/model"

type IndexData struct {
	List []*model.News
}

// Index render index view
func Index(w http.ResponseWriter, data *IndexData) {
	render(tpIndex, w, data)
}

// AdminLogin render admin view
func AdminLogin(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}

// AdminList render adminlist view
func AdminList(w http.ResponseWriter, data interface{}) {
	render(tpAdminList, w, data)
}

// AdminCreate render admin create view
func AdminCreate(w http.ResponseWriter, data interface{}) {
	render(tpAdminCreate, w, data)
}

// AdminEdit render admin edit view
func AdminEdit(w http.ResponseWriter, data interface{}) {
	render(tpAdminEdit, w, data)
}
