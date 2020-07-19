package json

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/http/json"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// DecoderServiceMock is for mocking
type DecoderServiceMock struct {
	mock.Mock
}

var _ json.DecoderService = &DecoderServiceMock{}

// ToInventoryCreateItemVo is for mocking
func (d *DecoderServiceMock) ToInventoryCreateItemVo(json []byte) (*inventory.CreateItemVO, error) {
	args := d.Called(json)
	return safeArgsGetCreateItemVo(args, 0), args.Error(1)
}

func safeArgsGetCreateItemVo(args mock.Arguments, idx int) *inventory.CreateItemVO {
	if val, ok := args.Get(idx).(*inventory.CreateItemVO); ok {
		return val
	}
	return nil
}
