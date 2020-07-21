package config

// Store encapsulates configuration properties
// to be injected
type Store interface {
	GetPort() int
}
