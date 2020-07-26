package entity

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// MockInventoryItemConstructor is for mocking
type MockInventoryItemConstructor struct {
	mock.Mock
}

var _ entity.InventoryItemConstructor = &MockInventoryItemConstructor{}

// NewAvailable is for mocking
func (i *MockInventoryItemConstructor) NewAvailable(name string, location string) (entity.InventoryItem, error) {
	args := i.Called(name, location)
	return safeArgsGetInventoryItem(args, 0), args.Error(1)
}

// Reincarnate is for mocking
func (i *MockInventoryItemConstructor) Reincarnate(id entity.ID, name string, location string, available bool) entity.InventoryItem {
	args := i.Called(id, name, location, available)
	return safeArgsGetInventoryItem(args, 0)
}

func safeArgsGetInventoryItem(args mock.Arguments, idx int) entity.InventoryItem {
	if val, ok := args.Get(idx).(entity.InventoryItem); ok {
		return val
	}
	return nil
}
