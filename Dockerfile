FROM golang:1.14 AS builder
LABEL maintainer="roman.atachiants@gmail.com"

# Copy to GOPATH
RUN mkdir -p /usr/local/go/src/github.com/stripe/veneur/
WORKDIR /usr/local/go/src/github.com/stripe/veneur/
COPY . .
RUN go build ./cmd/veneur/

# Final container
FROM debian AS final
ARG ENV

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /root/
COPY --from=builder /usr/local/go/src/github.com/stripe/veneur/veneur .
COPY --from=builder /usr/local/go/src/github.com/stripe/veneur/_envs/env_$ENV.yaml ./_envs/
CMD [ "/root/veneur" ]
