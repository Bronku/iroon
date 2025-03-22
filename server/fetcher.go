package server

import "net/http"

type fetcher func(r *http.Request) (any, int, error)
