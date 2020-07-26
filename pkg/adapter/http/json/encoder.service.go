package json

import (
	"encoding/json"
	"fmt"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// EncoderService converts items to JSON
type EncoderService interface {
	FromInventoryItemView(*inventory.ViewVO) ([]byte, error)
}

// EncoderServiceImpl implements EncoderService
type EncoderServiceImpl struct{}

// Check we implement the interface
var _ EncoderService = &EncoderServiceImpl{}

// NewEncoderServiceImpl is a constructor
func NewEncoderServiceImpl() *EncoderServiceImpl {
	return &EncoderServiceImpl{}
}

type jsonViewVO struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Available bool      `json:"available"`
}

// FromInventoryItemView converts a view to JSON
func (e *EncoderServiceImpl) FromInventoryItemView(view *inventory.ViewVO) ([]byte, error) {
	intermediary := &jsonViewVO{
		ID:        view.ID,
		Name:      view.Name,
		Location:  view.Location,
		Available: view.Available,
	}

	bytes, err := json.Marshal(intermediary)
	if err != nil {
		return nil, fmt.Errorf("could not convert inventory item view to json - marshal error: %w", err)
	}
	return bytes, nil
}
