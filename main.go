package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // change if running elsewhere
		Password: "",           // no password by default
		DB:       0,            // default DB
	})

	for {
		err := rdb.Set(ctx, "ping", "ok", 0).Err()
		if err == nil {
			break
		}
		fmt.Println("Waiting for Redis...")
		time.Sleep(time.Second)
	}

	err := rdb.Set(ctx, "greeting", "Hello Redis from Go!", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	val, err := rdb.Get(ctx, "greeting").Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Stored value:", val)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := "hello from port 8001"
		w.Write([]byte(response))
	})

	println("listening at port 8001")
	err = http.ListenAndServe(":8001", nil)
	if err != nil {
		panic(err)
	}

}
