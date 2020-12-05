package server

import (
	"net/http"
)

func (s *server) routes() {
	r := s.router

	// Fileserver Routes
	publicHandler := http.FileServer(http.Dir("./client/public/"))
	r.PathPrefix("/images/").
		Handler(publicHandler)
	r.PathPrefix("/assets/").
		Handler(publicHandler)
	r.PathPrefix("/markdown/").
		Handler(publicHandler)
	r.PathPrefix("/feeds/").
		Handler(publicHandler)

	r.HandleFunc("/", s.handleTemplate(nil, "pages/home"))
}
