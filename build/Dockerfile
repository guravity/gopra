FROM golang:1.19.1-alpine

RUN apk update && apk add git && apk add bash && apk add --no-cache alpine-sdk build-base

RUN mkdir /go/src/project

WORKDIR /go/src/project

ADD . /go/src/project

RUN go install golang.org/x/tools/cmd/goimports@latest \
    # && go inst
    all golang.org/x/tools/cmd/godoc@latest \
    # && go get -u github.com/kisielk/errcheck@latest