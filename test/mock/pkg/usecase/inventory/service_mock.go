package inventory

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// MockService is for mocking
type MockService struct {
	mock.Mock
}

var _ inventory.Service = &MockService{}

// Create is for mocking
func (s *MockService) Create(vo *inventory.CreateItemVO) (entity.ID, error) {
	args := s.Called(vo)
	return args.Get(0).(entity.ID), args.Error(1)
}

// ReadDetails is for mocking
func (s *MockService) ReadDetails(id entity.ID) (*inventory.ViewVO, error) {
	args := s.Called(id)
	return safeArgsGetViewVO(args, 0), args.Error(1)
}

// ReadAll is for mocking
func (s *MockService) ReadAll() ([]inventory.ThinViewVO, error) {
	args := s.Called()
	return safeArgsGetThinViewVOs(args, 0), args.Error(1)
}

// Update is for mocking
func (s *MockService) Update(id entity.ID, vo *inventory.UpdateItemVO) error {
	args := s.Called(id, vo)
	return args.Error(0)
}

// Delete is for mocking
func (s *MockService) Delete(id entity.ID) error {
	args := s.Called(id)
	return args.Error(0)
}

// Checkout is for mocking
func (s *MockService) Checkout(id entity.ID) error {
	args := s.Called(id)
	return args.Error(0)
}

// CheckIn is for mocking
func (s *MockService) CheckIn(id entity.ID) error {
	args := s.Called(id)
	return args.Error(0)
}
