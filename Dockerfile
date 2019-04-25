FROM golang:alpine

RUN apk update \
  && apk add git

COPY release/backend_linux_amd64 /go/bin/backend
COPY credentials.json .
COPY .env .
USER nobody:nobody

ENTRYPOINT ["/go/bin/backend"]