#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
RUN apk add --no-cache ca-certificates
WORKDIR /stevepro/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

ENTRYPOINT /stevepro/app
EXPOSE 3080
ENTRYPOINT ["/stevepro/app/run.sh"]