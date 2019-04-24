FROM golang:alpine

RUN apk update \
  && apk add git

COPY . /go/src/github.com/autom8ter/backend

ENV GO111MODULE=on
RUN cd /go/src/github.com/autom8ter/backend/backend && go build
RUN mv /go/src/github.com/autom8ter/backend/backend/backend /go/bin
# Perform any further action as an unprivileged user.
USER nobody:nobody

ENTRYPOINT ["/go/bin/backend"]