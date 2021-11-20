package inmemorystore

type (
	MockInMemoryStore struct {
		SetFn        func(key string, value string) error
		SetFnInvoked bool

		GetFn        func(key string) (string, error)
		GetFnInvoked bool

		FlushFn        func()
		FlushFnInvoked bool
	}
)
