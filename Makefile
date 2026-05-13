include .env
export 

.PHONY: run test build migrate-up migrate-down migrate-create

run: 
	go run ./cmd/bot

test:
	go test ./...

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up 

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down 1

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

