migrations_path = migrations

run:
	go run main.go

build_image:
	docker build --tag sm-threads-service

#lint-go:
#	staticcheck ./...

#lint-sql:
#	sqlfluff

# Useage: make create-mig sequence=something
# Output:
# migrate create -ext sql -dir workspaces/database/migrations -seq something
# .../workspaces/database/migrations/000006_something.up.sql
# .../workspaces/database/migrations/000006_something.down.sql
create-mig:
	migrate create -ext sql -dir ${migrations_path} -seq ${sequence}

migrate_threads:
	migrate -database ${DATABASE_URL} -path ${threads_migrations_path}