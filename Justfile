COMPOSE_FILE := "deploy/compose.yaml"
PROJECT_NAME := "group_manager"
ENV_FILE := ".env"

COMPOSE_CMD := "docker compose -f " + COMPOSE_FILE + " -p " + PROJECT_NAME + " --env-file " + ENV_FILE

sqlc:
    sqlc generate

postgres:
    docker exec -it group_manager_postgres psql --username group_manager

build:
    docker build -t group-manager:1.0.0 -f Dockerfile .

api:
    go run cmd/api/main.go

lint:
    golangci-lint run

up:
    {{ COMPOSE_CMD }} up -d

down:
    {{ COMPOSE_CMD }} down

restart:
    {{ COMPOSE_CMD }} restart

logs:
    {{ COMPOSE_CMD }} logs -f

status:
    {{ COMPOSE_CMD }} ps
