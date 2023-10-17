package proxy

import (
	"io"
	"log"
	"net/http"
	"time"
)

func Get(url string) string {

	if ok, e := load(url); e == nil {
		log.Println("get content successfully")
		return ok
	}

	res, err := http.Get("https://codesandbox.io/p/sandbox/k2ytyw")

	if err != nil {
		log.Println("start sandbox failed:" + err.Error())
		return ""
	}
	defer res.Body.Close()

	time.Sleep(3 * time.Second)

	txt, _ := load(url)
	return txt

}

func load(target string) (string, error) {
	res, err := http.Get("https://k2ytyw-8080.csb.app?url=" + target)

	if err != nil {
		log.Println("preview failed: " + err.Error())
		return "", err
	}
	defer res.Body.Close()

	txt, _ := io.ReadAll(res.Body)

	return string(txt), nil
}
