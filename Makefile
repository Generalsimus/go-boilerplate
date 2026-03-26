# Start your local development server
dev:
	air

# --- Database Commands ---
migrate-new:
	goose -dir db/migrations create $(name) sql

migrate-up:
	goose -dir db/migrations postgres "$(DB_URL)" up

migrate-down:
	goose -dir db/migrations postgres "$(DB_URL)" down

# Tell sqlc EXACTLY where the config file is!
sqlc:
	sqlc generate -f db/sqlc.yaml

# This runs the migration AND updates your Go code in one step
db-update: migrate-up sqlc