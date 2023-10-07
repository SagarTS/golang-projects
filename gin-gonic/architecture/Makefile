
include .env
MIGRATE=migrate -path migration -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -verbose
dev:
		gin appPort ${ServerPort} -i run server.go
migrate-up:
		$(MIGRATE) up
migrate-down:
		$(MIGRATE) down 
force:
		@read -p  "Which version do you want to force?" VERSION; \
		$(MIGRATE) force $$VERSION

goto:
		@read -p  "Which version do you want to migrate?" VERSION; \
		$(MIGRATE) goto $$VERSION

drop:
		$(MIGRATE) drop

create:
		@read -p  "What is the name of migration?" NAME; \
		${MIGRATE} create -ext sql -dir migration  $$NAME

.PHONY: migrate-up migrate-down force goto drop create