FROM golang:alpine AS builder
RUN apk add --no-cache git
RUN apk add --no-cache ca-certificates

WORKDIR /test/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

ENTRYPOINT /test/app
LABEL Name=20testapi Version=0.0.1
EXPOSE 3080
CMD ["./main"]