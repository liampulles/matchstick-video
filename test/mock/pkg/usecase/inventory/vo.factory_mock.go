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

// CreateViewVOsFromEntities is for mcoking
func (v *MockVOFactory) CreateViewVOsFromEntities(entities []entity.InventoryItem) []inventory.ViewVO {
	args := v.Called(entities)
	return safeArgsGetViewVOs(args, 0)
}

func safeArgsGetViewVO(args mock.Arguments, idx int) *inventory.ViewVO {
	if val, ok := args.Get(idx).(*inventory.ViewVO); ok {
		return val
	}
	return nil
}

func safeArgsGetViewVOs(args mock.Arguments, idx int) []inventory.ViewVO {
	if val, ok := args.Get(idx).([]inventory.ViewVO); ok {
		return val
	}
	return nil
}
