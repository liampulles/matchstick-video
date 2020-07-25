package config

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/config"
)

// StoreMock is for mocking
type StoreMock struct {
	mock.Mock
}

var _ config.Store = &StoreMock{}

// GetPort is for mocking
func (s *StoreMock) GetPort() int {
	args := s.Called()
	return args.Int(0)
}

// GetDbDriver is for mocking
func (s *StoreMock) GetDbDriver() string {
	args := s.Called()
	return args.String(0)
}

// GetMigrationSource is for mocking
func (s *StoreMock) GetMigrationSource() string {
	args := s.Called()
	return args.String(0)
}
