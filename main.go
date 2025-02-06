package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func form(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received a get request at: ", r.URL.String())
	tmpl, err := template.ParseFiles("form.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	type cake struct {
		Name   string
		ID     int
		Price  int
		Amount int
	}
	type tmplData struct {
		Name    string
		Surname string
		Phone   string
		Date    string
		Cakes   []cake
	}
	data := tmplData{
		Name:    "",
		Surname: "",
		Phone:   "",
		Date:    "",
		Cakes:   make([]cake, 0),
	}
	data.Cakes = append(data.Cakes, cake{"Hello", 0, 100, 0})
	data.Cakes = append(data.Cakes, cake{"World", 1, 120, 0})
	w.Header().Set("content-type", "text/html")
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("error executing the template: ", err)
	}
}

func addOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received a post request at: ", r.URL.String())
	err := r.ParseForm()
	if err != nil {
		fmt.Println("can't parse the form")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("received form: ", r.Form)
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("accepted <a href='/'>back</a>"))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received a get request at: ", r.URL.String())
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type tmplData struct {
		ID      int
		Name    string
		Surname string
	}
	data := make([]tmplData, 0)
	data = append(data, tmplData{0, "Jan", "Kowalski"})
	data = append(data, tmplData{1, "Jakub", "Bronk"})

	w.Header().Set("content-type", "text/html")
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("error executing the template: ", err)
	}
}

func main() {
	http.HandleFunc("GET /form", form)
	http.HandleFunc("GET /", index)
	http.HandleFunc("POST /", addOrder)
	http.ListenAndServe(":8080", nil)
}
