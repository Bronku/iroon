package server

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/Bronku/iroon/internal/store"
)

type Server struct {
	tmpl   map[string]*template.Template
	s      *store.Store
	routes map[string]route
	http.Handler
}

type route struct {
	function fetcher
	template string
}

func (h *Server) Close() {
}

//go:embed static/*
var static embed.FS

func (h *Server) loadHandler() {
	mux := http.NewServeMux()

	for i, e := range h.routes {
		mux.HandleFunc(i, h.render(e.function, e.template))
	}

	mux.HandleFunc("GET /", h.redirect("/orders", http.StatusSeeOther))

	fs := http.FileServerFS(static)
	mux.Handle("GET /static/", fs)

	h.Handler = mux
}

func New(store *store.Store) *Server {
	var server Server

	server.routes = map[string]route{
		"GET /order/":  {server.order, "order"},
		"GET /orders":  {server.orders, "orders"},
		"GET /cake/":   {server.cake, "cake"},
		"GET /cakes":   {server.cakes, "cakes"},
		"POST /order/": {server.postOrder, "confirmation"},
		"POST /cake/":  {server.postCake, "confirmation"},
	}

	server.loadTemplates()
	server.s = store
	server.loadHandler()

	return &server
}
