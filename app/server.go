package app

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

	"github.com/MSLacerda/ES2-Project-Backend/app/router"
)

func BuiltServer() {
	port := os.Getenv("PORT")
	// corsObj := handlers.AllowedOrigins([]string{"*"})
	// methodsObj := handlers.AllowedMethods([]string{"POST", "PUT", "GET", "DELETE", "OPTIONS"})
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                                       // All origins
		AllowedMethods: []string{"GET", "OPTIONS", "POST", "DELETE", "PUT"}, // Allowing only get, just an example
	})

	if port == "" {
		port = "8080"
	}

	route := router.GetRouter()

	log.Fatal(http.ListenAndServe(":"+port, c.Handler(route)))
}
