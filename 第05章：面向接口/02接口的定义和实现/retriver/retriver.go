package retriver

type Retriver interface {
	Get(url string) string
}
