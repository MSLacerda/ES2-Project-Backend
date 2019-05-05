package model

type Estoria struct {
	ID int64 `json:"id"`

	// Informação sobre a etapa da história (Como um (C),
	// eu quero (Q) ou de modo que (M))
	CodigoEntidade string `json:"codigo_entidade"`

	// A ser exibido
	Conteudo string `json:"conteudo"`

	// A correção limita-se a saber se a estoria formada
	// estão agrupadas corretamente
	Grupo int64 `json:"grupo"`
}
