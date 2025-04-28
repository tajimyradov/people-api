up:
	docker-compose up --build -d

down:
	docker-compose down -v


rebuild:
	docker-compose build people-api


logs:
	docker-compose logs -f people-api


swagger:
	swag init --generalInfo cmd/app/main.go --output docs


migrate-up:
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:7432/people_db?sslmode=disable" up

migrate-down:
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:7432/people_db?sslmode=disable" down

wait-db:
	until pg_isready -h localhost -p 7432 -U postgres; do echo "Waiting for postgres..."; sleep 2; done
