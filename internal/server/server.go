package server

import (
	"fmt"
	"hot-coffee/internal/config"
	"net/http"
)

func RunServer(cfg config.Config) error {
	router := Router()
	fmt.Printf("Server started on http://localhost:%d\n", cfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
}
