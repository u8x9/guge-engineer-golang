package handler

import (
	"io/ioutil"
	"net/http"
	"os"
)

func ListHandler(w http.ResponseWriter, r *http.Request) error {
	path := r.URL.Path[len("/list/"):]
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
