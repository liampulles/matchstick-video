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
func (m *MockEntityFactory) CreateFromVO(vo *inventory.CreateItemVO) (*entity.InventoryItem, error) {
	args := m.Called(vo)
	return args.Get(0).(*entity.InventoryItem), args.Error(1)
}
