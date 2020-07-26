package json

import (
	"encoding/json"
	"fmt"

	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// DecoderService converts JSON to structs
type DecoderService interface {
	ToInventoryCreateItemVo(json []byte) (*inventory.CreateItemVO, error)
}

// DecoderServiceImpl implements DecoderService
type DecoderServiceImpl struct{}

// Check we implement the interface
var _ DecoderService = &DecoderServiceImpl{}

// NewDecoderServiceImpl is a constructor
func NewDecoderServiceImpl() *DecoderServiceImpl {
	return &DecoderServiceImpl{}
}

// ToInventoryCreateItemVo implements DecoderService
func (d *DecoderServiceImpl) ToInventoryCreateItemVo(bytes []byte) (*inventory.CreateItemVO, error) {
	// TODO: Should use intermediate json vo for tags
	var result inventory.CreateItemVO
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, fmt.Errorf("could not unmarshal to inventory create item vo: %w", err)
	}
	return &result, nil
}
