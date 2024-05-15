FROM golang:1.22-alpine3.18

WORKDIR /app

COPY . .

RUN go mod download && go mod tidy

RUN go build -o ./ cmd/main.go

ENV BOT_API_CONFIG_PATH="./configs/config.yaml"

CMD ["./main"]