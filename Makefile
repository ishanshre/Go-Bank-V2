#!make
include .env

DB_URL=postgresql://${M_DB_USERNAME}:${M_DB_PASSWORD}@localhost:5432/${M_DB_NAME}?sslmode=disable


migrateUp:
	migrate -path db/migration -database "${DB_URL}" -verbose up

migrateDown:
	migrate -path db/migration -database "${DB_URL}" -verbose down

buildDocker:
	docker-compose up --build

runDocker:
	docker-compose up -d

closeDocker:
	docker-compose down

dlog:
	docker-compose logs

.PHONY: migrateUp migrateDown buildDocker runDocker closeDocker dlog