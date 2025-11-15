FROM golang:1.25-alpine

RUN apk add --no-cache git curl bash

WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o publisher .

CMD ["./publisher"]
