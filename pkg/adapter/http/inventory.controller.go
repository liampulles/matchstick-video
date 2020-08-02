package http

import (
	"net/http"

	"github.com/liampulles/matchstick-video/pkg/adapter/http/json"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

// TODO: Create controller integration tests

// InventoryControllerImpl defines controller methods
// dealing with the inventory resource.
type InventoryControllerImpl struct {
	inventoryService   inventory.Service
	decoderService     json.DecoderService
	encoderService     json.EncoderService
	responseFactory    ResponseFactory
	parameterConverter ParameterConverter
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
) *InventoryControllerImpl {

	return &InventoryControllerImpl{
		inventoryService:   inventoryService,
		encoderService:     encoderService,
		decoderService:     decoderService,
		responseFactory:    responseFactory,
		parameterConverter: parameterConverter,
	}
}

// GetHandlers implements the Controller interface
func (i *InventoryControllerImpl) GetHandlers() map[HandlerPattern]Handler {
	handlers := make(map[HandlerPattern]Handler)

	addHandler(handlers, http.MethodPost, "/inventory", i.Create)
	addHandler(handlers, http.MethodGet, "/inventory/{id}", i.ReadDetails)
	addHandler(handlers, http.MethodGet, "/inventory", i.ReadAll)
	addHandler(handlers, http.MethodPut, "/inventory/{id}", i.Update)
	addHandler(handlers, http.MethodDelete, "/inventory/{id}", i.Delete)
	addHandler(handlers, http.MethodPut, "/inventory/{id}/checkout", i.Checkout)
	addHandler(handlers, http.MethodPut, "/inventory/{id}/checkin", i.CheckIn)

	return handlers
}

// Create can be called to create an inventory item
func (i *InventoryControllerImpl) Create(request *Request) *Response {
	// Decode JSON request
	vo, err := i.decoderService.ToInventoryCreateItemVo(request.Body)
	if err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	// Delegate to service
	id, err := i.inventoryService.Create(vo)
	if err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	// Create response
	return i.responseFactory.CreateFromEntityID(201, id)
}

// ReadDetails can be called to get details on an inventory item
func (i *InventoryControllerImpl) ReadDetails(request *Request) *Response {
	// Extract ID from path params
	id, err := i.parameterConverter.ToEntityID(request.PathParam, "id")
	if err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	// Delegate to service
	vo, err := i.inventoryService.ReadDetails(id)
	if err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	// Encode to JSON
	json, err := i.encoderService.FromInventoryItemView(vo)
	if err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	// Create response
	return i.responseFactory.CreateJSON(200, json)
}

// ReadAll can be called to get details on all inventory items
func (i *InventoryControllerImpl) ReadAll(request *Request) *Response {
	// Delegate to service
	vos, err := i.inventoryService.ReadAll()
	if err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	// Encode to JSON
	json, err := i.encoderService.FromInventoryItemThinViews(vos)
	if err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	// Create response
	return i.responseFactory.CreateJSON(200, json)
}

// Update can be called to update some of the details
// of an inventory item.
func (i *InventoryControllerImpl) Update(request *Request) *Response {
	// Extract ID from path params
	id, err := i.parameterConverter.ToEntityID(request.PathParam, "id")
	if err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	// Decode JSON request
	vo, err := i.decoderService.ToInventoryUpdateItemVo(request.Body)
	if err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	// Delegate to service
	if err = i.inventoryService.Update(id, vo); err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	// Create response
	return i.responseFactory.CreateEmpty(204)
}

// Delete can be called to remove an inventory item from the system.
func (i *InventoryControllerImpl) Delete(request *Request) *Response {
	// Extract ID from path params
	id, err := i.parameterConverter.ToEntityID(request.PathParam, "id")
	if err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	// Delegate to service
	if err = i.inventoryService.Delete(id); err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	// Create response
	return i.responseFactory.CreateEmpty(204)
}

// Checkout can be called to checkout an inventory item.
func (i *InventoryControllerImpl) Checkout(request *Request) *Response {
	// Extract ID from path params
	id, err := i.parameterConverter.ToEntityID(request.PathParam, "id")
	if err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	// Delegate to service
	if err = i.inventoryService.Checkout(id); err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	// Create response
	return i.responseFactory.CreateEmpty(204)
}

// CheckIn can be called to check in an inventory item.
func (i *InventoryControllerImpl) CheckIn(request *Request) *Response {
	// Extract ID from path params
	id, err := i.parameterConverter.ToEntityID(request.PathParam, "id")
	if err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	// Delegate to service
	if err = i.inventoryService.CheckIn(id); err != nil {
		return i.responseFactory.CreateFromError(err)
	}

	// Create response
	return i.responseFactory.CreateEmpty(204)
}
