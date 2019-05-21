package app

import (
	"log"
	"net/http"
	"os"

	"github.com/MSLacerda/ES2-Project-Backend/app/router"
)

func BuiltServer() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	route := router.GetRouter()

	log.Fatal(http.ListenAndServe(":"+port, route))
}
