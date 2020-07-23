package http

import (
	"fmt"
	"strconv"

	"github.com/liampulles/matchstick-video/pkg/domain/entity"
)

// ParameterConverter converts parameters to various types
type ParameterConverter interface {
	ToEntityID(m map[string]string, param string) (entity.ID, error)
}

// ParameterConverterImpl implements ParameterConverter
type ParameterConverterImpl struct{}

// Check we implement the interface
var _ ParameterConverter = &ParameterConverterImpl{}

// NewParameterConverterImpl is a constructor
func NewParameterConverterImpl() *ParameterConverterImpl {
	return &ParameterConverterImpl{}
}

// ToEntityID implements ParameterConverter
func (p *ParameterConverterImpl) ToEntityID(m map[string]string, param string) (entity.ID, error) {
	v, err := getParam(m, param, "entity id")
	if err != nil {
		return entity.InvalidID, err
	}

	i, err := paramToInt64(v, "entity id")
	if err != nil {
		return entity.InvalidID, err
	}
	return entity.ID(i), nil
}

func getParam(m map[string]string, param string, errorType string) (string, error) {
	v, ok := m["id"]
	if !ok {
		return "", fmt.Errorf(
			"could not convert parameters to %s - \"%s\" is not in the parameter list",
			errorType, param)
	}
	return v, nil
}

func paramToInt64(value string, errorType string) (int64, error) {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf(
			"could not convert parameters to %s - cannot convert to int64",
			errorType)
	}
	return i, nil
}
