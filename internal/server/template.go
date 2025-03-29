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
	h.tmpl = make(map[string]*template.Template)
	pages := []string{"orders", "order", "cake", "cakes", "confirmation"}
	for _, page := range pages {
		tmpl, err := template.ParseFS(templates,
			"templates/layout/*.html",
			"templates/"+page+".html",
		)
		if err != nil {
			log.Fatal(err)
		}
		h.tmpl[page] = tmpl
	}
}

func (s *Server) render(fetch fetcher, templateName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, code, err := fetch(r)
		if err != nil {
			logging.ErrorPage(err, code).ServeHTTP(w, r)
			return
		}
		err = s.tmpl[templateName].ExecuteTemplate(w, "layout", data)
		if err != nil {
			logging.ErrorPage(err, http.StatusInternalServerError).ServeHTTP(w, r)
			return
		}
		w.Header().Set("content-type", "text/html")
	}
}
