package main

import (
	"log"
	"net/http"
)

func main() {
	route := GetRouter()

	log.Fatal(http.ListenAndServe(":8080", route))
}