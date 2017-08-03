package app

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/bosemian/gogo/pkg/model"
	"github.com/bosemian/gogo/pkg/view"
)

func adminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}
	view.AdminLogin(w, nil)
}

func adminDelete(w http.ResponseWriter, r *http.Request) {
	model.DeleteNews(r.FormValue("id"))
	http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
}

func adminList(w http.ResponseWriter, r *http.Request) {
	list, err := model.ListNews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.AdminList(w, &view.AdminListData{
		List: list,
	})
}

func adminEdit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	n, err := model.GetNews(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		n.Title = r.FormValue("title")
		n.Detail = r.FormValue("detail")

		if file, fHeader, err := r.FormFile("image"); err == nil {
			defer file.Close()
			t := time.Now().Format(time.RFC3339)
			tm := t[0:10]
			filename := tm + "-" + fHeader.Filename
			fp, err := os.Create("upload/" + filename)
			if err == nil {
				io.Copy(fp, file)
			}
			fp.Close()
			n.Image = "/upload/" + filename
		}

		err := model.UpdateNews(n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}

	view.AdminEdit(w, n)
}

func adminCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		n := model.News{
			Title:  r.FormValue("title"),
			Detail: r.FormValue("detail"),
		}

		if file, fHeader, err := r.FormFile("image"); err == nil {
			defer file.Close()
			t := time.Now().Format(time.RFC3339)
			tm := t[0:10]
			filename := tm + "-" + fHeader.Filename
			fp, err := os.Create("upload/" + filename)
			if err == nil {
				io.Copy(fp, file)
			}
			fp.Close()
			n.Image = "/upload/" + filename
		}

		err := model.CreateNews(n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}
	view.AdminCreate(w, nil)
}

func adminRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		err := model.Register(username, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}
	view.AdminRegiter(w, nil)
}
