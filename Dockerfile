FROM alpine:latest
MAINTAINER pixcai "pixcai@163.com"
COPY xurl /usr/bin/
EXPOSE 9900
ENTRYPOINT ["xurl"]
