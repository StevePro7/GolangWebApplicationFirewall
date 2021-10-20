#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
RUN apk add --no-cache ca-certificates
WORKDIR /test/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /test/app
LABEL Name=20testapi Version=0.0.1
EXPOSE 3080