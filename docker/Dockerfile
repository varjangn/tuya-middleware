FROM golang:1.22-alpine3.19 as builder

ENV config=docker

WORKDIR /app

COPY ./ /app/

RUN go mod download && \
    apk update && \
    apk add make && \
    make build

FROM golang:1.22-alpine3.19 as runner

COPY --from=builder ./app /app

WORKDIR /app

ENV config=docker

EXPOSE 5000

ENTRYPOINT /app/bin/Tuya-middleware
