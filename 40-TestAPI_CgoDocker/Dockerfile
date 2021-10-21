FROM golang:1.17.1 AS build

#RUN apk add --no-cache git
#RUN apk add --no-cache ca-certificates

WORKDIR /test/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

#RUN CGO_ENABLED=1 GOOS=linux go build -o /app/bin/server
RUN CGO_ENABLED=1 GOOS=linux go build -o /test/app/awesome
#RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /test/app/awesome -v ./...

ENTRYPOINT /test/app/awesome
EXPOSE 3080