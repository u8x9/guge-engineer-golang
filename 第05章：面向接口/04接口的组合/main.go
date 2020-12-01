package main

import (
	"fmt"

	"github.com/u8x9/guge-engineer-golang/第05章：面向接口/04接口的组合/retriver"
	"github.com/u8x9/guge-engineer-golang/第05章：面向接口/04接口的组合/retriver/real"
)

func download(r retriver.Retriver) string {
	return r.Get("https://httpbin.org/get")
}

func post(p retriver.Poster) string {
	return p.Post("https://httpbin.org/post", map[string]string{
		"name":   "u8x9",
		"course": "golang",
	})
}
func session(s retriver.RetriverPoster) string {
	return s.Post("https://httpbin.org/post", map[string]string{
		"name":    "u8x9",
		"course":  "golang",
		"content": "昨夜雨疏风骤，浓睡不消残酒",
	})
}
func main() {
	var r retriver.RetriverPoster
	r = &real.Retriver{}
	fmt.Println(download(r))
	fmt.Println(post(r))
	// fmt.Println(session(r))
}
