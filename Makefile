.PHONY: up down build logs run-backend run-frontend lint fmt clean

up:
	docker-compose up --build -d

down:
	docker-compose down

build:
	docker-compose build

logs:
	docker-compose logs -f


# === Configuration ===
FRONTEND_DIR=frontend
BACKEND_DIR=backend
FRONTEND_PORT=3000
BACKEND_PORT=8080

# === Native (non-Docker) Commands ===

run-backend:
	cd $(BACKEND_DIR) && go run main.go

run-frontend:
	cd $(FRONTEND_DIR) && npm install && npm run dev -- --port $(FRONTEND_PORT)

# === Helpers ===

fmt:
	cd $(BACKEND_DIR) && go fmt ./...

lint:
	cd $(BACKEND_DIR) && golangci-lint run

clean:
	docker-compose down -v --remove-orphans