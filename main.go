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

func handlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received Post request: ", r)
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Println(r.Form)
	tmpl, _ := template.ParseFS(templates, "templates/order/confirmation.html")
	_ = tmpl.Execute(w, struct{ Status string }{Status: "accepted"})
}

func main() {
	server := http.NewServeMux()

	server.Handle("GET /", http.FileServerFS(unwrap(fs.Sub(public, "public"))))
	server.HandleFunc("POST /", handlePost)

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Println("Can't start the server: ", err)
	}
}
