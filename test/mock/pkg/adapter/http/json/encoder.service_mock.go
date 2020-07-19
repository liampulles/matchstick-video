package json

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/http/json"
	"github.com/liampulles/matchstick-video/pkg/adapter/http/json/dto"
)

// EncoderServiceMock is for mocking
type EncoderServiceMock struct {
	mock.Mock
}

var _ json.EncoderService = &EncoderServiceMock{}

// FromInventoryItemView is for mocking
func (d *EncoderServiceMock) FromInventoryItemView(view *dto.InventoryItemView) ([]byte, error) {
	args := d.Called(view)
	return safeArgsGetBytes(args, 0), args.Error(1)
}

func safeArgsGetBytes(args mock.Arguments, idx int) []byte {
	if val, ok := args.Get(idx).([]byte); ok {
		return val
	}
	return nil
}
