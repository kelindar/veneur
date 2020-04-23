FROM golang:1.13 AS builder
LABEL maintainer="roman.atachiants@gmail.com"

# Copy the directory into the container outside of the gopath
RUN mkdir -p /go/src/github.com/stripe/veneur/
WORKDIR /go/src/github.com/stripe/veneur/
ADD . /go/src/github.com/stripe/veneur/

# Download and install any required third party dependencies into the container.
ENV GO111MODULE off
RUN go build -o /go/bin/veneur ./cmd/veneur/main.go

# Base image for runtime
FROM debian:latest
RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /root/
COPY --from=builder /go/bin/veneur .
COPY --from=builder /go/src/github.com/stripe/veneur/example.yaml ./config.yaml
RUN chmod +x ./veneur
CMD ["./veneur", "-f", "config.yaml"]