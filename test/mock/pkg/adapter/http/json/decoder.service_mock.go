package json

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/http/json"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// MockDecoderService is for mocking
type MockDecoderService struct {
	mock.Mock
}

var _ json.DecoderService = &MockDecoderService{}

// ToInventoryCreateItemVo is for mocking
func (d *MockDecoderService) ToInventoryCreateItemVo(json []byte) (*inventory.CreateItemVO, error) {
	args := d.Called(json)
	return safeArgsGetCreateItemVo(args, 0), args.Error(1)
}

// ToInventoryUpdateItemVo is for mocking
func (d *MockDecoderService) ToInventoryUpdateItemVo(json []byte) (*inventory.UpdateItemVO, error) {
	args := d.Called(json)
	return safeArgsGetUpdateItemVo(args, 0), args.Error(1)
}

func safeArgsGetCreateItemVo(args mock.Arguments, idx int) *inventory.CreateItemVO {
	if val, ok := args.Get(idx).(*inventory.CreateItemVO); ok {
		return val
	}
	return nil
}

func safeArgsGetUpdateItemVo(args mock.Arguments, idx int) *inventory.UpdateItemVO {
	if val, ok := args.Get(idx).(*inventory.UpdateItemVO); ok {
		return val
	}
	return nil
}
