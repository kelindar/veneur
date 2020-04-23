FROM debian:latest AS builder
LABEL maintainer="roman.atachiants@gmail.com"

# add ca certificates for http secured connection
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/cache/apk/*

# copy the binary
WORKDIR /root/  
RUN go build -o /root/veneur ./cmd/veneur/.
RUN chmod +x /root/veneur

# Expose the port and start the service
EXPOSE 8027
CMD ["/root/veneur"]