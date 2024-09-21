createauth:
	docker compose exec postgres createdb --username=rendres --owner=rendres auth

createdb:
	docker compose exec postgres createdb --username=rendres --owner=rendres wedding

dropdb:
	docker compose exec postgres dropdb -U rendres wedding

migrateup:
	migrate -path db/migrations -database "postgresql://rendres:s3cr3t@localhost:5432/wedding?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://rendres:s3cr3t@localhost:5432/wedding?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/presedo93/wedding/back/db/sqlc Store

.PHONY: createauth createdb dropdb migrateup migratedown sqlc test server mock
