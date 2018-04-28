FROM alpine:latest
LABEL maintainer="pixcai<pixcai@163.com>" version="0.1"
RUN mkdir -p /go/bin
COPY xurl /go/bin
EXPOSE 9900
ENTRYPOINT ["/go/bin/xurl"]
