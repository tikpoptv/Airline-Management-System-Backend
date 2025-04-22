FROM golang:1.23
WORKDIR /app
COPY . .
RUN go build -o server ./cmd/main.go

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
