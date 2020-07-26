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

type jsonCreateItemVO struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

// ToInventoryCreateItemVo implements DecoderService
func (d *DecoderServiceImpl) ToInventoryCreateItemVo(bytes []byte) (*inventory.CreateItemVO, error) {
	var intermediary jsonCreateItemVO
	if err := json.Unmarshal(bytes, &intermediary); err != nil {
		return nil, fmt.Errorf("could not unmarshal to inventory create item vo: %w", err)
	}

	result := &inventory.CreateItemVO{
		Name:     intermediary.Name,
		Location: intermediary.Location,
	}
	return result, nil
}
