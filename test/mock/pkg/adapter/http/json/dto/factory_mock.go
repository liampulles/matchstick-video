package dto

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/http/json/dto"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// FactoryMock is for mocking
type FactoryMock struct {
	mock.Mock
}

var _ dto.Factory = &FactoryMock{}

// CreateInventoryItemViewFromEntity is for mocking
func (f *FactoryMock) CreateInventoryItemViewFromEntity(e entity.InventoryItem) *dto.InventoryItemView {
	args := f.Called(e)
	return args.Get(0).(*dto.InventoryItemView)
}
