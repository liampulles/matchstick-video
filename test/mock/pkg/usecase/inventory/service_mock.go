package inventory

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// ServiceMock is for mocking
type ServiceMock struct {
	mock.Mock
}

var _ inventory.Service = &ServiceMock{}

// Create is for mocking
func (s *ServiceMock) Create(vo *inventory.CreateItemVO) (entity.ID, error) {
	args := s.Called(vo)
	return args.Get(0).(entity.ID), args.Error(1)
}

// Read is for mocking
func (s *ServiceMock) Read(id entity.ID) (entity.InventoryItem, error) {
	args := s.Called(id)
	return safeArgsGetInventoryItem(args, 0), args.Error(1)
}

// Update is for mocking
func (s *ServiceMock) Update(id entity.ID, vo *inventory.UpdateItemVO) error {
	args := s.Called(id, vo)
	return args.Error(0)
}

// Delete is for mocking
func (s *ServiceMock) Delete(id entity.ID) error {
	args := s.Called(id)
	return args.Error(0)
}

// IsAvailable is for mocking
func (s *ServiceMock) IsAvailable(id entity.ID) (bool, error) {
	args := s.Called(id)
	return args.Bool(0), args.Error(1)
}

// Checkout is for mocking
func (s *ServiceMock) Checkout(id entity.ID) error {
	args := s.Called(id)
	return args.Error(0)
}

// CheckIn is for mocking
func (s *ServiceMock) CheckIn(id entity.ID) error {
	args := s.Called(id)
	return args.Error(0)
}
