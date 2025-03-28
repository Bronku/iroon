package server

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/Bronku/iroon/internal/logging"
)

//go:embed templates/*
var templates embed.FS

func (h *Server) loadTemplates() {
	var err error
	h.tmpl, err = template.ParseFS(templates, "templates/*")
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) render(fetch fetcher, templateName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html")
		data, code, err := fetch(r)
		if err != nil {
			logging.ErrorPage(err, code).ServeHTTP(w, r)
			return
		}
		err = s.tmpl.ExecuteTemplate(w, templateName, data)
		if err != nil {
			logging.ErrorPage(err, http.StatusInternalServerError).ServeHTTP(w, r)
		}
	}
}
