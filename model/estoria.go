package model

type Estoria struct {
	Estoria int64 `json:"estoria"`// referência a estoria
	PEstoria int64 `json:"p_estoria"` // referência a palavra da estoria
	Conteudo string `json:"conteudo"` // conteudo da estoria
}
