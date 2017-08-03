package app

import "net/http"
import "github.com/bosemian/gogo/pkg/model"

// Mount handler mux
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", index) // list all news
	mux.Handle("/upload/", http.StripPrefix("/upload", http.FileServer(http.Dir("upload"))))
	mux.Handle("/news/", http.StripPrefix("/news", http.HandlerFunc(newsView)))

	mux.HandleFunc("/register/", adminRegister)
	mux.HandleFunc("/login", adminLogin)

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/delete", adminDelete)
	adminMux.HandleFunc("/logout", adminLogout)
	adminMux.HandleFunc("/list", adminList)
	adminMux.HandleFunc("/create", adminCreate)
	adminMux.HandleFunc("/edit", adminEdit)

	mux.Handle("/admin/", http.StripPrefix("/admin", onlyAdmin(adminMux)))

}

func onlyAdmin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := model.GetSession(r)
		ok, err := model.CheckUserID(sess.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !ok {
			http.Redirect(w, r, "/", http.StatusFound)
		}
		h.ServeHTTP(w, r)
	})
}
