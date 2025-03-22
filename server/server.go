package server

import (
	"html/template"
	"net/http"

	"github.com/Bronku/iroon/store"
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

func (h *Server) openStore() error {
	var err error
	h.s, err = store.OpenStore("./foo.db")
	if err != nil {
		h.s.Close()
	}
	return err
}

func (h *Server) loadHandler() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /order/", h.render(h.getOrder, "order.html"))
	mux.HandleFunc("GET /", h.render(h.index, "index.html"))
	mux.HandleFunc("POST /order/", h.postOrder)

	h.Handler = mux
}

func New() (*Server, error) {
	var server Server
	var err error

	err = server.loadTemplates()
	if err != nil {
		return nil, err
	}

	err = server.openStore()
	if err != nil {
		return nil, err
	}

	server.loadHandler()

	return &server, nil
}
