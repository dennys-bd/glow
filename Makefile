migrate:
	go run ops/migrate/migrate.go

start:
	go run server.go

prepare:
	docker-compose up -d db pgadmin

install:
	go get

migratetest:
	(\
		export ENVIRONMENT=TEST; \
		make migrate \
	)

test:
	(\
		export ENVIRONMENT=TEST; \
		go test github.com/dennys-bd/glow/usecase \
	)

startdocker:
	docker-compose up backend

migratedocker:
	docker-compose run backend go run ops/migrate/migrate.go

