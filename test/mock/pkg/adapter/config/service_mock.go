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
