package app

import "net/http"

// Mount handler mux
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", index) // list all news
	mux.Handle("/upload/", http.StripPrefix("/upload", http.FileServer(http.Dir("upload"))))
	mux.Handle("/news/", http.StripPrefix("/news", http.HandlerFunc(newsView)))

	mux.HandleFunc("/register/", adminRegister)

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/delete", adminDelete)
	adminMux.HandleFunc("/login", adminLogin)
	adminMux.HandleFunc("/list", adminList)
	adminMux.HandleFunc("/create", adminCreate)
	adminMux.HandleFunc("/edit", adminEdit)

	mux.Handle("/admin/", http.StripPrefix("/admin", onlyAdmin(adminMux)))

}

func onlyAdmin(h http.Handler) http.Handler {
	return h
}
