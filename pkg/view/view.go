package view

import "net/http"
import "github.com/bosemian/gogo/pkg/model"

type IndexData struct {
	List []*model.News
}

type AdminListData struct {
	List []*model.News
}

// Index render index view
func Index(w http.ResponseWriter, data *IndexData) {
	render(tpIndex, w, data)
}

// News render News page
func News(w http.ResponseWriter, data *model.News) {
	render(tpNews, w, data)
}

// AdminLogin render admin view
func AdminLogin(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}

// AdminRegiter render register view
func AdminRegiter(w http.ResponseWriter, data interface{}) {
	render(tpAdminRegister, w, data)
}

// AdminList render adminlist view
func AdminList(w http.ResponseWriter, data *AdminListData) {
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
