package iroon

import (
	"log"
	"net/http"

	"github.com/Bronku/iroon/internal/auth"
	"github.com/Bronku/iroon/internal/logging"
	"github.com/Bronku/iroon/internal/server"
)

func Run() {
	h := server.New()
	defer h.Close()

	var handler http.Handler = h
	handler = logging.Middleware(handler)
	handler = auth.New().Middleware(handler)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
