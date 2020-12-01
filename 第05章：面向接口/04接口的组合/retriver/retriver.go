package retriver

type Retriver interface {
	Get(url string) string
}
type Poster interface {
	Post(url string, form map[string]string) string
}

type RetriverPoster interface {
	Retriver
	Poster
}
