package http

import (
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// ParameterConverter converts parameters to various types
type ParameterConverter interface {
	ToEntityID(m map[string]string, param string) (entity.ID, error)
}
