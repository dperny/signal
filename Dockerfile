FROM golang:1.7-alpine

RUN mkdir -p /go/src/github.com/dperny/signal
WORKDIR /go/src/github.com/dperny/signal

COPY . /go/src/github.com/dperny/signal

CMD ["go", "run", "main.go"]
