FROM golang:1.24.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/mesto cmd/main.go

FROM alpine:latest AS runner
COPY --from=builder /app/bin/mesto mesto
EXPOSE 8080
CMD [ "./mesto" ]
