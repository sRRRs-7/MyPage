DB_URL="postgresql://root:secret@localhost:5432/mypage?sslmode=disable"

go:
	go run main.go

postgres:
	docker run --name pg_mypage -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it pg_mypage createdb --username=root --owner=root mypage

dropdb:
	docker exec -it pg_mypage dropdb mypage

migrateinit:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate


.PHONY: go postgres createdb dropdb migrateinit migrateup migratedown migrateup1 migratedown1 sqlc