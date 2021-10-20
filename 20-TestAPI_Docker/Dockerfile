FROM golang:alpine AS builder
RUN apk add --no-cache git
RUN apk add --no-cache ca-certificates

WORKDIR /test/app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /test/app/main -v ./...

ENTRYPOINT /test/app
LABEL Name=20testapi Version=0.0.1
EXPOSE 3080