FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN apk add build-base

ENTRYPOINT ["go", "run", "./cmd", "-S=all", "-C=conf.dev.yaml"]
