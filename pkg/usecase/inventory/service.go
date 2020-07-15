package inventory

import (
	"fmt"

	"github.com/liampulles/matchstick-video/pkg/domain"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// Service performs operations on inventories.
type Service interface {
	Create(*CreateItemVO) (*entity.InventoryItem, error)
	Read(entity.ID) (*entity.InventoryItem, error)
	Update(*UpdateItemVO) error
	Delete(entity.ID) error

	IsAvailable(id entity.ID) (bool, error)
	Checkout(id entity.ID) error
}

// ServiceImpl implements Service
type ServiceImpl struct {
	validator           Validator
	inventoryRepository Repository
	entityFactory       EntityFactory
}

// Make sure ServiceImpl implements Service!
var _ Service = &ServiceImpl{}

// NewServiceImpl is a constructor
func NewServiceImpl(validator Validator,
	inventoryRepository Repository,
	entityFactory EntityFactory) *ServiceImpl {
	return &ServiceImpl{
		validator:           validator,
		inventoryRepository: inventoryRepository,
		entityFactory:       entityFactory,
	}
}

// Create implements the Service interface
func (s *ServiceImpl) Create(vo *CreateItemVO) (*entity.InventoryItem, error) {
	if err := s.validator.ValidateCreateItemVO(vo); err != nil {
		return nil, fmt.Errorf("could not create inventory item - validation error: %w", err)
	}

	entity, err := s.entityFactory.CreateFromVO(vo)
	if err != nil {
		return nil, fmt.Errorf("could not create inventory item - factory error: %w", err)
	}

	saved, err := s.inventoryRepository.Save(entity)
	if err != nil {
		return nil, fmt.Errorf("could not create inventory item - repository error: %w", err)
	}

	return saved, nil
}

// Read implements the Service interface
func (s *ServiceImpl) Read(id entity.ID) (*entity.InventoryItem, error) {
	found, err := s.inventoryRepository.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("could not read inventory item - repository error: %w", err)
	}

	return found, nil
}

// Update implements the Service interface
func (s *ServiceImpl) Update(vo *UpdateItemVO) error {
	if err := s.validator.ValidateUpdateItemVO(vo); err != nil {
		return fmt.Errorf("could not update inventory item - validation error: %w", err)
	}

	found, err := s.inventoryRepository.FindByID(vo.ID)
	if err != nil {
		return fmt.Errorf("could not update inventory item - repository error: %w", err)
	}

	update(found, vo)

	_, err = s.inventoryRepository.Save(found)
	if err != nil {
		return fmt.Errorf("could not update inventory item - repository error: %w", err)
	}
	return nil
}

// Delete implements the Service interface
func (s *ServiceImpl) Delete(id entity.ID) error {
	return &domain.NotImplementedError{
		Package: "inventory",
		Struct:  "ServiceImpl",
		Method:  "Delete",
	}
}

// IsAvailable implements the Service interface
func (s *ServiceImpl) IsAvailable(id entity.ID) (bool, error) {
	return false, &domain.NotImplementedError{
		Package: "inventory",
		Struct:  "ServiceImpl",
		Method:  "IsAvailable",
	}
}

// Checkout implements the Service interface
func (s *ServiceImpl) Checkout(id entity.ID) error {
	return &domain.NotImplementedError{
		Package: "inventory",
		Struct:  "ServiceImpl",
		Method:  "IsAvailable",
	}
}

func update(e *entity.InventoryItem, vo *UpdateItemVO) {
	e.Name = vo.Name
}
