package http

import (
	"net/http"

	"github.com/liampulles/matchstick-video/pkg/adapter/http/json"
	"github.com/liampulles/matchstick-video/pkg/adapter/http/json/dto"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// InventoryControllerImpl defines controller methods
// dealing with the inventory resource.
type InventoryControllerImpl struct {
	inventoryService   inventory.Service
	decoderService     json.DecoderService
	encoderService     json.EncoderService
	responseFactory    ResponseFactory
	parameterConverter ParameterConverter
	dtoFactory         dto.Factory
}

// Check we implement the interface
var _ Controller = &InventoryControllerImpl{}

// NewInventoryControllerImpl is a constructor
func NewInventoryControllerImpl(
	inventoryService inventory.Service,
	decoderService json.DecoderService,
	encoderService json.EncoderService,
	responseFactory ResponseFactory,
	parameterConverter ParameterConverter,
	dtoFactory dto.Factory,
) *InventoryControllerImpl {

	return &InventoryControllerImpl{
		inventoryService:   inventoryService,
		encoderService:     encoderService,
		decoderService:     decoderService,
		responseFactory:    responseFactory,
		parameterConverter: parameterConverter,
		dtoFactory:         dtoFactory,
	}
}

// GetHandlers implements the Controller interface
func (i *InventoryControllerImpl) GetHandlers() map[HandlerPattern]Handler {
	handlers := make(map[HandlerPattern]Handler)

	addHandler(handlers, http.MethodPost, "/inventory", i.Create)
	addHandler(handlers, http.MethodGet, "/inventory/{id}", i.ReadDetails)

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

// ReadDetails can be called to get details on an inventory item
func (i *InventoryControllerImpl) ReadDetails(
	pathParam map[string]string,
	queryParam map[string]string,
	body []byte,
) *Response {
	id, err := i.parameterConverter.ToEntityID(pathParam, "id")
	if err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	e, err := i.inventoryService.ReadDetails(id)
	if err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	view := i.dtoFactory.CreateInventoryItemViewFromEntity(e)

	json, err := i.encoderService.FromInventoryItemView(view)
	if err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	return i.responseFactory.Create(200, json)
}
