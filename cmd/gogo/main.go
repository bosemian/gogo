package main

import (
	"log"
	"net/http"

	"github.com/bosemian/gogo/pkg/model"

	"github.com/bosemian/gogo/pkg/app"
)

const (
	port     = ":8080"
	mongoURL = "192.168.99.100:27017"
)

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)
	err := model.Init(mongoURL)
	if err != nil {
		log.Fatalf("cannot init model %v", err)
	}
	http.ListenAndServe(port, mux)
	log.Println("server has started...")
}
