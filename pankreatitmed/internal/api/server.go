package api

import (
	"net/http"

	"pankreatitmed/internal/app/handler"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer(h *handler.Handler) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/criteria", h.CriteriaPage)
	mux.HandleFunc("/criterion", h.CriterionPage)
	mux.HandleFunc("/order", h.OrderPage)

	fs := http.FileServer(http.Dir("resources"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return &Server{mux: mux}
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}
