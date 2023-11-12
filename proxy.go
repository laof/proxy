package proxy

import (
	"io"
	"log"
	"net/http"
)

// https://codesandbox.io/p/sandbox/k2ytyw

func Get(target string) string {
	res, err := http.Get("https://k2ytyw-8080.csb.app?url=" + target)

	if err != nil {
		log.Println("preview failed: " + err.Error())
		return ""
	}
	defer res.Body.Close()

	txt, _ := io.ReadAll(res.Body)

	return string(txt)
}
