# Build Stage
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o people-api cmd/app/main.go

# Production Stage 
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/people-api .
COPY --from=builder /app/.env .
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/wait-for-it.sh .


EXPOSE 8080

CMD ["./people-api"]
