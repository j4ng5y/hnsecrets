all: test build publish

.PHONY: test
test: testApiGateway testSecrets

.PHONY: testApiGateway
testApiGateway:
	@cd svcs/api-gateway
	@go test -race -coverprofile coverage.out ./...

.PHONY: testSecrets
testSecrets:
	@cd svcs/secrets
	@go test -race -coverprofile coverage.out ./...

.PHONY: coverage
coverage:
	@go tool cover -html svcs/api-gateway/coverage.out
	@go tool cover -html svcs/secrets/coverage.out

.PHONY: build
build: buildApiGateway buildSecrets

.PHONY: buildApiGateway
buildApiGateway:
	@cd svcs/api-gateway
	@docker build --no-cache --name j4ng5y/hnsecrets:gateway-latest .

.PHONY: buildSecrets
buildSecrets:
	@cd svcs/secrets
	@docker build --no-cache --name j4ng5y/hnsecrets:secrets-latest .

.PHONY: deploy
deploy: deployApiGateway deploySecrets

.PHONY: deployApiGateway
deployApiGateway:
	@docker push j4ng5y/hnsecrets:gateway-latest

.PHONY: deploySecrets
deploySecrets:
	@docker push j4ng5y/hnsecrets:secrets-latest