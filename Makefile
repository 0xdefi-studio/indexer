run_postgres:
	DB_PORT=5432 docker compose --env-file .env.dev.local up db

run_postgres_bg:
	DB_PORT=5432 docker compose --env-file .env.dev.local up db -d

run_fomo3d:
	docker compose  up blend_indexer

