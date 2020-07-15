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

// FindByID is for mocking
func (m *MockRepository) FindByID(id entity.ID) (entity.InventoryItem, error) {
	args := m.Called(id)
	return safeArgsGetInventoryItem(args, 0), args.Error(1)
}

// Save is for mocking
func (m *MockRepository) Save(e entity.InventoryItem) (entity.InventoryItem, error) {
	args := m.Called(e)
	return safeArgsGetInventoryItem(args, 0), args.Error(1)
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
