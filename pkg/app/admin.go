package app

import "net/http"
import "github.com/bosemian/gogo/pkg/view"

func adminLogin(w http.ResponseWriter, r *http.Request) {
	view.AdminLogin(w, nil)
}

func adminList(w http.ResponseWriter, r *http.Request) {
	view.AdminList(w, nil)
}

func adminEdit(w http.ResponseWriter, r *http.Request) {

}

func adminCreate(w http.ResponseWriter, r *http.Request) {

}
