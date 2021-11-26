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


# WAF
RUN cd /tmp
RUN git clone https://github.com/SpiderLabs/ModSecurity.git
RUN cd ModSecurity
#RUN git checkout ${MODSEC_VERSION:-v3.0.4}  # not work
#RUN git submodule update --init
CMD /build.sh

#ENV LD_LIBRARY_PATH=/usr/local/modsecurity/lib/





RUN echo 'ping localhost &' > /bootstrap.sh
RUN echo 'sleep infinity' >> /bootstrap.sh
RUN chmod +x /bootstrap.sh

CMD /bootstrap.sh