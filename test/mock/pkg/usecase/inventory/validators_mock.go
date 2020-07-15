package inventory

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// MockValidator is for mocking
type MockValidator struct {
	mock.Mock
}

var _ inventory.Validator = &MockValidator{}

// ValidateCreateItemVO is for mocking
func (m *MockValidator) ValidateCreateItemVO(vo *inventory.CreateItemVO) error {
	args := m.Called(vo)
	return args.Error(0)
}

// ValidateUpdateItemVO is for mocking
func (m *MockValidator) ValidateUpdateItemVO(vo *inventory.UpdateItemVO) error {
	args := m.Called(vo)
	return args.Error(0)
}
