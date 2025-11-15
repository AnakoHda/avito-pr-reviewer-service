
MIGRATION_DIR = ./migrations
DATABASE = postgres
POSTGRES_URL = postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)

 .PHONY: migrate-up migrate-down migrate-status clean

bin/goose:
	@ECHO "Installing goose migration tool..."
	@mkdir bin
	@GOBIN=$(PWD)/bin go install github.com/pressly/goose/v3/cmd/goose@v3.25.0

migrate-up: bin/goose
	@./bin/goose -dir $(MIGRATION_DIR) postgres "$(POSTGRES_URL)" up
migrate-down: bin/goose
	@./bin/goose -dir $(MIGRATION_DIR) postgres "$(POSTGRES_URL)" down
migrate-status: bin/goose
	@./bin/goose -dir $(MIGRATION_DIR) postgres "$(POSTGRES_URL)" status

clean:
	rm -rf bin