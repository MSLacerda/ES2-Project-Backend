# ES2-Project-Backend

## Rotas

### Casos de uso

- Buscar caso de uso

Há apenas um caso de uso; Serão listados os conteúdos (atores, casos de uso, etc), já em ordem aleatória, para exibição para o usuário.

```http
/api/casos GET
```
Resposta
```json

  [
    {
      "codigo": 1,      
      "conteudo": "usuário"
    },
    {     
      "codigo": 2,      
      "conteudo": "Requisita comando"
    }
  ]

```

- Conferir caso de uso

Para conferir se o caso de uso está correto o é esperado o seguinte: para exibição ao usuário de um lado há um imagem do documento de caso de uso com números no lugar dos conteúdos; do outro lado há a lista de conteúdos. É esperado que o usuário possa associar cada número da imagem a um conteúdo da lista. É esperado uma lista de códigos, o código do usuário ```codigo_usuario``` e o código correto ```codigo```

```http
/api/casos PUT
```

Esperado
```json
  [
    {
      "codigo" : 1,
      "codigo_usuario": 2
    },
    {
      "codigo" : 3,
      "codigo_usuario": 3
    }
  ]
```

Resposta
```json
{
  "correto": true
}
```

### Estorias de usuário

- Buscar estoria de usuário

Lista todas as partes que vão compor uma estória de usuário (já embaralhadas)

```http
/api/estorias
```
Resposta
```json

  [
    {
      "estoria": 1,
      "p_estoria": 15,
      "conteudo": "da"
    },
    {
      "estoria": 1,
      "p_estoria": 16,
      "conteudo": "instruções"
    }
  ]

```

- Conferir estoria de usuário

Confere uma estória de usuário. Para isso recebe uma lista das palavras da estoria e compara se a ordem das mesmas é a esperada. Também se espera o ```ìd``` da estoria na URL.

```http
/api/estorias/{estoria_id}
```

Esperado
```json

  [
    {
      "estoria": 1,
      "p_estoria": 15,
      "conteudo": "da"
    },
    {
      "estoria": 1,
      "p_estoria": 16,
      "conteudo": "instruções"
    }
  ]

```

Resposta
```json
{
  "correto": true
}
```
