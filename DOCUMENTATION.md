# Documentação do serviço de Votação

## Introdução

Este documento descreve o funcionamento do serviço de votação. O serviço gerencia enquetes (polls) e votos por meio de rotas HTTP.


## Estrutura dos Handlers

### `PollHandler`

O `PollHandler` é a estrutura principal deste módulo e implementa as funcionalidades relacionadas à criação de enquetes e votos.

#### Construtor

```go
func NewPollHandler() *PollHandler
```
Cria e retorna uma nova instância de `PollHandler`.

#### Propriedades

```go
type PollHandler struct {
	commom.CommonsHandler
	service services.PollService
}
```
- `CommonsHandler` - Estrutura base para handlers HTTP.
- `service` - Instância do serviço de enquete.

## Rotas

A função `RegisterRoutes` registra as rotas manipuladas pelo `PollHandler`:

```go
func (v PollHandler) RegisterRoutes(router *http.ServeMux)
```

### Endpoints

#### `GET /poll`

```go
func (v PollHandler) Poll(w http.ResponseWriter, r *http.Request)
```
Obtém a lista de enquetes disponíveis.

**Resposta:** JSON contendo as enquetes registradas.

#### `POST /poll`

```go
func (v PollHandler) NewPoll(w http.ResponseWriter, r *http.Request)
```
Cria uma nova enquete a partir dos dados fornecidos no corpo da requisição.

**Requisição:** JSON com os dados da enquete.

**Resposta:** JSON com status da operação.

#### `POST /vote`

```go
func (v PollHandler) Vote(w http.ResponseWriter, r *http.Request)
```
Registra um novo voto em uma enquete existente.

**Requisição:** JSON contendo os detalhes do voto.

**Resposta:** JSON com status da operação.


<br>
<br>

### `CaptchaHandler`

O `CaptchaHandler` é a estrutura principal deste módulo e implementa a funcionalidade de geração de CAPTCHA.

#### Construtor

```go
func NewVoteHandler() *CaptchaHandler
```
Cria e retorna uma nova instância de `CaptchaHandler`.

#### Propriedades

```go
type CaptchaHandler struct {
	commom.CommonsHandler
	captchaService captcha.CaptchaService
}
```
- `CommonsHandler` - Estrutura base para handlers HTTP.
- `captchaService` - Instância do serviço de geração de CAPTCHA.

## Rotas

A função `RegisterRoutes` registra as rotas manipuladas pelo `CaptchaHandler`:

```go
func (c CaptchaHandler) RegisterRoutes(router *http.ServeMux)
```

### Endpoints

#### `GET /captcha`

```go
func (c CaptchaHandler) Captcha(w http.ResponseWriter, r *http.Request)
```
Gera um novo CAPTCHA e retorna a resposta.

**Resposta:** JSON contendo os dados do CAPTCHA gerado.



