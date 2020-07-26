package inventory

import (
	"fmt"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// Service performs operations on inventories.
type Service interface {
	Create(*CreateItemVO) (entity.ID, error)
	ReadDetails(entity.ID) (entity.InventoryItem, error)
	Update(entity.ID, *UpdateItemVO) error
	Delete(entity.ID) error

	IsAvailable(entity.ID) (bool, error)
	Checkout(entity.ID) error
	CheckIn(entity.ID) error
}

// ServiceImpl implements Service
type ServiceImpl struct {
	inventoryRepository Repository
	entityFactory       EntityFactory
	entityModifier      EntityModifier
}

// Make sure ServiceImpl implements Service!
var _ Service = &ServiceImpl{}

// NewServiceImpl is a constructor
func NewServiceImpl(
	inventoryRepository Repository,
	entityFactory EntityFactory,
	entityModifier EntityModifier) *ServiceImpl {
	return &ServiceImpl{
		inventoryRepository: inventoryRepository,
		entityFactory:       entityFactory,
		entityModifier:      entityModifier,
	}
}

// Create implements the Service interface
func (s *ServiceImpl) Create(vo *CreateItemVO) (entity.ID, error) {
	// Create new entity
	e, err := s.entityFactory.CreateFromVO(vo)
	if err != nil {
		return entity.InvalidID, fmt.Errorf("could not create inventory item - factory error: %w", err)
	}

	// Persist it
	id, err := s.inventoryRepository.Create(e)
	if err != nil {
		return entity.InvalidID, fmt.Errorf("could not create inventory item - repository error: %w", err)
	}

	return id, nil
}

// ReadDetails implements the Service interface
func (s *ServiceImpl) ReadDetails(id entity.ID) (entity.InventoryItem, error) {
	found, err := s.inventoryRepository.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("could not read inventory item - repository error: %w", err)
	}

	return found, nil
}

// Update implements the Service interface
func (s *ServiceImpl) Update(id entity.ID, vo *UpdateItemVO) error {
	// Retrieve entity
	found, err := s.inventoryRepository.FindByID(id)
	if err != nil {
		return fmt.Errorf("could not update inventory item - repository error: %w", err)
	}

	// Modify it
	if err := s.entityModifier.ModifyWithUpdateItemVO(found, vo); err != nil {
		return fmt.Errorf("could not update inventory item - modifier error: %w", err)
	}

	// Persist it
	err = s.inventoryRepository.Update(found)
	if err != nil {
		return fmt.Errorf("could not update inventory item - repository error: %w", err)
	}
	return nil
}

// Delete implements the Service interface
func (s *ServiceImpl) Delete(id entity.ID) error {
	if err := s.inventoryRepository.DeleteByID(id); err != nil {
		return fmt.Errorf("could not delete inventory item - repository error: %w", err)
	}
	return nil
}

// IsAvailable implements the Service interface
func (s *ServiceImpl) IsAvailable(id entity.ID) (bool, error) {
	found, err := s.inventoryRepository.FindByID(id)
	if err != nil {
		return false, fmt.Errorf("could not determine if inventory item is available - repository error: %w", err)
	}

	return found.IsAvailable(), nil
}

// Checkout implements the Service interface
func (s *ServiceImpl) Checkout(id entity.ID) error {
	// Retrieve entity
	found, err := s.inventoryRepository.FindByID(id)
	if err != nil {
		return fmt.Errorf("could not checkout inventory item - repository error: %w", err)
	}

	// Checkout the entity
	err = found.Checkout()
	if err != nil {
		return fmt.Errorf("could not checkout inventory item - entity error: %w", err)
	}

	// Persist the updated entity
	err = s.inventoryRepository.Update(found)
	if err != nil {
		return fmt.Errorf("could not checkout inventory item - repository error: %w", err)
	}
	return nil
}

// CheckIn implements the Service interface
func (s *ServiceImpl) CheckIn(id entity.ID) error {
	// Retrieve the entity
	found, err := s.inventoryRepository.FindByID(id)
	if err != nil {
		return fmt.Errorf("could not check in inventory item - repository error: %w", err)
	}

	// Check in the entity
	err = found.CheckIn()
	if err != nil {
		return fmt.Errorf("could not check in inventory item - entity error: %w", err)
	}

	// Persist the modified entity
	err = s.inventoryRepository.Update(found)
	if err != nil {
		return fmt.Errorf("could not check in inventory item - repository error: %w", err)
	}
	return nil
}
