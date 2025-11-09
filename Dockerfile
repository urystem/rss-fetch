FROM golang:1.24.4-alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# статическая сборка
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w -extldflags '-static'" -o rss cmd/main.go

FROM scratch

WORKDIR /app
COPY --from=builder /app/rss .

ENTRYPOINT ["/app/rss", "fetch"]
