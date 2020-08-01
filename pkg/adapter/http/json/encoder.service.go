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
	FromInventoryItemThinViews([]inventory.ThinViewVO) ([]byte, error)
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

type jsonThinViewVO struct {
	ID   entity.ID `json:"id"`
	Name string    `json:"name"`
}

// FromInventoryItemView converts a view to JSON
func (e *EncoderServiceImpl) FromInventoryItemView(view *inventory.ViewVO) ([]byte, error) {
	intermediary := mapViewIntermediary(view)

	bytes, err := json.Marshal(intermediary)
	if err != nil {
		return nil, fmt.Errorf("could not convert inventory item view to json - marshal error: %w", err)
	}
	return bytes, nil
}

// FromInventoryItemThinViews views to JSON
func (e *EncoderServiceImpl) FromInventoryItemThinViews(views []inventory.ThinViewVO) ([]byte, error) {
	intermediaries := make([]jsonThinViewVO, 0)
	for _, view := range views {
		intermediary := mapThinViewIntermediary(&view)
		intermediaries = append(intermediaries, *intermediary)
	}

	bytes, err := json.Marshal(intermediaries)
	if err != nil {
		return nil, fmt.Errorf("could not convert inventory item views to json - marshal error: %w", err)
	}
	return bytes, nil
}

func mapViewIntermediary(view *inventory.ViewVO) *jsonViewVO {
	return &jsonViewVO{
		ID:        view.ID,
		Name:      view.Name,
		Location:  view.Location,
		Available: view.Available,
	}
}

func mapThinViewIntermediary(view *inventory.ThinViewVO) *jsonThinViewVO {
	return &jsonThinViewVO{
		ID:   view.ID,
		Name: view.Name,
	}
}
