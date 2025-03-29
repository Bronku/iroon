package server

import (
	"html/template"
	"net/http"

	"github.com/Bronku/iroon/internal/store"
)

type Server struct {
	tmpl map[string]*template.Template
	s    *store.Store
	http.Handler
}

func (h *Server) Close() {
}

func (h *Server) loadHandler() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /order/", h.render(h.order, "order"))
	mux.HandleFunc("GET /", h.redirect("/orders", http.StatusSeeOther))
	mux.HandleFunc("GET /orders", h.render(h.orders, "orders"))
	mux.HandleFunc("POST /order/", h.render(h.postOrder, "confirmation"))
	mux.HandleFunc("GET /cakes", h.render(h.cakes, "cakes"))
	mux.HandleFunc("GET /cake/", h.render(h.cake, "cake"))
	mux.HandleFunc("POST /cake/", h.render(h.postCake, "confirmation"))

	h.Handler = mux
}

func New(store *store.Store) *Server {
	var server Server

	server.loadTemplates()
	server.s = store
	server.loadHandler()

	return &server
}
