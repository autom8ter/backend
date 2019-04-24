FROM golang:alpine AS builder
COPY ./main.go /go/src/github.com/autom8ter/,placeholder,/main.go
COPY ./vendor /go/src/github.com/autom8ter/,placeholder,/vendor

RUN set -ex && \
  cd /go/src/github.com/autom8ter/,placeholder, && \
  CGO_ENABLED=0 go build \
        -tags netgo \
        -v -a \
        -ldflags '-extldflags "-static"' && \
  mv ./,placeholder, /usr/bin/,placeholder,

FROM busybox
ENV SECRET=,placeholder,
# Retrieve the binary from the previous stage
COPY --from=builder /usr/bin/,placeholder, /usr/local/bin/,placeholder,

# Set the binary as the entrypoint of the container
ENTRYPOINT [ ",placeholder," ]