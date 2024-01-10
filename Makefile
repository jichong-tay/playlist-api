postgres-image:
	docker pull postgres:16-alpine

postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root foodpanda-playlist

dropdb:
	docker exec -it postgres16 dropdb foodpanda-playlist

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/foodpanda-playlist?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/foodpanda-playlist?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

sqlc:
	sqlc generate

dev-up:
	docker-compose up -d

dev-down:
	docker-compose down

.PHONY: postgres-image postgres createdb dropdb migrateup migratedown sqlc test dev-up dev-down