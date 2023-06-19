#!make
include .env

DB_URL=postgresql://${M_DB_USERNAME}:${M_DB_PASSWORD}@localhost:5432/${M_DB_NAME}?sslmode=disable


migrateUp:
	migrate -path db/migration -database "${DB_URL}" -verbose up

migrateDown:
	migrate -path db/migration -database "${DB_URL}" -verbose down

migrateForce: 
	migrate -path migrations -database "${DB_URL}" force $(version)

migrateCreate:
	migrate create -ext sql -dir migrations -seq $(fileName)

buildDocker:
	docker-compose up --build

runDocker:
	docker-compose up -d

closeDocker:
	docker-compose down

dlog:
	docker-compose logs

sqlcGen:
	sqlc generate

test:
	go test -v -cover ./...
.PHONY: migrateUp migrateDown buildDocker runDocker closeDocker dlog test