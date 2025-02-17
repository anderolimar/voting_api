# Documentação do serviço de Votação

## Introdução

Este documento descreve o funcionamento do serviço de votação. O serviço gerencia enquetes (polls) e votos por meio de rotas HTTP.


### Endpoints

#### `GET /poll`

Obtém a ultima de enquetes disponíveis.

**Resposta:** JSON contendo uma enquete.

```json
{
    "statusCode": 200,
    "poll": {
        "id": "67b27749ad94fd1682caf84a",
        "title": "Escolha quem voçê quer eliminar nesse paredão?",
        "options": [
            {
                "index": 1,
                "title": "Participante 1",
                "quantity": 0
            },
            {
                "index": 2,
                "title": "Participante 2",
                "quantity": 0
            }
        ]
    }
}

```

#### `POST /poll`

Cria uma nova enquete a partir dos dados fornecidos no corpo da requisição.

**Requisição:** JSON com os dados da enquete.
```json
{
    "title": "Escolha quem voçê quer eliminar nesse paredão?",
    "options": [
        { "index": 1, "title": "Participante 1" },
        { "index": 2, "title": "Participante 2" }
    ]
}
```

**Resposta:** JSON com status da operação.
```json
{
    "statusCode": 200
}
```

#### `POST /vote`

Registra um novo voto em uma enquete existente.

**Requisição:** JSON contendo os detalhes do voto.

```json
{
    "captchaID":"1fj0QkqzXOx13cZy8VNh",
    "captchaInput":"0239",
    "vote":2,
    "pollID":"67b2b459aa46985b1ede01ef"
}
```

**Resposta:** JSON com resultado parcial da enquete.

```json
{
    "statusCode": 200,
	"poll": {
		"id": "67b2b459aa46985b1ede01ef",
		"title": "Escolha quem voçê quer eliminar nesse paredão?",
		"options": [
			{ "index": 1, "title": "Participante 1" },
			{ "index": 2, "title": "Participante 2" }
		]
	}
}
```

**Resposta com erro** CAPTCHA Invalido

```json
{
    "code": "INVALID_CAPTCHA",
    "message": "Invalid CAPTCHA answer",
    "statusCode": 400
}
```

#### `GET /captcha`

Gera um novo CAPTCHA e retorna a resposta.

**Resposta:** JSON contendo os dados do CAPTCHA gerado.

```json
{
    "statusCode": 200,
    "id": "bKEy32motr44f1g8keOG",
    "base64Img": "data:image/png;base64,iVBORw0KGgoA..."
}
```

