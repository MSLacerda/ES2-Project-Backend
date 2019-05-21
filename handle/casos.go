package handle

import (
	"github.com/MSLacerda/ES2-Project-Backend/data"
	"log"
	"net/http"
)

func BuscarCaso (w http.ResponseWriter, r *http.Request) {
	caso, err := data.BuscarCaso()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = JSON(w, 200, caso)
	if err != nil {
		log.Println(err)
	}
}

func ConferirCaso (w http.ResponseWriter, r *http.Request) {
	log.Printf("todo")
}
