# My dockerfile
FROM golang:alpine
MAINTAINER Shunsuke Tamiya <shunsuketamiya@posteo.net>

RUN set -e && apk add --no-cache gcc libc-dev libgcc

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

COPY . /go/src/github.com/ablce9/go-assignment

RUN cd /go/src/github.com/ablce9/go-assignment && go install

CMD [ "go-assignment" ]
