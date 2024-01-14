run:
	go run ./cmd/.
test:
	go test --race ./...

rm:
	docker compose down all --volumes --remove-orphans

migrate-new:
	sql-migrate new -config=./environment/mysql/migrate/migrate.yml

migrate-status:
	sql-migrate new -config=./environment/mysql/migrate/migrate.yml

migrate-dev:
	DB_USER=api DB_PASSWORD=P@ssw0rd DB_HOST=127.0.0.1 DB_PORT=3306 DB_NAME_DEV=devdb sql-migrate up -config=./environment/mysql/migrate/migrate.yml -env="development"

migrate-stg:
	DB_USER=root DB_PASSWORD=P@ssw0rd DB_HOST=127.0.0.1 DB_PORT=3306 DB_NAME_DEV=stgdb sql-migrate up -config=./environment/mysql/migrate/migrate.yml -env="staging"

migrate-prod:
	DB_USER=root DB_PASSWORD=P@ssw0rd DB_HOST=127.0.0.1 DB_PORT=3306 DB_NAME_DEV=proddb sql-migrate up -config=./environment/mysql/migrate/migrate.yml -env="production"