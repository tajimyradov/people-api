

# People API

A RESTful People API built in Go using Gin, PostgreSQL, Swagger, and Docker. Follows Clean Architecture principles.

## Features

- Gin framework
- PostgreSQL database
- Swagger documentation
- Clean architecture
- Dockerized setup
- Makefile for automation

## Prerequisites

- Docker
- Docker Compose
- Make (optional)
- Go (only for local non-Docker usage)

## 📦 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/tajimyradov/people-api.git
cd people-api
```

### 2. Set up environment

Create a `.env` file in the root directory:

```env
SERVER_PORT=8080
DB_HOST=people-db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=people_db

AGE_API_DOMAIN=https://api.agify.io
GENDER_API_DOMAIN=https://api.genderize.io
NATIONALITY_API_DOMAIN=https://api.nationalize.io
```

### 3. Run with Docker

```bash
make up
```

### 4. Stop containers

```bash
make down
```

## 📜 Makefile Commands

```bash
make up            # Build and run the app and database
make down          # Stop and remove containers
make rebuild       # Rebuild the app container
make logs          # Follow logs from containers
make swagger       # Generate Swagger documentation
make migrate-up    # Run database migrations using golang-migrate
```

## 🔍 Swagger Documentation

Once the server is running, visit:

```
http://localhost:8080/swagger/index.html
```

To regenerate Swagger docs:

```bash
make swagger
```

## 🧪 Example API Requests (cURL)

```bash
# Create a person
curl -X POST http://localhost:8080/people \
-H "Content-Type: application/json" \
-d '{"name":"Alice","age":30}'

# List all people
curl http://localhost:8080/people

# Get one person
curl http://localhost:8080/people/1

# Update a person
curl -X PUT http://localhost:8080/people/1 \
-H "Content-Type: application/json" \
-d '{"name":"Bob","age":35}'

# Delete a person
curl -X DELETE http://localhost:8080/people/1
```

## 🧱 Running Migrations Manually

If you're using `golang-migrate` CLI locally:

```bash
migrate -path ./migrations -database "postgres://postgres:postgres@localhost:7432/people_db?sslmode=disable" up
```

## 📁 Project Structure

```
.
├── cmd/app                 # Main app entrypoint
├── config                  # Config loader
├── docs                    # Swagger docs
├── internal
│   ├── handler             # HTTP handlers
│   ├── repository          # DB repository
│   ├── server              # Router setup
│   └── service             # Business logic
├── migrations              # DB migrations
├── pkg                     # Utilities (logger, db)
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```
