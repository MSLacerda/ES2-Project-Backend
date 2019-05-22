package handle

import (
	"encoding/json"
	"github.com/MSLacerda/ES2-Project-Backend/model"
	"log"
	"math/rand"
	"net/http"

	"github.com/MSLacerda/ES2-Project-Backend/data"
)

func BuscarCaso(w http.ResponseWriter, r *http.Request) {
	caso, err := data.BuscarCaso()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Para deixar ordem aleatória
	respCaso := make([]model.Caso, len(caso))
	ordem := rand.Perm(len(caso))

	for i, e := range ordem {
		respCaso[i] = caso[e]
	}

	err = JSON(w, 200, respCaso)
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
