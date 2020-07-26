package http

import (
	"github.com/stretchr/testify/mock"

	"github.com/liampulles/matchstick-video/pkg/adapter/http"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// MockParameterConverter is for mocking
type MockParameterConverter struct {
	mock.Mock
}

var _ http.ParameterConverter = &MockParameterConverter{}

// ToEntityID is for mocking
func (p *MockParameterConverter) ToEntityID(m map[string]string, param string) (entity.ID, error) {
	args := p.Called(m, param)
	return args.Get(0).(entity.ID), args.Error(1)
}
