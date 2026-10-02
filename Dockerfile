FROM golang:1.26-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./build/alutec-inventory-be ./cmd/api/

FROM alpine:3.19

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/build/alutec-inventory-be /app/alutec-inventory-be
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /app/entrypoint.sh /app/entrypoint.sh

WORKDIR /app

RUN chmod +x entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]
