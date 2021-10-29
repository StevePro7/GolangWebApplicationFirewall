FROM debian:10-slim
RUN apt-get update && apt-get install -y gcc

FROM golang:alpine AS builder
RUN apk add --no-cache git
RUN apk add --no-cache ca-certificates


WORKDIR /test/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

#RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o /test/app/awesome -v ./...
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /test/app/awesome -v ./...
ENTRYPOINT /test/app/awesome
EXPOSE 3080