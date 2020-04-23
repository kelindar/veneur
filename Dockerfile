FROM golang:1.14 AS builder
LABEL maintainer="roman.atachiants@gmail.com"

# Copy the directory into the container outside of the gopath
RUN mkdir -p /go-build/src/github.com/stripe/veneur/
WORKDIR /go-build/src/github.com/stripe/veneur/
ADD . /go-build/src/github.com/stripe/veneur/

# Download and install any required third party dependencies into the container.
RUN go build -o /go/bin/veneur ./cmd/veneur/

# Base image for runtime
FROM alpine:latest
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /root/
COPY --from=builder /go/bin/veneur .
CMD ["./veneur"]