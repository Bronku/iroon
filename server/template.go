package server

import (
	"embed"
	"html/template"
)

//go:embed templates/*
var templates embed.FS

func (h *Server) loadTemplates() error {
	var err error
	h.tmpl, err = template.ParseFS(templates, "templates/*")
	return err
}
