check-swag:
	which swag || go install github.com/swaggo/swag/cmd/swag

check-mock:
	which mockgen || go install github.com/golang/mock/mockgen@v1.6.0

swag: check-swag
	go mod tidy && swag init

mock: check-mock
	mockgen -source ./repositories/poll.go -destination ./mocks/mock_repo_poll.go -package mocks
	mockgen -source ./clients/captcha/captcha.go -destination ./mocks/mock_cli_captcha.go -package mocks
	mockgen -source ./clients/pubsub/pubsub.go -destination ./mocks/mock_cli_pubsub.go -package mocks
envup: 
	docker-compose build
	docker-compose up -d

envdown: 
	docker-compose down

test: 
	go test ./...

# attack:
# 	vegeta attack -duration=10s -rate=100 -targets=target.conf | vegeta report