package db

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/db"
)

// MockErrorParser is for mocking
type MockErrorParser struct {
	mock.Mock
}

var _ db.ErrorParser = &MockErrorParser{}

// FromDBRowScan is for mocking
func (s *MockErrorParser) FromDBRowScan(err error, _type string) error {
	args := s.Called(err, _type)
	return args.Error(0)
}
