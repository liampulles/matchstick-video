package inventory

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// MockEntityModifier is for mocking
type MockEntityModifier struct {
	mock.Mock
}

var _ inventory.EntityModifier = &MockEntityModifier{}

// ModifyWithUpdateItemVO is for mocking
func (m *MockEntityModifier) ModifyWithUpdateItemVO(e entity.InventoryItem, vo *inventory.UpdateItemVO) error {
	args := m.Called(e, vo)
	return args.Error(0)
}
