.PHONY: all run up down migrate migrate-drop

up:
	@echo "Starting Docker containers"
	docker-compose up -d

down:
	@echo "Stopping Docker containers"
	docker-compose down

migrate:             
	docker exec -i s_db psql -U postgres -d secret_service< migrations/00001_inital_tokens.sql

migrate-drop:
	docker exec -i s_db psql -U postgres -d secret_service< migrations/00002_inital_drop.sql

run:
	@echo "Running server locally"
	go run ./cmd/main.go

all: up 