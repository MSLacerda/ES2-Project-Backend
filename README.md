# ES2-Project-Backend

## Rotas

### Casos de uso

- Listar casos de uso (por partes)

Lista todos os dados dos casos de uso para a disposição do usuário a fim de montar possível resposta para o problema.

```http
/api/casos GET
```
Resposta:
```json
{
  "data": [
    {
      "id": 1,
      "grupo": 1,
      "codigo_entidade": "A",
      "conteudo": "usuário"
    },
    {
      "id": 2,
      "grupo": 1,
      "codigo_entidade": "C",
      "conteudo": "executa comando"
    }
  ]
}
```

- Buscar caso de uso (por partes)

Busca um caso específico (por partes)

```http
/api/casos/{grupo} GET
```
Resposta: O "grupo" significa que o conteúdo listado é um parte (caso de uso, ator, etc) do caso de uso 1 (se o grupo for 1).
```json
{
  "data": [
    {
      "id": 1,
      "grupo": 1,
      "codigo_entidade": "A",
      "conteudo": "usuário"
    },
    {
      "id": 3,
      "grupo": 1,
      "codigo_entidade": "C",
      "conteudo": "requisita comando"
    }
  ]
}
```

- Conferir caso de uso

Confere se o caso de uso resolvido pelo usuário está correto

```http
/api/casos/{grupo} PUT
```
Corpo: todos os "conteúdos" (atores, casos de uso, extends) pertencentes a um mesmo caso de uso. No caso do "extends", a avaliação ocorre em pares de caso de uso em que **A** entends **B**
```json
{
  "grupo": 1,
  "atores": [3, 4, 5],
  "casos_de_uso": [6, 1, 2],
  "extends": [6, 1]
}
```

Resposta
```json
{
  "correto": true
}
```

### Estorias de usuário

- Listar estorias de usuário (por partes)

```http
/api/estorias
```
Resposta
```json
{
  "data": [
    {
      "id": 1,
      "grupo": 1,
      "codigo_entidade": "Q",
      "conteúdo": "quero poder executar obter ajuda"
    }
  ]
}
```
