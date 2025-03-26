package server

import (
	"html/template"
	"net/http"

	"github.com/Bronku/iroon/internal/store"
)

type Server struct {
	tmpl *template.Template
	s    *store.Store
	http.Handler
}

func (h *Server) Close() {
	if h.s != nil {
		h.s.Close()
	}
}

func (h *Server) loadHandler() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /order/", h.render(h.order, "order.html"))
	mux.HandleFunc("GET /", h.render(h.index, "index.html"))
	mux.HandleFunc("POST /order/", h.render(h.postOrder, "confirmation.html"))
	mux.HandleFunc("GET /cakes", h.render(h.cakes, "cakes.html"))
	mux.HandleFunc("GET /cake/", h.render(h.cake, "cake.html"))
	mux.HandleFunc("POST /cake/", h.render(h.postCake, "confirmation.html"))

	h.Handler = mux
}

func New() *Server {
	var server Server

	server.loadTemplates()
	server.s = store.OpenStore("./foo.db")
	server.loadHandler()

	return &server
}
