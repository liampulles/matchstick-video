package http

import (
	"net/http"

	"github.com/liampulles/matchstick-video/pkg/adapter/http/json"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// InventoryControllerImpl defines controller methods
// dealing with the inventory resource.
type InventoryControllerImpl struct {
	inventoryService inventory.Service
	decoderService   json.DecoderService
	responseFactory  ResponseFactory
}

// Check we implement the interface
var _ Controller = &InventoryControllerImpl{}

// NewInventoryControllerImpl is a constructor
func NewInventoryControllerImpl(
	inventoryService inventory.Service,
	decoderService json.DecoderService,
	responseFactory ResponseFactory,
) *InventoryControllerImpl {
	return &InventoryControllerImpl{
		inventoryService: inventoryService,
		decoderService:   decoderService,
		responseFactory:  responseFactory,
	}
}

// GetHandlers implements the Controller interface
func (i *InventoryControllerImpl) GetHandlers() map[HandlerPattern]Handler {
	handlers := make(map[HandlerPattern]Handler)

	addHandler(handlers, http.MethodPost, "/inventory", i.Create)

	return handlers
}

// Create can be called to create an inventory item
func (i *InventoryControllerImpl) Create(
	pathParam map[string]string,
	queryParam map[string]string,
	body []byte,
) *Response {

	vo, err := i.decoderService.ToInventoryCreateItemVo(body)
	if err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	id, err := i.inventoryService.Create(vo)
	if err != nil {
		return i.responseFactory.CreateFromError(err)
	}
	return i.responseFactory.CreateFromEntityID(201, id)
}
