
.PHONY: start-and-migrate
start-and-migrate: stop-and-delete networkcreate postgres-image postgres wait-for-postgres createdb migrateup

.PHONY: stop-and-delete
stop-and-delete:
	-docker stop postgres16
	-docker rm postgres16
	-docker network rm playlist-network

.PHONY: postgres-image
postgres-image:
	docker pull postgres:16-alpine

.PHONY: postgres
postgres:
	docker run --name postgres16 --network playlist-network -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

.PHONY: wait-for-postgres
wait-for-postgres:postgres
	sleep 10  # wait time to run make createdb

.PHONY: createdb
createdb: wait-for-postgres
	docker exec -it postgres16 createdb --username=root --owner=root foodpanda-playlist

.PHONY: dropdb
dropdb:
	docker exec -it postgres16 dropdb foodpanda-playlist

.PHONY: migrateup
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@127.0.0.1:5433/foodpanda-playlist?sslmode=disable" -verbose up

.PHONY: migrateup1
migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@127.0.0.1:5433/foodpanda-playlist?sslmode=disable" -verbose up 1

.PHONY: migratedown
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@127.0.0.1:5433/foodpanda-playlist?sslmode=disable" -verbose down

.PHONY: migratedown1
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@127.0.0.1:5433/foodpanda-playlist?sslmode=disable" -verbose down 1

.PHONY: test
test:
	go test -v -cover ./...
 
 .PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: server
server:
	go run main.go

.PHONY: mock
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/jichong-tay/playlist-api/db/sqlc Store

.PHONY: networkcreate
networkcreate:
	-docker network create playlist-network

.PHONY: delete-network
delete-network:
	-docker network rm playlist-network

.PHONY: networkconnectdb
networkconnectdb:
	docker network connect playlist-network postgres16

.PHONY: playlist
playlist:
	docker run --name playlist --network playlist-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres16:5432/foodpanda-playlist?sslmode=disable" playlist:latest 

