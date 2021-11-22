FROM goboring/golang:1.16.7b7

RUN apt-get -y update &&  \
    apt-get -y upgrade && \
    apt-get install --no-install-recommends -y \
        git \
        ca-certificates \   
        apt-utils \
        autoconf \
        automake \
        build-essential \
        git \
        libcurl4-openssl-dev \
        libgeoip-dev \
        liblmdb-dev \
        libpcre++-dev \ 
        libtool \
        libxml2-dev \
        libyajl-dev \
        pkgconf \
        wget \
        zlib1g-dev


ENV GOROOT=/usr/local/go
ENV GOCGO=/usr/local/go-cgo

ENV CGO_ENABLED=1


WORKDIR /test/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN /test/app/bootstrap.sh
ENV LD_LIBRARY_PATH=/usr/local/modsecurity/lib/


RUN go build main.go
ENTRYPOINT /test/app/main
EXPOSE 3080
