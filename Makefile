run:
	go run cmd/api/main.go

migration:
	@read -p "migration file name: " file;\
	go run cmd/migration/main.go -file=$$file

compose-up:
	cp .env.dev .env
	docker compose up -d

compose-down:
	docker-compose down