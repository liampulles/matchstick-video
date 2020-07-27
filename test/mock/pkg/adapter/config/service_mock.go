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

// GetDbUser is for mocking
func (s *MockStore) GetDbUser() string {
	args := s.Called()
	return args.String(0)
}

// GetDbPassword is for mocking
func (s *MockStore) GetDbPassword() string {
	args := s.Called()
	return args.String(0)
}

// GetDbHost is for mocking
func (s *MockStore) GetDbHost() string {
	args := s.Called()
	return args.String(0)
}

// GetDbPort is for mocking
func (s *MockStore) GetDbPort() int {
	args := s.Called()
	return args.Int(0)
}

// GetDbName is for mocking
func (s *MockStore) GetDbName() string {
	args := s.Called()
	return args.String(0)
}
