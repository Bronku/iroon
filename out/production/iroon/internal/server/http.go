package server

import "net/http"

func (s *Server) redirect(path string, code int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, path, code)
	}
}
