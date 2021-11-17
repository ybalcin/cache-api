package ports

type (
	// Cache output port
	Cache interface {
		Set(key string, value string) error
		Get(key string) (string, error)
		FlushCache()
	}
)
