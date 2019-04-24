FROM golang:alpine

RUN apk update \
  && apk add git

COPY release/backend_linux_amd64 /go/bin/backend
USER nobody:nobody

ENTRYPOINT ["/go/bin/backend"]