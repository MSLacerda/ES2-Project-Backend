package data

import (
	"encoding/json"
	"fmt"
	"github.com/MSLacerda/ES2-Project-Backend/model"
	"io/ioutil"
	"log"
	"os"
)

func BuscarCaso() ([]model.Caso, error) {
	type respCasos struct {
		Casos []model.Caso `json:"casos"`
	}

	casos := respCasos{}

	jsonFile, err := os.Open("casos.json")
	if err != nil {
		log.Println(err)
		return casos.Casos, fmt.Errorf("Erro ao buscar dados dos casos")
	}

	defer jsonFile.Close()

	byteValue,_ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &casos)

	return casos.Casos, nil
}
