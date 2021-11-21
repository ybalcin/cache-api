package adapters

type (
	MockInMemoryStore struct {
		AddToMemoryFn      func(key string, value string) error
		AddToMemoryInvoked bool

		GetFromMemoryFn      func(key string) (string, error)
		GetFromMemoryInvoked bool

		ClearAllMemoryFn      func()
		ClearAllMemoryInvoked bool

		LoadToMemoryFromFileFn      func()
		LoadToMemoryFromFileInvoked bool

		StartSaveToFileFromMemoryTaskFn      func()
		StartSaveToFileFromMemoryTaskInvoked bool
	}
)

func (m *MockInMemoryStore) AddToMemory(key string, value string) error {
	m.AddToMemoryInvoked = true
	return m.AddToMemoryFn(key, value)
}

func (m *MockInMemoryStore) GetFromMemory(key string) (string, error) {
	m.GetFromMemoryInvoked = true
	return m.GetFromMemoryFn(key)
}

func (m *MockInMemoryStore) ClearAllMemory() {
	m.ClearAllMemoryInvoked = true
	m.ClearAllMemoryFn()
}

func (m *MockInMemoryStore) LoadToMemoryFromFile() {
	m.LoadToMemoryFromFileInvoked = true
	m.LoadToMemoryFromFileFn()
}

func (m *MockInMemoryStore) StartSaveToFileFromMemoryTask() {
	m.StartSaveToFileFromMemoryTaskInvoked = true
	m.StartSaveToFileFromMemoryTaskFn()
}

type (
	MockCacheAdapter struct {
		SetFn      func(key string, value string) error
		SetInvoked bool

		GetFn      func(key string) (string, error)
		GetInvoked bool

		FlushCacheFn func()
		FlushInvoked bool
	}
)

func (m *MockCacheAdapter) Set(key string, value string) error {
	m.SetInvoked = true
	return m.SetFn(key, value)
}

func (m *MockCacheAdapter) Get(key string) (string, error) {
	m.GetInvoked = true
	return m.GetFn(key)
}

func (m *MockCacheAdapter) FlushCache() {
	m.FlushInvoked = true
	m.FlushCacheFn()
}
