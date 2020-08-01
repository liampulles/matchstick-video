package inventory

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// MockVOFactory is for mocking
type MockVOFactory struct {
	mock.Mock
}

var _ inventory.VOFactory = &MockVOFactory{}

// CreateViewVOFromEntity is for mocking
func (v *MockVOFactory) CreateViewVOFromEntity(e entity.InventoryItem) *inventory.ViewVO {
	args := v.Called(e)
	return safeArgsGetViewVO(args, 0)
}

// CreateThinViewVOsFromEntities is for mocking
func (v *MockVOFactory) CreateThinViewVOsFromEntities(entities []entity.InventoryItem) []inventory.ThinViewVO {
	args := v.Called(entities)
	return safeArgsGetThinViewVOs(args, 0)
}

func safeArgsGetViewVO(args mock.Arguments, idx int) *inventory.ViewVO {
	if val, ok := args.Get(idx).(*inventory.ViewVO); ok {
		return val
	}
	return nil
}

func safeArgsGetThinViewVOs(args mock.Arguments, idx int) []inventory.ThinViewVO {
	if val, ok := args.Get(idx).([]inventory.ThinViewVO); ok {
		return val
	}
	return nil
}
