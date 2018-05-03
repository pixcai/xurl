# xurl
URL shortener service

### Install dependices
Redis: `go get -u github.com/go-redis/redis`

Mux: `go get -u github.com/gorilla/mux`

### Run
`go run src/main.go src/routes.go src/xurl.go`

And then visit [http://127.0.0.1:9900](http://127.0.0.1:9900) in browser

### Run in Docker
`docker-compose run --rm xurl_build`

`docker-compose up --build xurl`

### License
MIT
