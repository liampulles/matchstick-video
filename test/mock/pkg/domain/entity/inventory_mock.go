package entity

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// MockInventoryItem is for mocking
type MockInventoryItem struct {
	mock.Mock
	// Used to distinguish instances
	Data string
}

var _ entity.InventoryItem = &MockInventoryItem{}

// ID is for mocking
func (i *MockInventoryItem) ID() entity.ID {
	args := i.Called()
	return args.Get(0).(entity.ID)
}

// Name is for mocking
func (i *MockInventoryItem) Name() string {
	args := i.Called()
	return args.String(0)
}

// Location is for mocking
func (i *MockInventoryItem) Location() string {
	args := i.Called()
	return args.String(0)
}

// IsAvailable is for mocking
func (i *MockInventoryItem) IsAvailable() bool {
	args := i.Called()
	return args.Bool(0)
}

// InitID is for mocking
func (i *MockInventoryItem) InitID(id entity.ID) error {
	args := i.Called(id)
	return args.Error(0)
}

// Checkout is for mocking
func (i *MockInventoryItem) Checkout() error {
	args := i.Called()
	return args.Error(0)
}

// CheckIn is for mocking
func (i *MockInventoryItem) CheckIn() error {
	args := i.Called()
	return args.Error(0)
}

// ChangeName is for mocking
func (i *MockInventoryItem) ChangeName(name string) error {
	args := i.Called(name)
	return args.Error(0)
}

// ChangeLocation is for mocking
func (i *MockInventoryItem) ChangeLocation(location string) error {
	args := i.Called(location)
	return args.Error(0)
}
