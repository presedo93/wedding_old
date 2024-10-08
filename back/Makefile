DB_URL=postgresql://rendres:s3cr3t@localhost:5432/wedding?sslmode=disable

createauth:
	docker compose exec postgres createdb --username=rendres --owner=rendres logto

dropauth:
	docker compose exec postgres dropdb -U rendres logto

createdb:
	docker compose exec postgres createdb --username=rendres --owner=rendres wedding

dropdb:
	docker compose exec postgres dropdb -U rendres wedding

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

serve:
	go run main.go

mockdb:
	mockgen -package mockdb -destination db/mock/store.go github.com/presedo93/wedding/back/db/sqlc Store

.PHONY: createauth createdb dropdb migrateup migratedown sqlc test serve mockdb
