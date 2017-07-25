FROM golang:alpine

EXPOSE 8000/tcp

ENTRYPOINT ["static"]

RUN \
    apk add --update git && \
    rm -rf /var/cache/apk/*

RUN mkdir -p /go/src/static
WORKDIR /go/src/static

COPY . /go/src/static

RUN go get -v -d
RUN go install -v
