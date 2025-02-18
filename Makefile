
check-mock:
	which mockgen || go install github.com/golang/mock/mockgen@v1.6.0

mock: check-mock
	mockgen -source ./repositories/poll.go -destination ./mocks/mock_repo_poll.go -package mocks
	mockgen -source ./clients/captcha/captcha.go -destination ./mocks/mock_cli_captcha.go -package mocks
	mockgen -source ./clients/pubsub/pubsub.go -destination ./mocks/mock_cli_pubsub.go -package mocks

envup: 
	docker-compose build
	docker-compose up -d

envdown: 
	docker-compose down

runservice:
	go run main.go service

runworker:
	go run main.go worker

test: 
	go test ./...

setup:
	curl --location 'http://localhost:8080/poll' --header 'Content-Type: application/json' --data '{"title": "Escolha quem voçê quer eliminar nesse paredão?", "options": [{ "index": 1, "title": "Participante 1" },{ "index": 2, "title": "Participante 2" },{ "index": 3, "title": "Participante 3" },{ "index": 4, "title": "Participante 4" }]}'
	curl --location 'http://localhost:8081/poll' --header 'Content-Type: application/json' --data '{"title": "Escolha quem voçê quer eliminar nesse paredão?", "options": [{ "index": 1, "title": "Participante 1" },{ "index": 2, "title": "Participante 2" },{ "index": 3, "title": "Participante 3" },{ "index": 4, "title": "Participante 4" }]}'

attack:
	vegeta attack -duration=10s -rate=1000 -targets=target.conf | vegeta report