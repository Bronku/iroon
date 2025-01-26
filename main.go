package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
)

//go:embed public
var public embed.FS

//go:embed templates
var templates embed.FS

func unwrap[T any](output T, err error) T {
	if err != nil {
		panic(err)
	}
	return output
}

func getNewOrder(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		_, _ = fmt.Fprint(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "new_order.gohtml", nil)
	if err != nil {
		_, _ = fmt.Fprint(w, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func postNewOrder(w http.ResponseWriter, r *http.Request) {

}

func main() {
	server := http.NewServeMux()

	server.Handle("/", http.FileServerFS(unwrap(fs.Sub(public, "public"))))
	server.HandleFunc("GET /order/new", getNewOrder)
	server.HandleFunc("POST /order/new", postNewOrder)

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Println("Can't start the server: ", err)
	}
}
