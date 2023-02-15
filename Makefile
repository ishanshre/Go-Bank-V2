#!make
include .env

DB_URL=postgresql://${DB_USERNAME}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable


migrateUp:
	migrate -path db/migration -database "${DB_URL}" -verbose up

buildDocker:
	docker-compose up -d --build

runDocker:
	docker-compose up -d

closeDocker:
	docker-compose down

dlog:
	docker-compose logs