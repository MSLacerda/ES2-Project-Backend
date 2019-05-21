package data

import (
	"encoding/json"
	"fmt"
	"github.com/MSLacerda/ES2-Project-Backend/model"
	"io/ioutil"
	"log"
	"os"
)

/*func ListarEstorias() ([]model.Estoria, error) {
	estorias := make([]model.Estoria, 0)

	db := database.GetConn()

	err := db.Model(&estorias).Select()
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Erro ao buscar estorias de usuário")
	}

	return estorias, nil
}

func BuscarEstoria(id int64) (model.Estoria, error) {
	estoria := model.Estoria{ID: id}

	db := database.GetConn()

	err := db.Model(&estoria).Select()
	if err != nil {
		log.Println(err)
		return estoria, fmt.Errorf("Erro ao buscar estorias de usuário")
	}

	return estoria, nil
}*/

func ListarEstorias() ([]model.Estoria, error) {
	type respEstoria struct {
		Estorias []model.Estoria `json:"estorias"`
	}

	estorias := respEstoria{}

	jsonFile, err := os.Open("estorias.json")
	if err != nil {
		log.Println(err)
		return estorias.Estorias, fmt.Errorf("Erro ao buscar dados das estorias")
	}

	defer jsonFile.Close()

	byteValue,_ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &estorias)

	return estorias.Estorias, nil
}

func BuscarEstoria(id int64) ([]model.Estoria, error) {
	respEstorias := make([]model.Estoria, 0)

	estorias, err := ListarEstorias()
	if err != nil {
		log.Println(err)
		return respEstorias, err
	}

	for _, e := range estorias {
		if e.Estoria == id {
			respEstorias = append(respEstorias, e)
		}
	}

	return respEstorias, nil
}

func ConferirEstoria(estoria []model.Estoria, id int64) bool {
	estorias, err := BuscarEstoria(id)
	if err != nil {
		return false

	}

	if len(estorias) != len(estoria) {
		return false
	}

	for i, e := range estoria {
		if e.Conteudo != estorias[i].Conteudo {
			return false
		}
	}

	return true
}