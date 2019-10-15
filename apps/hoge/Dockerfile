FROM golang:1.13.1-alpine3.10 as builder

WORKDIR /app
COPY go.mod main.go ./
RUN go build -o main main.go

FROM alpine:3.10.2
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]