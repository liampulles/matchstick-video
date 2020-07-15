package entity

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// InventoryItemMock is for mocking
type InventoryItemMock struct {
	mock.Mock
	// Used to distinguish instances
	Data string
}

var _ entity.InventoryItem = &InventoryItemMock{}

// ID is for mocking
func (i *InventoryItemMock) ID() entity.ID {
	args := i.Called()
	return args.Get(0).(entity.ID)
}

// Name is for mocking
func (i *InventoryItemMock) Name() string {
	args := i.Called()
	return args.String(0)
}

// Location is for mocking
func (i *InventoryItemMock) Location() string {
	args := i.Called()
	return args.String(0)
}

// IsAvailable is for mocking
func (i *InventoryItemMock) IsAvailable() bool {
	args := i.Called()
	return args.Bool(0)
}

// Checkout is for mocking
func (i *InventoryItemMock) Checkout() error {
	args := i.Called()
	return args.Error(0)
}

// CheckIn is for mocking
func (i *InventoryItemMock) CheckIn() error {
	args := i.Called()
	return args.Error(0)
}

// ChangeName is for mocking
func (i *InventoryItemMock) ChangeName(name string) error {
	args := i.Called(name)
	return args.Error(0)
}

// ChangeLocation is for mocking
func (i *InventoryItemMock) ChangeLocation(location string) error {
	args := i.Called(location)
	return args.Error(0)
}
