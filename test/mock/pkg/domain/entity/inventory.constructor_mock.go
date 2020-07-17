package entity

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// InventoryItemConstructorMock is for mocking
type InventoryItemConstructorMock struct {
	mock.Mock
}

var _ entity.InventoryItemConstructor = &InventoryItemConstructorMock{}

// NewAvailable is for mocking
func (i *InventoryItemConstructorMock) NewAvailable(name string, location string) entity.InventoryItem {
	args := i.Called(name, location)
	return safeArgsGetInventoryItem(args, 0)
}

// Reincarnate is for mocking
func (i *InventoryItemConstructorMock) Reincarnate(id entity.ID, name string, location string, available bool) entity.InventoryItem {
	args := i.Called(id, name, location, available)
	return safeArgsGetInventoryItem(args, 0)
}

func safeArgsGetInventoryItem(args mock.Arguments, idx int) entity.InventoryItem {
	if val, ok := args.Get(idx).(entity.InventoryItem); ok {
		return val
	}
	return nil
}
