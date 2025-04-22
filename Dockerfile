FROM golang:1.23 as builder
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]
