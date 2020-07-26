package inventory

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// VOFactoryMock is for mocking
type VOFactoryMock struct {
	mock.Mock
}

var _ inventory.VOFactory = &VOFactoryMock{}

// CreateViewVOFromEntity is for mocking
func (v *VOFactoryMock) CreateViewVOFromEntity(e entity.InventoryItem) *inventory.ViewVO {
	args := v.Called(e)
	return safeArgsGetViewVO(args, 0)
}

func safeArgsGetViewVO(args mock.Arguments, idx int) *inventory.ViewVO {
	if val, ok := args.Get(idx).(*inventory.ViewVO); ok {
		return val
	}
	return nil
}
