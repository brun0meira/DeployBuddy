# Default environment file
ENV_FILE := .env

# Docker compose files
COMPOSE_DEV := compose.dev.yml
COMPOSE_PROD := compose.prod.yml

# Enhanced cleanup function
define docker_cleanup
	echo "Cleaning up Docker containers and volumes..."; \
	docker stop $(1) || true; \
	docker rm $(1) || true; \
	docker compose -f $(2) down -v
endef

# Ensure .env file exists
$(ENV_FILE):
	@if [ ! -f $(ENV_FILE) ]; then \
		echo "Creating .env from .env.example"; \
		cp .env.example $(ENV_FILE); \
	else \
		echo ".env file already exists"; \
	fi

# General targets to manage the database containers
.PHONY: manage-db-dev
manage-db-dev:
	@echo "Starting $(CONTAINER_DEV)..."
	docker compose -f $(COMPOSE_DEV) down
	docker compose -f $(COMPOSE_DEV) up -d

.PHONY: manage-db-prod
manage-db-prod:
	@echo "Starting $(CONTAINER_PROD)..."
	docker compose -f $(COMPOSE_PROD) down
	docker compose -f $(COMPOSE_PROD) up -d

# Development environment setup
.PHONY: dev
dev: $(ENV_FILE) manage-db-dev
	@echo "Stopping any process on port 8080..."
	@npx kill-port 8080
	@echo "Tidying up Go modules..."
	@go mod tidy
	@echo "Starting the development server using Air..."
	@air

# Production environment setup
.PHONY: prod
prod: $(ENV_FILE) manage-db-prod
	@npx kill-port 8080
	@docker compose -f compose.prod.yml up --build -d

# Reset - stops and removes the database container
.PHONY: reset-dev
reset-dev:
	@$(call docker_cleanup,$(CONTAINER_NAME_DEV),compose.dev.yml)

.PHONY: reset-prod
reset-prod:
	@$(call docker_cleanup,$(CONTAINER_NAME_PROD),compose.prod.yml)

# Test - runs the tests with database container | Run mod tidy, start in dev mode, run tests and clean up
.PHONY: test
test: manage-db-dev
	@go mod tidy
	@go test -v ./...
	@$(call docker_cleanup,$(CONTAINER_NAME_DEV),compose.dev.yml)

# Deploy - deploys the application
.PHONY: deploy
deploy:
	echo "Deploying application..."; \
	scp -i $(SSH_KEY) ~/.ssh/id_ed25519 ubuntu@$(IP):~/.ssh/id_ed25519; \
	chmod +x deploy.sh; \
	./deploy.sh $(SSH_KEY) $(IP) $(CLONE_SSH_KEY) || { echo "Deployment failed"; exit 1; }

# Build - builds the application
.PHONY: build
build:
	go build -o bin/app cmd/app/main.go

# Cleaning up the project
.PHONY: clean
clean:
	@echo "Cleaning up containers..."
	docker compose -f $(COMPOSE_DEV) down -v
	docker compose -f $(COMPOSE_PROD) down -v
	@echo "Removing temporary files..."
	@rm -rf bin .air tmp go.sum vendor .DS_Store .vscode .idea

# Swagger - swag init --parseDependency --parseInternal --generalInfo ./cmd/app/main.go --output ./docs
.PHONY: swag
swag:
	@echo "Generating Swagger documentation..."
	@swag init --parseDependency --parseInternal --generalInfo ./cmd/app/main.go --output ./docs
