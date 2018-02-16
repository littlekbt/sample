FROM ubuntu

MAINTAINER littlekbt

RUN apt update && \
    apt install -y git && \
    git clone https://github.com/nginx/nginx && \
    cd ./nginx && \
    apt install -y libpcre3 libpcre3-dev zlib1g-dev libssl-dev gcc make && \
    ./auto/configure --with-http_v2_module --with-http_ssl_module && \
    make

WORKDIR /go
