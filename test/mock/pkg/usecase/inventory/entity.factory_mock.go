package inventory

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// MockEntityFactory is for mocking
type MockEntityFactory struct {
	mock.Mock
}

var _ inventory.EntityFactory = &MockEntityFactory{}

// CreateFromVO is for mocking
func (m *MockEntityFactory) CreateFromVO(vo *inventory.CreateItemVO) (entity.InventoryItem, error) {
	args := m.Called(vo)
	return safeArgsGetInventoryItem(args, 0), args.Error(1)
}

func safeArgsGetInventoryItem(args mock.Arguments, idx int) entity.InventoryItem {
	if val, ok := args.Get(idx).(entity.InventoryItem); ok {
		return val
	}
	return nil
}
