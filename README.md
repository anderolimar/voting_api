# Voting API


### **API de Votação**  

A API de Votação permite que os usuários criem enquetes e disponibilize a votação de maneira simples e eficiente.  

![Fluxo de Votação](./assets/vote_workflow.jpg)

###  **Iniciar ambiente**
```sh
make envup
```
- **`docker-compose build`**: Constrói os containers do projeto usando o `docker-compose`.
- **`docker-compose up -d`**: Inicia os containers em **modo desacoplado** (`-d`), ou seja, em segundo plano.

###  **Gerar enquete**
```sh
make setup
```
- Faz uma **requisição HTTP POST** para `http://localhost:8080/poll`.
- Envia um **payload JSON** que cria uma votação com 4 opções (participantes).
- Acesse a enquete no endereço [http://localhost:8080](http://localhost:8080)

### 4. **Finalizar ambiente**
```sh
make envdown
```
- Para e remove os containers e redes criados pelo `docker-compose`.

### 5. **Executar testes**
```sh
make test
```
- Executa **todos os testes unitários** do projeto escritos em Go.
- O `./...` faz com que os testes sejam rodados em todos os pacotes do projeto.


### **Gerar mock**
```sh
make mock: check-mock
```
- Depende do comando `check-mock` (garante que `mockgen` esteja instalado).
- Gera mocks para facilitar testes unitários:
  - Cria um mock para `poll.go` e salva em `mocks/mock_repo_poll.go`.
  - Cria um mock para `captcha.go` e salva em `mocks/mock_cli_captcha.go`.
  - Cria um mock para `pubsub.go` e salva em `mocks/mock_cli_pubsub.go`.

## Documentação

A documentação dos endpoint encontrasse no arquivo [DOCUMENTATION.md](./DOCUMENTATION.md)

## **Possíveis Melhorias Futuras:**  
- Autenticação para rota de criação de enquete  
- Melhoria nos logs
- Adicionar rate limit 
- Utilizar outro tipo de fila de mensagens
- Busca de enquete por ID
- Implementar testes de integração
- Adicionar test coverage 

