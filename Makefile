compose-up:
	cp .env.dev .env
	docker compose up -d

compose-down:
	docker-compose down