package model

type Caso struct {
	ID int64 `json:"id"`
	Grupo int64 `json:"questao"`
	CodigoEntidade string `string:"codigo_entidade"`
	Conteudo string `string:"conteudo"`
}
