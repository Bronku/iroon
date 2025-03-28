package iroon

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Bronku/iroon/internal/auth"
	"github.com/Bronku/iroon/internal/logging"
	"github.com/Bronku/iroon/internal/server"
	"github.com/Bronku/iroon/internal/store"
)

func Run() {
	s := store.OpenStore("./foo.db")
	defer s.Close()
	h := server.New(s)
	defer h.Close()

	var handler http.Handler = h
	handler = logging.Middleware(handler)
	handler = auth.New(s).Middleware(handler)
	fmt.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
