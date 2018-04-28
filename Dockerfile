FROM alpine:latest
LABEL maintainer="pixcai<pixcai@163.com>" version="0.1"
RUN mkdir -p /go/src
COPY . /go/src/xurl
WORKDIR /go/src/xurl
EXPOSE 9900
ENTRYPOINT ["./xurl"]
