package cache

// Service describes a caching service
type Service interface {
	Get(key string) (string, error)
	GetCookies() (interface{}, error)
}
