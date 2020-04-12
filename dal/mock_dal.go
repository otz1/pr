package dal

import "log"

type MockResultCacheDAL struct {
	BasicResultCacheDAL
}

func NewMockResultCacheDAL() *MockResultCacheDAL {
	return &MockResultCacheDAL{
		NewBasicResultCacheDAL(),
	}
}

func (m *MockResultCacheDAL) Query(query string) {
	key := m.Hash(query)
	log.Println("looking in mock for", key)
}