// Code generated by mockgen; DO NOT EDIT.

package a

type MockDB struct {
	GetFunc func(id string) int
	SetFunc func(id string, v int)
}

func (m *MockDB) Get(id string) int {
	return m.GetFunc(id)
}

func (m *MockDB) Set(id string, v int) {
	m.SetFunc(id, v)
}

