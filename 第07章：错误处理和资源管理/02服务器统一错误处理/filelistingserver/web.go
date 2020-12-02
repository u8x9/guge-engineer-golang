package main

import (
	"log"
	"net/http"
	"os"

	"github.com/u8x9/guge-engineer-golang/第07章：错误处理和资源管理/02服务器统一错误处理/filelistingserver/handler"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			log.Println("Error handling request:", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}

			http.Error(w, http.StatusText(code), code)
			return
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(handler.ListHandler))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
