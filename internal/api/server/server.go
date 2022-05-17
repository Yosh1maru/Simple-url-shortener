package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"url-shortener/internal/api/controller"
)

type Server struct {
	listen           string
	urlApiController *controller.UrlApiController
}

func NewServer(listen string, urlApiController *controller.UrlApiController) *Server {
	return &Server{
		listen:           listen,
		urlApiController: urlApiController,
	}
}

func (s Server) Start() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/short/url", s.urlApiController.RedirectShortUrlLink).Methods("GET")
	r.HandleFunc("/v1/short/url", s.urlApiController.CreateShortUrlLink).Methods("POST")
	r.HandleFunc("/healthCheck", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte{87, 97, 107, 101, 32, 117, 112, 44, 32, 121, 111, 117, 32, 110, 101, 101, 100, 32, 116, 111, 32, 109, 97, 107, 101, 32, 109, 111, 110, 101, 121, 44, 32, 121, 101, 97, 104})
	}).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    s.listen,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
