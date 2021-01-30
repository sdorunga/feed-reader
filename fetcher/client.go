package fetcher

type Client interface {
	Get(string) (string, error)
}
