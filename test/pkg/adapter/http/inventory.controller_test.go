package http_test

import (
	"fmt"
	goHttp "net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"

	httpMocks "github.com/liampulles/matchstick-video/test/mock/pkg/adapter/http"
	jsonMocks "github.com/liampulles/matchstick-video/test/mock/pkg/adapter/http/json"
	dtoMocks "github.com/liampulles/matchstick-video/test/mock/pkg/adapter/http/json/dto"
	entityMocks "github.com/liampulles/matchstick-video/test/mock/pkg/domain/entity"
	inventoryMocks "github.com/liampulles/matchstick-video/test/mock/pkg/usecase/inventory"

	"github.com/liampulles/matchstick-video/pkg/adapter/http"
	"github.com/liampulles/matchstick-video/pkg/adapter/http/json/dto"
	"github.com/liampulles/matchstick-video/pkg/domain/entity"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

type InventoryControllerTestSuite struct {
	suite.Suite
	mockInventoryService   *inventoryMocks.ServiceMock
	mockDecoderService     *jsonMocks.DecoderServiceMock
	mockEncoderService     *jsonMocks.EncoderServiceMock
	mockResponseFactory    *httpMocks.ResponseFactoryMock
	mockParameterConverter *httpMocks.ParameterConverterMock
	mockDtoFactory         *dtoMocks.FactoryMock
	sut                    *http.InventoryControllerImpl
}

func TestInventoryControllerTestSuite(t *testing.T) {
	suite.Run(t, new(InventoryControllerTestSuite))
}

func (suite *InventoryControllerTestSuite) SetupTest() {
	suite.mockInventoryService = &inventoryMocks.ServiceMock{}
	suite.mockDecoderService = &jsonMocks.DecoderServiceMock{}
	suite.mockEncoderService = &jsonMocks.EncoderServiceMock{}
	suite.mockResponseFactory = &httpMocks.ResponseFactoryMock{}
	suite.mockParameterConverter = &httpMocks.ParameterConverterMock{}
	suite.mockDtoFactory = &dtoMocks.FactoryMock{}
	suite.sut = http.NewInventoryControllerImpl(
		suite.mockInventoryService,
		suite.mockDecoderService,
		suite.mockEncoderService,
		suite.mockResponseFactory,
		suite.mockParameterConverter,
		suite.mockDtoFactory,
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
	// Setup fixture
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
	// Setup fixture
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
	// Setup fixture
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

func (suite *InventoryControllerTestSuite) TestRead_WhenParameterConverterFails_ShouldFail() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(entity.InvalidID, mockErr)
	suite.mockResponseFactory.On("CreateFromError", mockErr).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.Read(pathParamFixture, nil, nil)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestRead_WhenInventoryServiceFails_ShouldFail() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	mockID := entity.ID(101)
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(mockID, nil)
	suite.mockInventoryService.On("Read", mockID).
		Return(nil, mockErr)
	suite.mockResponseFactory.On("CreateFromError", mockErr).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.Read(pathParamFixture, nil, nil)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestRead_WhenEncoderServiceFails_ShouldFail() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	mockID := entity.ID(101)
	mockEntity := &entityMocks.InventoryItemMock{Data: "some.data"}
	mockView := &dto.InventoryItemView{Name: "some.name"}
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(mockID, nil)
	suite.mockInventoryService.On("Read", mockID).
		Return(mockEntity, nil)
	suite.mockDtoFactory.On("CreateInventoryItemViewFromEntity", mockEntity).
		Return(mockView)
	suite.mockEncoderService.On("FromInventoryItemView", mockView).
		Return(nil, mockErr)
	suite.mockResponseFactory.On("CreateFromError", mockErr).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.Read(pathParamFixture, nil, nil)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestRead_WhenEncoderServicePasses_ShouldReturnAsExpected() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockID := entity.ID(101)
	mockEntity := &entityMocks.InventoryItemMock{Data: "some.data"}
	mockView := &dto.InventoryItemView{Name: "some.name"}
	mockJson := []byte("some.json")
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(mockID, nil)
	suite.mockInventoryService.On("Read", mockID).
		Return(mockEntity, nil)
	suite.mockDtoFactory.On("CreateInventoryItemViewFromEntity", mockEntity).
		Return(mockView)
	suite.mockEncoderService.On("FromInventoryItemView", mockView).
		Return(mockJson, nil)
	suite.mockResponseFactory.On("Create", uint(200), mockJson).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.Read(pathParamFixture, nil, nil)

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
