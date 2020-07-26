package config

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/config"
)

// MockStore is for mocking
type MockStore struct {
	mock.Mock
}

var _ config.Store = &MockStore{}

// GetPort is for mocking
func (s *MockStore) GetPort() int {
	args := s.Called()
	return args.Int(0)
}

// GetDbDriver is for mocking
func (s *MockStore) GetDbDriver() string {
	args := s.Called()
	return args.String(0)
}

// GetMigrationSource is for mocking
func (s *MockStore) GetMigrationSource() string {
	args := s.Called()
	return args.String(0)
}
