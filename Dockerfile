FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./cmd/finance-tracker-api/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 3000

CMD ["./server", "-storage=postgres"]