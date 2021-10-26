FROM goboring/golang:1.16.7b7
RUN apt-get -y update &&  \
    apt-get -y upgrade && \
    apt-get install --no-install-recommends -y \
        git \
        ca-certificates

WORKDIR /test/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /test/app/awesome -v ./...
ENTRYPOINT /test/app/awesome
EXPOSE 3080