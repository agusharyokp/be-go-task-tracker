
createdb:
	docker exec -it go-task-tracker-db-1 createdb --username=root --owner=root task-tracker

dropdb:
	docker exec -it go-task-tracker-db-1 dropdb task-tracker

migrateup:
	migrate -path db/migration/ -database "postgresql://root:root@localhost:15432/task-tracker?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://root:root@localhost:15432/task-tracker?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: createdb dropdb migrateup migratedown sqlc