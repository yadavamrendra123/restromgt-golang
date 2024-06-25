
stop:
	docker-compose down
start:
	docker-compose up
rerun:
	docker-compose down
	docker-compose up
build:
	docker-compose build
tidy:
	go mod tidy
help:
	echo "Usage: make [command]"; \
	echo ""; \
	echo "Available commands:"; \
	echo "  stop     - Stop Docker containers (docker-compose down)"; \
	echo "  start    - Start Docker containers (docker-compose up)"; \
	echo "  rerun    - Stop and then start Docker containers"; \
	echo "  build    - Build Docker images (docker-compose build)"; \
	echo "  tidy     - Tidy Go modules (go mod tidy)"; \
	echo "  help     - Display this help message"
