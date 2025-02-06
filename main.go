package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type cake struct {
	Name   string
	ID     int
	Price  int
	Amount int
}

type order struct {
	ID       int
	Name     string
	Surname  string
	Phone    string
	Location string
	Date     time.Time
	Paid     int // increments of 0.01
	Cakes    []cake
}

type handler struct {
	tmpl   *template.Template
	cakes  []cake
	orders []order
}

func (h *handler) form(w http.ResponseWriter, r *http.Request) {
	data := order{
		ID:    -1,
		Date:  time.Now(),
		Cakes: h.cakes,
	}
	w.Header().Set("content-type", "text/html")
	err := h.tmpl.ExecuteTemplate(w, "form.html", data)
	if err != nil {
		fmt.Println("error executing the template: ", err)
	}
}

func (h *handler) addOrder(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("can't parse the form")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("received form: ", r.Form)

	var n order
	n.ID, err = strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("can't parse order id: ", err)
		return
	}
	n.Name = r.FormValue("name")
	n.Surname = r.FormValue("surname")
	n.Phone = r.FormValue("phone")
	n.Location = r.FormValue("location")
	n.Date, err = time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("can't parse order date: ", err)
		return
	}
	n.Paid, err = strconv.Atoi(r.FormValue("paid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("can't parse order paid: ", err)
		return
	}
	n.Cakes = make([]cake, 0)

	h.orders = append(h.orders, n)

	fmt.Println("parsed order: ", n)

	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("accepted <a href='/'>back</a>"))
}

func (h *handler) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	err := h.tmpl.ExecuteTemplate(w, "index.html", h.orders)
	if err != nil {
		fmt.Println("error executing the template: ", err)
	}
}

func logger(in http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log := fmt.Sprint(r.Method, " ", r.URL.String())
		in(w, r)
		log += fmt.Sprint(" ", time.Since(start))
		fmt.Println(log)
	}
}

func main() {
	var h handler
	templates, err := template.ParseFiles("index.html", "form.html")
	if err != nil {
		log.Fatal("can't parse templates: ", err)
	}

	h.tmpl = templates

	orders := make([]order, 0)
	orders = append(orders, order{0, "Albert", "Camus", "123456789", "Kartuzy", time.Now().AddDate(0, 0, 7), 0, nil})
	orders = append(orders, order{1, "George", "Orwell", "", "Kartuzy", time.Now().AddDate(0, 1, 0), 0, nil})
	orders = append(orders, order{2, "Karl", "Marx", "0700", "Somonino", time.Now(), 0, nil})
	h.orders = orders

	cakes := make([]cake, 0)
	cakes = append(cakes, cake{"Sernik", 0, 120, 0})
	cakes = append(cakes, cake{"Malinowa chmurka", 1, 120, 0})
	h.cakes = cakes

	http.HandleFunc("GET /form", logger(h.form))
	http.HandleFunc("GET /", logger(h.index))
	http.HandleFunc("POST /", logger(h.addOrder))
	http.ListenAndServe(":8080", nil)
}
