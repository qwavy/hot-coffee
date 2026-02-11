package main

import (
	"hot-coffee/internal/server"
	"net/http"
)

func main() {
	r := server.Router()
	http.ListenAndServe("localhost:8080", r)
}
