package handler

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}
func (e userError) Message() string {
	return string(e)
}

func ListHandler(w http.ResponseWriter, r *http.Request) error {
	if !strings.HasPrefix(r.URL.Path, prefix) {
		return userError("path must start with " + prefix)
	}
	path := r.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	w.Write(buf)
	return nil
}
