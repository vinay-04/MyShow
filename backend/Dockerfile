FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

ENV PORT=8080 \
    DB_HOST=postgres \
    DB_PORT=5432 \
    DB_USER=postgres \
    DB_PASSWORD=admin1234 \
    DB_NAME=myshow \
    JWT_SECRET=your-secret-key

EXPOSE 8080

CMD ["./main"]

