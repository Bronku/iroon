package server

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates/*
var templates embed.FS

func (h *Server) loadTemplates() error {
	var err error
	h.tmpl, err = template.ParseFS(templates, "templates/*")
	return err
}

func (s *Server) render(fetch fetcher, templateName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html")
		data, code, err := fetch(r)
		if err != nil {
			errorPage(err, code).ServeHTTP(w, r)
			return
		}
		err = s.tmpl.ExecuteTemplate(w, templateName, data)
		if err != nil {
			errorPage(err, http.StatusInternalServerError).ServeHTTP(w, r)
		}
	}
}
