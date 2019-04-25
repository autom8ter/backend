FROM golang:alpine

RUN apk update \
  && apk add git

WORKDIR app
COPY release/backend_linux_amd64 /go/bin/backend
COPY credentials.json .
USER nobody:nobody

ENTRYPOINT ["/go/bin/backend"]