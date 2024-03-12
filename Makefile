PROJECT_NAME := "tempocerto-api"

include Makefile.application.mk

.PHONY: all dep build clean lint

setup: ## Setup project
	@echo "Get the dependencies..."
	@make dep --silent 
	@echo "Install staticcheck to lint..."
	@go install honnef.co/go/tools/cmd/staticcheck@2022.1.2
	@echo "Install gosec to lint..."
	@go install github.com/securego/gosec/v2/cmd/gosec@v2.13.1
	@echo "Configuring hooks..."
	@git config core.hooksPath hooks/
	@chmod +x ./hooks/pre-commit
	@echo "Done."

all: build

dep-dev-run: ## Run development dependencies
	@docker-compose -f docker-compose.yml up -d  --build --remove-orphans

dep-dev-stop: ## Stop development dependencies
	@docker-compose stop

dep-dev-status: ## Show status from development dependencies
	@docker-compose status

dev-start-with-db: dep-dev-run ## Run application and dependencies
	@docker-compose -f  docker-compose.yml -f docker-compose.app.yml up -d  --build --remove-orphans

dev-down: 
	@docker-compose -f  docker-compose.yml -f docker-compose.app.yml down -v

dev-start-db: dep-dev-run ## Run application and dependencies
	@docker-compose up -d  --build --remove-orphans

security: ## Check security vulnerabilities
	@govulncheck ./...

view-doc: ## Run view doc web application
	@cd docs/c4/tempocerto-systems && c4builder site --watch