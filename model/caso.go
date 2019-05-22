package model

type Caso struct {
	Codigo int `json:"codigo"`
	CodigoUsuario int `json:"codigo_usuario,omitempty"`
	Conteudo string `json:"conteudo"`
}
