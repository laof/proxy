package main

import (
	"log"

	"github.com/laof/proxy"
)

func main() {
	txt := proxy.Get("http://www.baidu.com")
	log.Println(txt)
}
