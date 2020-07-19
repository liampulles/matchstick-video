package json

import (
	"github.com/liampulles/matchstick-video/pkg/adapter/http/json/dto"
)

// EncoderService converts items to JSON
type EncoderService interface {
	FromInventoryItemView(*dto.InventoryItemView) ([]byte, error)
}
