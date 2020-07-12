FROM golang:latest

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY . /usr/src/app

RUN go get -u -v -f all
CMD go build -o lb
ENTRYPOINT ["./lb"]

