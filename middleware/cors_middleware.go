package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func CorsConfig() http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(http.DefaultServeMux)

	return handler
}
