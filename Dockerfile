FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go build -o server ./cmd/main.go

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
