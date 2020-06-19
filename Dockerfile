FROM golang:1.14-alpine

USER nobody:nogroup

ENV CGO_ENABLED=0 XDG_CACHE_HOME=/tmp/.cache

WORKDIR /go/src/armstrong
COPY . .
RUN go install -v
