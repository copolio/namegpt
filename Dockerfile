FROM golang:1.20-alpine3.18

RUN mkdir /namegpt

ADD . /namegpt

WORKDIR /namegpt

RUN go mod download

RUN go build -o ./bin/app ./cmd/namegpt/main.go

CMD ["/namegpt/bin/app"]