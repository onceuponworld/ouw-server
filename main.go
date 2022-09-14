package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/onceuponworld/ouw-sdk"
)


var (
	conf		= flag.String("conf", APP_CONF, "config file path")
)


var rds *redis.Client

var ctx = context.Background()

var app ouwsdk.AppConfig


func initRedis() {

	rds = redis.NewClient(&redis.Options{
		Addr: ouwsdk.Addr(app.Store.Host, app.Store.Port),
		Password: "",
		DB: 0,
	})

} // initRedis


func main() {

	flag.Parse()

	ouwsdk.ParseConfig(*conf, app)

	initRedis()

	http.HandleFunc("/world", worldHandler)

	address := ouwsdk.Addr(app.Host, app.Port)

	fmt.Printf("Starting %s on %s...", APP_NAME, address)

	log.Fatal(http.ListenAndServe(address, nil))

} // main
