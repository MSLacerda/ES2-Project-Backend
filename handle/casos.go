package handle

import (
	"encoding/json"
	"github.com/MSLacerda/ES2-Project-Backend/model"
	"log"
	"net/http"

	"github.com/MSLacerda/ES2-Project-Backend/data"
)

func BuscarCaso(w http.ResponseWriter, r *http.Request) {
	caso, err := data.BuscarCaso()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = JSON(w, 200, caso)
	if err != nil {
		log.Println(err)
	}
}

func ConferirCaso(w http.ResponseWriter, r *http.Request) {
	type respCaso struct {
		Casos []model.Caso
	}

	dados := respCaso{}

	err := json.NewDecoder(r.Body).Decode(&dados.Casos)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := data.ConferirCaso(dados.Casos)

	type respApresentacao struct {
		Correto bool `json:"correto"`
	}

	respCasoAprensetacao := respApresentacao{}
	respCasoAprensetacao.Correto = resp

	err = JSON(w, 200, respCasoAprensetacao)
	if err != nil {
		log.Println(err)
	}
}
