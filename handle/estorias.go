package handle

import (
	"encoding/json"
	"fmt"
	"github.com/MSLacerda/ES2-Project-Backend/data"
	"github.com/MSLacerda/ES2-Project-Backend/model"
	"github.com/gorilla/mux"
	"math/rand"
	"reflect"
	"strconv"

	"log"
	"net/http"
)

func ListarEstorias (w http.ResponseWriter, r *http.Request) {
	log.Printf("1")
	estorias, err := data.ListarEstorias()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = JSON(w, 200, estorias)
	if err != nil {
		log.Println(err)
	}
}

func BuscarEstoria (w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	estoria, err := data.BuscarEstoria(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Para deixar ordem aleatória
	respEstoria := make([]model.Estoria, len(estoria))
	ordem := rand.Perm(len(estoria))

	for i, e := range ordem {
		respEstoria[i] = estoria[e]
	}

	err = JSON(w, 200, respEstoria)
	if err != nil {
		log.Println(err)
	}
}

func ConferirEstoria (w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

	type respEstoria struct {
		Estorias []model.Estoria `json:"estorias"`
	}

	dados := respEstoria{}

	err = json.NewDecoder(r.Body).Decode(&dados.Estorias)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := data.ConferirEstoria(dados.Estorias, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = JSON(w, 200, resp)
	if err != nil {
		log.Println(err)
	}
}

// Auxiliares

func JSON(res http.ResponseWriter, status int, retorno interface{}) error {

	var retornoJSON []byte
	if retorno != nil && reflect.TypeOf(retorno).Kind().String() == "string" {
		var retornoMAP map[string]interface{}
		if err := json.Unmarshal([]byte(retorno.(string)), &retornoMAP); err != nil {
			return fmt.Errorf("Erro ao converter json")
		}
		retornoJSON = []byte(retorno.(string))
	} else if retorno != nil {
		var err error
		retornoJSON, err = json.Marshal(retorno)
		if err != nil {
			return fmt.Errorf("Erro ao converter json")
		}
	}
	res.Header().Set("content-type", "application-json")
	res.WriteHeader(status)
	if retorno != nil {
		res.Write(retornoJSON)
	}
	return nil
}