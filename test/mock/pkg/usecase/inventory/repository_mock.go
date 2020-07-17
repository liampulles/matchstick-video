package inventory

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// MockRepository is for mocking
type MockRepository struct {
	mock.Mock
}

var _ inventory.Repository = &MockRepository{}

// Create is for mocking
func (m *MockRepository) Create(e entity.InventoryItem) (entity.ID, error) {
	args := m.Called(e)
	return args.Get(0).(entity.ID), args.Error(1)
}

// FindByID is for mocking
func (m *MockRepository) FindByID(id entity.ID) (entity.InventoryItem, error) {
	args := m.Called(id)
	return safeArgsGetInventoryItem(args, 0), args.Error(1)
}

// Update is for mocking
func (m *MockRepository) Update(e entity.InventoryItem) error {
	args := m.Called(e)
	return args.Error(0)
}

// DeleteByID is for mocking
func (m *MockRepository) DeleteByID(id entity.ID) error {
	args := m.Called(id)
	return args.Error(0)
}

func safeArgsGetInventoryItem(args mock.Arguments, idx int) entity.InventoryItem {
	if val, ok := args.Get(idx).(entity.InventoryItem); ok {
		return val
	}
	return nil
}
