FROM golang:latest

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY . /usr/src/app

# Install dependencies (bad way)
RUN go get github.com/KHYehor/gRPCBalancer/src/grpc/calculate
RUN go get github.com/KHYehor/gRPCBalancer/src/grpc/health
RUN go get github.com/KHYehor/gRPCBalancer/src/server
RUN go get google.golang.org/grpc

RUN go build -o lb
CMD ["./lb"]

