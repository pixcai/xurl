# xurl
URL shortener service

# Install dependices
Redis: `go get -u github.com/go-redis/redis`

Mux: `go get -u github.com/gorilla/mux`

# Run
`go run src/main.go src/routes.go src/xurl.go`

And then visit [http://127.0.0.1:9900](http://127.0.0.1:9900) in browser

# Build for Docker
`docker build -t xurl:build -f Dockerfile.build .`

`docker run --rm -v $GOPATH/src:/go/src xurl:build`

# Run in Docker
`docker build -t xurl .`

`docker run -d --name redis -v /data/redis:/data redis:latest redis-server`

`docker run -d --name xurl --link redis -e "XURL_REDIS_URL=redis:6379" xurl`

# License
MIT
