package main

import (
	"github.com/go-redis/redis"
	"log"
	"net/http"
	"net/url"
	"time"
)

var Redis *redis.Client

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr: RedisAddr,
	})
	if _, err := Redis.Ping().Result(); err != nil {
		log.Fatal("xUrlServer: ", err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()

		originUrl, err := url.ParseRequestURI(r.FormValue("url"))
		if err != nil {
			w.Write([]byte("网址格式错误"))
			return
		}

		hashs, err := Transform(originUrl.String())
		if err != nil {
			w.Write([]byte("无法压缩网址"))
			return
		}

		now := time.Now()
		hash := hashs[0]
		value := make(map[string]interface{})
		value["longUrl"] = originUrl.String()
		value["createdAt"] = now

		if status := Redis.HMSet(hash, value); status.Err() != nil {
			w.Write([]byte("服务器内部错误"))
			return
		}
		Redis.ExpireAt(hash, now.Add(72*time.Hour))
		w.Write([]byte(OriginUrl + hash))
	} else {
		http.ServeFile(w, r, "index.html")
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	longUrls, err := Redis.HMGet(r.URL.Path[1:], "longUrl").Result()
	if err != nil || longUrls[0] == nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, longUrls[0].(string), http.StatusMovedPermanently)
	}
}
