# Variables
DB_DRIVER = postgres
DB_USER = admin
DB_PASS = admin
DB_NAME = alert-bot-db
DB_HOST = localhost
DB_PORT = 5432

DB_MIGRATIONS_DIR = migrations

DB_CONN_STR = "${DB_DRIVER}://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

migrate-up:
	@echo "Applying database migrations..."
	@migrate -source file://$(DB_MIGRATIONS_DIR) -database $(DB_CONN_STR) up

migrate-down:
	@echo "Reverting database migrations..."
	@migrate -source file://$(DB_MIGRATIONS_DIR) -database $(DB_CONN_STR) down