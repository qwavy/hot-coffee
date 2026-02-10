package router

import (
	"hot-coffee/internal/dal"
	"hot-coffee/internal/handler"
	"hot-coffee/internal/service"
	"net/http"
)

func RunServer(cfg config.Config) error {
	router := Router()
	fmt.Printf("Server started on http://localhost:%d\n", cfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
}
