all: test build-local

.PHONY: api
api:
	@protoc --proto_path=api/. --go_out=svcs/api-gateway/pkg/gateway/api --go_opt=paths=source_relative \
    --go-grpc_out=svcs/api-gateway/pkg/gateway/api --go-grpc_opt=paths=source_relative \
    api/secrets-gateway.proto
	@protoc --proto_path=api/. --go_out=svcs/secrets/pkg/secrets/api --go_opt=paths=source_relative \
    --go-grpc_out=svcs/secrets/pkg/secrets/api --go-grpc_opt=paths=source_relative \
    api/secrets-gateway.proto

.PHONY: test
test: testApiGateway testSecrets

.PHONY: testApiGateway
testApiGateway:
	@cd svcs/api-gateway && go test -race -coverprofile coverage.out ./...

.PHONY: testSecrets
testSecrets:
	@cd svcs/secrets && go test -race -coverprofile coverage.out ./...

.PHONY: coverage
coverage:
	@go tool cover -html svcs/api-gateway/coverage.out
	@go tool cover -html svcs/secrets/coverage.out

.PHONY: build-local
build-local: api buildLocalApiGateway buildLocalSecrets

.PHONY: buildLocalApiGateway
buildLocalApiGateway:
	@cd svcs/api-gateway && go build -a -o bin/api-gateway cmd/gateway/gateway.go

.PHONY: buildLocalSecrets
buildLocalSecrets:
	@cd svcs/secrets && go build -a -o bin/secrets cmd/secrets/secrets.go

.PHONY: build
build: api buildApiGateway buildSecrets

.PHONY: buildApiGateway
buildApiGateway:
	@cd svcs/api-gateway && docker build --no-cache -t j4ng5y/hnsecrets:gateway-latest .

.PHONY: buildSecrets
buildSecrets:
	@cd svcs/secrets && docker build --no-cache -t j4ng5y/hnsecrets:secrets-latest .

.PHONY: deploy
deploy: deployApiGateway deploySecrets

.PHONY: deployApiGateway
deployApiGateway:
	@docker push j4ng5y/hnsecrets:gateway-latest

.PHONY: deploySecrets
deploySecrets:
	@docker push j4ng5y/hnsecrets:secrets-latest

run: svcs/api-gateway/bin/gateway svcs/secrets/bin/secrets
	bash -c svcs/api-gateway/bin/gateway
	bash -c svcs/secrets/bin/secrets