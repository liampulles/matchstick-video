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
	FromInventoryItemViews([]inventory.ViewVO) ([]byte, error)
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
	intermediary := mapIntermediary(view)

	bytes, err := json.Marshal(intermediary)
	if err != nil {
		return nil, fmt.Errorf("could not convert inventory item view to json - marshal error: %w", err)
	}
	return bytes, nil
}

// FromInventoryItemViews views to JSON
func (e *EncoderServiceImpl) FromInventoryItemViews(views []inventory.ViewVO) ([]byte, error) {
	intermediaries := make([]jsonViewVO, 0)
	for _, view := range views {
		intermediary := mapIntermediary(&view)
		intermediaries = append(intermediaries, *intermediary)
	}

	bytes, err := json.Marshal(intermediaries)
	if err != nil {
		return nil, fmt.Errorf("could not convert inventory item views to json - marshal error: %w", err)
	}
	return bytes, nil
}

func mapIntermediary(view *inventory.ViewVO) *jsonViewVO {
	return &jsonViewVO{
		ID:        view.ID,
		Name:      view.Name,
		Location:  view.Location,
		Available: view.Available,
	}
}
