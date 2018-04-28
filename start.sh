#!/usr/bin/env bash

docker rmi -f xurl:0.1 xurl:build > /dev/null

docker build -t xurl:build -f Dockerfile.build .
docker run --rm -v $GOPATH/src:/go/src xurl:build

docker build -t xurl:0.1 .
clear
docker run -d --rm --name xurl-0.1 --link redis -e XURL_REDIS_URL="redis:6379" -p 9900:9900 xurl:0.1
