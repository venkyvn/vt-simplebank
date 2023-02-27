FROM golang:1.19.6-alpine3.17
WORKDIR /app
COPY . .
RUN go build -o main main.go

EXPOSE 8080
