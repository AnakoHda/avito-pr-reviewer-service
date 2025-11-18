
MIGRATION_DIR = ./migrations
DATABASE = postgres

 .PHONY: migrate-up migrate-down migrate-status clean

bin/goose:
	@ECHO "Installing goose migration tool..."
	@mkdir bin
	@GOBIN=$(PWD)/bin go install github.com/pressly/goose/v3/cmd/goose@v3.25.0

bin/codegen_install:
	@ECHO "Installing goose codegen tool..."
	@mkdir -p bin
	@GOBIN=$(PWD)/bin go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
codegen: bin/codegen_install
	bin/oapi-codegen --config=api/codegen.yaml api/openapi.yml

#migrate-create: bin/goose
#	@./bin/goose -dir $(MIGRATION_DIR) create additional_indexes sql
migrate-up: bin/goose
	@./bin/goose -dir $(MIGRATION_DIR) postgres "$(POSTGRES_URL)" up
migrate-down: bin/goose
	@./bin/goose -dir $(MIGRATION_DIR) postgres "$(POSTGRES_URL)" down
migrate-status: bin/goose
	@./bin/goose -dir $(MIGRATION_DIR) postgres "$(POSTGRES_URL)" status

clean:
	rm -rf bin