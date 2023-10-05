FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app .

EXPOSE 8888

CMD ["./main"]