package real

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type Retriver struct {
	UserAgent string
	Timeout   time.Duration
}

func (r *Retriver) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	return string(result)
}

func (r *Retriver) Post(urlStr string, form map[string]string) string {
	data := url.Values{}
	for k, v := range form {
		data.Set(k, v)
	}
	resp, err := http.PostForm(urlStr, data)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	return string(result)
}
