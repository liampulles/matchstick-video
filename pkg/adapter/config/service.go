package config

import (
	"fmt"

	goConfig "github.com/liampulles/go-config"
)

// Store encapsulates configuration properties
// to be injected
type Store interface {
	GetPort() int
}

// StoreImpl implements store
type StoreImpl struct {
	port int
}

// Check we implement the interface
var _ Store = &StoreImpl{}

// NewStoreImpl is a constructor
func NewStoreImpl(source goConfig.Source) (*StoreImpl, error) {
	typedSource := goConfig.NewTypedSource(source)
	store := &StoreImpl{
		port: 8080,
	}

	if err := goConfig.LoadProperties(typedSource,
		goConfig.IntProp("PORT", &store.port, false),
	); err != nil {
		return nil, fmt.Errorf("could not fetch config: %w", err)
	}

	return store, nil
}

// GetPort implements the store interface
func (s *StoreImpl) GetPort() int {
	return s.port
}
