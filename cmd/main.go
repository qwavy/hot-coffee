package main

import (
	"hot-coffee/internal/router"
	"net/http"
)

func main() {
	r := router.Router()

	http.ListenAndServe("localhost:8080", r)

}
