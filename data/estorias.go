package data

import (
	"fmt"
	"github.com/MSLacerda/ES2-Project-Backend/database"
	"github.com/MSLacerda/ES2-Project-Backend/model"
	"log"
)

func ListarEstorias() ([]model.Estoria, error) {
	var estorias []model.Estoria

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
}
