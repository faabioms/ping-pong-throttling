FROM golang:1.16-alpine

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN go build -o ping_pong_app .

CMD /app/ping_pong_app