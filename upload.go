package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"upload/cache"
	"upload/context/global"
	"upload/control"
)

var configFile = flag.String("conf", "config.json", "config file")

func main() {
	// readin bind address
	flag.Parse()

	// parser config variables
	global.Parser(*configFile)

	// init control and cache
	control.InitRoute()
	cache.InitTemplates()

	// startup service
	log.Println("ListenAndServe: address", global.Conf.BindPort)
	addr := fmt.Sprintf(":%d", global.Conf.BindPort)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
