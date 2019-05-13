package app

import (
	"log"
	"net/http"

	"github.com/MSLacerda/ES2-Project-Backend/app/router"
)

func BuiltServer() {
	route := router.GetRouter()

	log.Fatal(http.ListenAndServe(":8080", route))
}
