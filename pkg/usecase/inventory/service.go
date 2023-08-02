package inventory

import (
	"fmt"

	"github.com/liampulles/matchstick-video/pkg/adapter/db/sql/inventory"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// Service performs operations on inventories.
type Service interface {
	Create(*CreateItemVO) (entity.ID, error)
	ReadDetails(entity.ID) (*ViewVO, error)
	ReadAll() ([]ThinViewVO, error)
	Update(entity.ID, *UpdateItemVO) error
	Delete(entity.ID) error

	Checkout(entity.ID) error
	CheckIn(entity.ID) error
}

// ServiceImpl implements Service
type ServiceImpl struct {
	entityFactory  EntityFactory
	entityModifier EntityModifier
	voFactory      VOFactory
}

// Make sure ServiceImpl implements Service!
var _ Service = &ServiceImpl{}

// NewServiceImpl is a constructor
func NewServiceImpl(
	entityFactory EntityFactory,
	entityModifier EntityModifier,
	voFactory VOFactory) *ServiceImpl {
	return &ServiceImpl{
		entityFactory:  entityFactory,
		entityModifier: entityModifier,
		voFactory:      voFactory,
	}
}

// Create creates a new entity from a request vo, and persists it.
func (s *ServiceImpl) Create(vo *CreateItemVO) (entity.ID, error) {
	// Create new entity
	e, err := s.entityFactory.CreateFromVO(vo)
	if err != nil {
		return entity.InvalidID, fmt.Errorf("could not create inventory item - factory error: %w", err)
	}

	// Persist it
	id, err := inventory.Create(e)
	if err != nil {
		return entity.InvalidID, fmt.Errorf("could not create inventory item - repository create error: %w", err)
	}

	return id, nil
}

// ReadDetails retrieves an entity and returns a view of it.
func (s *ServiceImpl) ReadDetails(id entity.ID) (*ViewVO, error) {
	// Retrieve entity
	found, err := inventory.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("could not read inventory item - repository find error: %w", err)
	}

	// Create VO
	vo := s.voFactory.CreateViewVOFromEntity(found)

	return vo, nil
}

// ReadAll retrieves all entities and returns views of them.
func (s *ServiceImpl) ReadAll() ([]ThinViewVO, error) {
	// Retrieve entity
	found, err := inventory.FindAll()
	if err != nil {
		return nil, fmt.Errorf("could not read inventory items - repository find error: %w", err)
	}

	// Create VO
	vos := s.voFactory.CreateThinViewVOsFromEntities(found)

	return vos, nil
}

// Update modifies an existing entity as directed by a vo, and
// persists the changes.
func (s *ServiceImpl) Update(id entity.ID, vo *UpdateItemVO) error {
	// Retrieve entity
	found, err := inventory.FindByID(id)
	if err != nil {
		return fmt.Errorf("could not update inventory item - repository find error: %w", err)
	}

	// Modify it
	if err := s.entityModifier.ModifyWithUpdateItemVO(found, vo); err != nil {
		return fmt.Errorf("could not update inventory item - modifier error: %w", err)
	}

	// Persist it
	err = inventory.Update(found)
	if err != nil {
		return fmt.Errorf("could not update inventory item - repository update error: %w", err)
	}
	return nil
}

// Delete wipes the entity from storage.
func (s *ServiceImpl) Delete(id entity.ID) error {
	if err := inventory.DeleteByID(id); err != nil {
		return fmt.Errorf("could not delete inventory item - repository delete error: %w", err)
	}
	return nil
}

// Checkout marks an entity as unavailable, and persists that information.
func (s *ServiceImpl) Checkout(id entity.ID) error {
	// Retrieve entity
	found, err := inventory.FindByID(id)
	if err != nil {
		return fmt.Errorf("could not checkout inventory item - repository find error: %w", err)
	}

	// Checkout the entity
	err = found.Checkout()
	if err != nil {
		return fmt.Errorf("could not checkout inventory item - entity error: %w", err)
	}

	// Persist the updated entity
	err = inventory.Update(found)
	if err != nil {
		return fmt.Errorf("could not checkout inventory item - repository update error: %w", err)
	}
	return nil
}

// CheckIn marks an entity as available, and persists that information.
func (s *ServiceImpl) CheckIn(id entity.ID) error {
	// Retrieve the entity
	found, err := inventory.FindByID(id)
	if err != nil {
		return fmt.Errorf("could not check in inventory item - repository find error: %w", err)
	}

	// Check in the entity
	err = found.CheckIn()
	if err != nil {
		return fmt.Errorf("could not check in inventory item - entity error: %w", err)
	}

	// Persist the modified entity
	err = inventory.Update(found)
	if err != nil {
		return fmt.Errorf("could not check in inventory item - repository update error: %w", err)
	}
	return nil
}
