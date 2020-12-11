FROM golang:1.15

RUN mkdir /app

ADD . /app
WORKDIR /app

RUN go mod download

RUN go build

EXPOSE  80

ENTRYPOINT [ "go", "run", "server/server.go" ]
