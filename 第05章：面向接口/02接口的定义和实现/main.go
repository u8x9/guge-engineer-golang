package main

import (
	"fmt"

	"github.com/u8x9/guge-engineer-golang/第05章：面向接口/02接口的定义和实现/retriver"
	"github.com/u8x9/guge-engineer-golang/第05章：面向接口/02接口的定义和实现/retriver/mock"
	"github.com/u8x9/guge-engineer-golang/第05章：面向接口/02接口的定义和实现/retriver/real"
)

func download(r retriver.Retriver) string {
	return r.Get("https://github.com")
}

func getDownload() {
	var r retriver.Retriver

	r = &mock.Retriver{Content: "foobar"}
	fmt.Println(download(r))

	r = &real.Retriver{}
	fmt.Println(download(r))
}

func main() {
	getDownload()
}
