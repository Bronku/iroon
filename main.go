package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed public
var public embed.FS

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
	fmt.Println(r.Form)
	w.WriteHeader(http.StatusAccepted)
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
