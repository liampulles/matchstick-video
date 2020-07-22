package json

import (
	"encoding/json"
	"fmt"

	"github.com/liampulles/matchstick-video/pkg/adapter/http/json/dto"
)

// EncoderService converts items to JSON
type EncoderService interface {
	FromInventoryItemView(*dto.InventoryItemView) ([]byte, error)
}

// EncoderServiceImpl implements EncoderService
type EncoderServiceImpl struct{}

// Check we implement the interface
var _ EncoderService = &EncoderServiceImpl{}

// NewEncoderServiceImpl is a constructor
func NewEncoderServiceImpl() *EncoderServiceImpl {
	return &EncoderServiceImpl{}
}

// FromInventoryItemView implements EncoderService
func (e *EncoderServiceImpl) FromInventoryItemView(view *dto.InventoryItemView) ([]byte, error) {
	bytes, err := json.Marshal(view)
	if err != nil {
		return nil, fmt.Errorf("could not convert inventory item view to json - marshal error: %w", err)
	}
	return bytes, nil
}
