# Build stage
FROM golang:alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o main main.go

# Run stage
FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 3002 4002
CMD ["/app/main"]
