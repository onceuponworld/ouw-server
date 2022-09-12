package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)


var (
	conf		= flag.String("conf", APP_CONF, "config file path")
)


var app AppConfig


func main() {

	readConf()

	http.HandleFunc("/world", worldHandler)

	fmt.Printf("Starting %s service on %s...", APP_NAME, addr())
	log.Fatal(http.ListenAndServe(addr(), nil))

} // main
