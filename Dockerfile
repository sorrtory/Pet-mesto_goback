FROM golang:1.24.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod verify

# Create tables in db
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/migrate cmd/migrate/migrate.go
# Server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/mesto cmd/main.go

FROM alpine:latest AS runner
COPY --from=builder /app/bin/* .
COPY --from=builder /app/cmd/migrate/migrations/ migrations
COPY --from=builder /app/web/public/ public

EXPOSE 8080
CMD ./migrate up && ./mesto
