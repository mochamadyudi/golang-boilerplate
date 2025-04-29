APP_NAME := yuyuid
VERSION := 1.0.0
BUILD_DIR := dist

PLATFORMS := \
  linux/amd64 \
  linux/arm64 \
  windows/amd64 \
  darwin/amd64 \
  darwin/arm64

MIGRATE_CMD=migrate

DB_DRIVER=postgres
DB_HOST=127.0.0.1
DB_PORT=5432
DB_NAME=db_yuyuid
DB_USER=yuyuid
DB_PASS=
DB_TZ=Asia/Jakarta
DB_SSL_MODE=disable

DB_DSN=$(DB_DRIVER)://$(DB_USER)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)


# migration path
MIGRATION_DIR=$(shell pwd)/internal/database/migrations
SEEDER_DIR=$(shell pwd)/internal/database/migrations

.PHONY: all migrate-up migrate-down migrate-force migrate-drop migrate-new run

dev:
	air
run:
	go run cmd/app/main.go

mu:
	$(MIGRATE_CMD) -path $(MIGRATION_DIR) -database "$(DB_DSN)" up
md:
	$(MIGRATE_CMD) -path $(MIGRATION_DIR) -database "$(DB_DSN)" down
mf:
	$(MIGRATE_CMD) -path $(MIGRATION_DIR) -database "$(DB_DSN)" force $(version)
md:
	$(MIGRATE_CMD) -path $(MIGRATION_DIR) -database "$(DB_DSN)" drop -f
mn:
	$(MIGRATE_CMD) create -ext sql -dir $(MIGRATION_DIR) -seq $(name)
seed:
	go run cmd/seeder/main.go


.PHONY: build clean

build:
	@echo "🔨 Building binaries...";
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
		GOOS=$${platform%%/*}; \
		GOARCH=$${platform##*/}; \
		output_name=$(BUILD_DIR)/$(APP_NAME)-$${GOOS}-$${GOARCH}; \
		if [ "$$GOOS" = "windows" ]; then output_name=$${output_name}.exe; fi; \
		echo ">> Building for $$GOOS/$$GOARCH..."; \
		GOOS=$$GOOS GOARCH=$$GOARCH go build -o $$output_name ./cmd/app/main.go || exit 1; \
	done
	@echo "✅ Build complete.";

clean:
	@echo "🔨 Building binaries...";
	rm -rf $(BUILD_DIR)
