package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Response struct {
	Message string `json:"message"`
	Hash    string `json:"hash"`
}

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
		w.Header().Set("Content-Type", "application/json;charset=utf-8")

		originUrl := r.FormValue("url")
		if _, err := url.ParseRequestURI(originUrl); err != nil {
			res, _ := json.Marshal(Response{
				Message: "网址格式错误",
			})
			w.Write(res)
			return
		}

		hashs, err := Transform(originUrl)
		if err != nil {
			res, _ := json.Marshal(Response{
				Message: "无法压缩网址",
			})
			w.Write(res)
			return
		}

		now := time.Now()
		hash := hashs[0]
		value := make(map[string]interface{})
		value["longUrl"] = originUrl
		value["createdAt"] = now

		if status := Redis.HMSet(hash, value); status.Err() != nil {
			res, _ := json.Marshal(Response{
				Message: "服务器内部错误",
			})
			w.Write(res)
			return
		}
		Redis.ExpireAt(hash, now.Add(72*time.Hour))
		res, _ := json.Marshal(Response{
			Message: "OK",
			Hash:    hash,
		})
		w.Write(res)
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
