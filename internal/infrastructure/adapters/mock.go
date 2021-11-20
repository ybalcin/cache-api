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
