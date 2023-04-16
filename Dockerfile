FROM golang:1.19 AS builder

WORKDIR /app
COPY . ./

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o /sol-price-bot cmd/bot/main.go

FROM alpine:latest

COPY --from=builder /sol-price-bot /sol-price-bot

EXPOSE 3000

ENTRYPOINT [ "/sol-price-bot" ]