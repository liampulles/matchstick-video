package config

import (
	"github.com/stretchr/testify/mock"
)

// StoreMock is for mocking
type StoreMock struct {
	mock.Mock
}

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
