package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"strings"
)

func unwrap[T any](output T, err error) T {
	if err != nil {
		panic(err)
	}
	return output
}

func LoadTemplates(fs fs.ReadDirFS) map[string]*template.Template {
	files, err := fs.ReadDir("templates")
	if err != nil {
		fmt.Println("can't read the templates directory")
		return nil
	}
	tmpl := make(map[string]*template.Template)
	for _, e := range files {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		name = strings.Split(name, ".")[0]
		tmpl[name], err = template.ParseFS(fs, "templates/"+name+".html")
		if err != nil {
			fmt.Println("can't open template: " + name)
		}
	}
	return tmpl
}
