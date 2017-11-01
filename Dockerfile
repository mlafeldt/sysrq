FROM golang:alpine

ADD . /go/src/github.com/mlafeldt/sysrq
WORKDIR /go/src/github.com/mlafeldt/sysrq

RUN go install -v ./cmd/sysrq

ENTRYPOINT ["sysrq"]
