build:
	cp config/config.json.tmpl config.json && go build -o app cmd/main.go

migration-up:
	migrate -source file://internal/db/migrations/psql-migrations -database "postgres://localhost:5432/sample?sslmode=disable" up

migration-down:
	migrate -source file://internal/db/migrations/psql-migrations -database "postgres://localhost:5432/sample?sslmode=disable" down	

create-migration:
	migrate create -ext sql -dir internal/db/migrations/psql-migrations -seq $(name)