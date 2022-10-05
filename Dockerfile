FROM golang:1.19.2-alpine

USER nobody:nogroup

ARG GOOS

ENV CGO_ENABLED=0 XDG_CACHE_HOME=/tmp/.cache

WORKDIR /go/src/armstrong
COPY . .
RUN go install -v
