package mock

import (
	"strings"
)

type Retriver struct {
	Content string
}

func (r *Retriver) Get(url string) string {
	buf := strings.Builder{}
	buf.WriteString(`Content of "`)
	buf.WriteString(url)
	buf.WriteString(`" is "`)
	buf.WriteString(r.Content)
	buf.WriteString(`"`)
	return buf.String()
}
func (r *Retriver) Post(url string, form map[string]string) string {
	r.Content = form["content"]
	return "OK"
}
