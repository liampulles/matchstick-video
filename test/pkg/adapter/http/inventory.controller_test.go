package http_test

import (
	"fmt"
	goHttp "net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"

	httpMocks "github.com/liampulles/matchstick-video/test/mock/pkg/adapter/http"
	jsonMocks "github.com/liampulles/matchstick-video/test/mock/pkg/adapter/http/json"
	inventoryMocks "github.com/liampulles/matchstick-video/test/mock/pkg/usecase/inventory"

	"github.com/liampulles/matchstick-video/pkg/adapter/http"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

type InventoryControllerTestSuite struct {
	suite.Suite
	mockInventoryService *inventoryMocks.ServiceMock
	mockDecoderService   *jsonMocks.DecoderServiceMock
	mockResponseFactory  *httpMocks.ResponseFactoryMock
	sut                  *http.InventoryControllerImpl
}

func TestInventoryControllerTestSuite(t *testing.T) {
	suite.Run(t, new(InventoryControllerTestSuite))
}

func (suite *InventoryControllerTestSuite) SetupTest() {
	suite.mockInventoryService = &inventoryMocks.ServiceMock{}
	suite.mockDecoderService = &jsonMocks.DecoderServiceMock{}
	suite.mockResponseFactory = &httpMocks.ResponseFactoryMock{}
	suite.sut = http.NewInventoryControllerImpl(
		suite.mockInventoryService, suite.mockDecoderService, suite.mockResponseFactory,
	)
}

func (suite *InventoryControllerTestSuite) TestGetHandlers_ShouldReturnAllHandlers() {
	// Setup expectations
	expected := map[http.HandlerPattern]http.Handler{
		http.HandlerPattern{
			Method:      goHttp.MethodPost,
			PathPattern: "/inventory",
		}: suite.sut.Create,
	}

	// Exercise SUT
	actual := suite.sut.GetHandlers()

	// Verify results
	if err := equalKeys(expected, actual); err != nil {
		suite.Failf("Unexpected result.", "%s", err)
	}
}

func (suite *InventoryControllerTestSuite) TestCreate_WhenDecoderServiceFails_ShouldFail() {
	// Setup expectations
	bodyFixture := []byte("some.body")

	// Setup expectations
	expected := &http.Response{
		StatusCode: 501,
		Body:       []byte("some.error"),
	}

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockDecoderService.On("ToInventoryCreateItemVo", bodyFixture).
		Return(nil, mockErr)
	suite.mockResponseFactory.On("CreateFromError", mockErr).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.Create(nil, nil, bodyFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestCreate_WhenInventoryServiceFails_ShouldFail() {
	// Setup expectations
	bodyFixture := []byte("some.body")

	// Setup expectations
	expected := &http.Response{
		StatusCode: 501,
		Body:       []byte("some.error"),
	}

	// Setup mocks
	mockVo := &inventory.CreateItemVO{Name: "some.name"}
	mockErr := fmt.Errorf("mock.error")
	suite.mockDecoderService.On("ToInventoryCreateItemVo", bodyFixture).
		Return(mockVo, nil)
	suite.mockInventoryService.On("Create", mockVo).
		Return(entity.InvalidID, mockErr)
	suite.mockResponseFactory.On("CreateFromError", mockErr).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.Create(nil, nil, bodyFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestCreate_WhenInventoryServicePasses_ShouldReturnEntityResponse() {
	// Setup expectations
	bodyFixture := []byte("some.body")

	// Setup expectations
	expected := &http.Response{
		StatusCode: 201,
		Body:       []byte("some.entity.id"),
	}

	// Setup mocks
	mockVo := &inventory.CreateItemVO{Name: "some.name"}
	mockId := entity.ID(101)
	suite.mockDecoderService.On("ToInventoryCreateItemVo", bodyFixture).
		Return(mockVo, nil)
	suite.mockInventoryService.On("Create", mockVo).
		Return(mockId, nil)
	suite.mockResponseFactory.On("CreateFromEntityID", uint(201), mockId).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.Create(nil, nil, bodyFixture)

	// Verify results
	suite.Equal(expected, actual)
}

// EqualKeys matches the keys of a map
func equalKeys(expected, actual map[http.HandlerPattern]http.Handler) error {
	expectedKeys := keys(expected)
	actualKeys := keys(actual)
	if !reflect.DeepEqual(expectedKeys, actualKeys) {
		return fmt.Errorf("Key mismatch. Expected: %v, Actual: %v", expectedKeys, actualKeys)
	}
	return nil
}

func keys(m map[http.HandlerPattern]http.Handler) []http.HandlerPattern {
	result := make([]http.HandlerPattern, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}
