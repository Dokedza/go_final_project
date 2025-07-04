FROM golang:1.23.4 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp

FROM alpine:latest
WORKDIR /app/

COPY --from=builder /app .
COPY web ./web
COPY scheduler.db .

ENV TODO_PORT=7540
ENV TODO_DBFILE=./scheduler.db

EXPOSE ${TODO_PORT}

CMD ["./myapp"]