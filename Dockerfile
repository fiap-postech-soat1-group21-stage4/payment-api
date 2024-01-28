FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go build -o paymentapi payment-api/cmd/main.go

CMD ["./paymentapi"]