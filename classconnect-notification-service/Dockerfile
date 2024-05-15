FROM golang:1.22-alpine3.18

WORKDIR /app

COPY . .

RUN go mod download && go mod tidy

RUN go build -o ./ cmd/main.go

ENV NOTIFICATION_SERVICE_CONFIG_PATH="./configs/config.yaml"

EXPOSE 8083

CMD ["./main"]