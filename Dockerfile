FROM golang:alpine
WORKDIR /cwk

COPY . /cwk

RUN go build
