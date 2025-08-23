.PATH: run
run:
	go run cmd/app/main.go

.PATH: build_docker
build_docker:
	go build -o ./bin/migrate cmd/migrator/migrate.go
	go build -o ./bin/app cmd/app/main.go

.PATH: run_build_docker
run_build_docker:
	./migrate --migrations-path=./migrations && ./app

.PATH: migrate
migrate:
	go run cmd/migrator/migrate.go --migrations-path=./migrations

.PATH: delete_and_migrate
delete_and_migrate:
	psql -h localhost -p 5432 -U postgres -d fiber_book -f migrations/1_init.down.sql
	psql -h localhost -p 5432 -U postgres -d fiber_book -f migrations/1_init.up.sql
