FROM golang:alpine AS builder
RUN apk add --no-cache git
RUN apk add --no-cache ca-certificates

WORKDIR /test/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /test/app/awesome -v ./...
ENTRYPOINT /test/app/awesome
EXPOSE 3080