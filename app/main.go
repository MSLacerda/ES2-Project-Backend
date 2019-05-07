package main

import (
	"github.com/MSLacerda/ES2-Project-Backend/app/router"
	"log"
	"net/http"
)

func main() {
	route := router.GetRouter()

	log.Fatal(http.ListenAndServe(":8080", route))
}