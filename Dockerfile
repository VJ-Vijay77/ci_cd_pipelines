FROM golang:1.19.0-alpine3.16 

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go clean --modcache

RUN go build -o main .

CMD ["/app/main"]



