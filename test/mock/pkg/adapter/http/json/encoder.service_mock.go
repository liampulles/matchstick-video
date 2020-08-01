package json

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/http/json"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// MockEncoderService is for mocking
type MockEncoderService struct {
	mock.Mock
}

var _ json.EncoderService = &MockEncoderService{}

// FromInventoryItemView is for mocking
func (d *MockEncoderService) FromInventoryItemView(view *inventory.ViewVO) ([]byte, error) {
	args := d.Called(view)
	return safeArgsGetBytes(args, 0), args.Error(1)
}

// FromInventoryItemViews is for mcoking
func (d *MockEncoderService) FromInventoryItemViews(views []inventory.ViewVO) ([]byte, error) {
	args := d.Called(views)
	return safeArgsGetBytes(args, 0), args.Error(1)
}

func safeArgsGetBytes(args mock.Arguments, idx int) []byte {
	if val, ok := args.Get(idx).([]byte); ok {
		return val
	}
	return nil
}
