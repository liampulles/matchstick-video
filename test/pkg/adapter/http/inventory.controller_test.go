package http_test

import (
	"fmt"
	goHttp "net/http"
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
	mockInventoryService   *inventoryMocks.MockService
	mockDecoderService     *jsonMocks.MockDecoderService
	mockEncoderService     *jsonMocks.MockEncoderService
	mockResponseFactory    *httpMocks.MockResponseFactory
	mockParameterConverter *httpMocks.MockParameterConverter
	sut                    *http.InventoryControllerImpl
}

func TestInventoryControllerTestSuite(t *testing.T) {
	suite.Run(t, new(InventoryControllerTestSuite))
}

func (suite *InventoryControllerTestSuite) SetupTest() {
	suite.mockInventoryService = &inventoryMocks.MockService{}
	suite.mockDecoderService = &jsonMocks.MockDecoderService{}
	suite.mockEncoderService = &jsonMocks.MockEncoderService{}
	suite.mockResponseFactory = &httpMocks.MockResponseFactory{}
	suite.mockParameterConverter = &httpMocks.MockParameterConverter{}
	suite.sut = http.NewInventoryControllerImpl(
		suite.mockInventoryService,
		suite.mockDecoderService,
		suite.mockEncoderService,
		suite.mockResponseFactory,
		suite.mockParameterConverter,
	)
}

func (suite *InventoryControllerTestSuite) TestGetHandlers_ShouldReturnAllHandlers() {
	// Setup expectations
	expected := []http.HandlerPattern{
		http.HandlerPattern{
			Method:      goHttp.MethodPost,
			PathPattern: "/inventory",
		},
		http.HandlerPattern{
			Method:      goHttp.MethodGet,
			PathPattern: "/inventory/{id}",
		},
		http.HandlerPattern{
			Method:      goHttp.MethodPut,
			PathPattern: "/inventory/{id}",
		},
		http.HandlerPattern{
			Method:      goHttp.MethodDelete,
			PathPattern: "/inventory/{id}",
		},
		http.HandlerPattern{
			Method:      goHttp.MethodPut,
			PathPattern: "/inventory/{id}/checkout",
		},
		http.HandlerPattern{
			Method:      goHttp.MethodPut,
			PathPattern: "/inventory/{id}/checkin",
		},
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
	requestFixture := &http.Request{
		Body: bodyFixture,
	}

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
	actual := suite.sut.Create(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestCreate_WhenInventoryServiceFails_ShouldFail() {
	// Setup fixture
	bodyFixture := []byte("some.body")
	requestFixture := &http.Request{
		Body: bodyFixture,
	}

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
	actual := suite.sut.Create(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestCreate_WhenInventoryServicePasses_ShouldReturnEntityResponse() {
	// Setup fixture
	bodyFixture := []byte("some.body")
	requestFixture := &http.Request{
		Body: bodyFixture,
	}

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
	actual := suite.sut.Create(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestReadDetails_WhenParameterConverterFails_ShouldFail() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}
	requestFixture := &http.Request{
		PathParam: pathParamFixture,
	}

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
	actual := suite.sut.ReadDetails(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestReadDetails_WhenInventoryServiceFails_ShouldFail() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}
	requestFixture := &http.Request{
		PathParam: pathParamFixture,
	}

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
	suite.mockInventoryService.On("ReadDetails", mockID).
		Return(nil, mockErr)
	suite.mockResponseFactory.On("CreateFromError", mockErr).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.ReadDetails(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestReadDetails_WhenEncoderServiceFails_ShouldFail() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}
	requestFixture := &http.Request{
		PathParam: pathParamFixture,
	}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockErr := fmt.Errorf("mock.error")
	mockID := entity.ID(101)
	mockView := &inventory.ViewVO{Name: "some.name"}
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(mockID, nil)
	suite.mockInventoryService.On("ReadDetails", mockID).
		Return(mockView, nil)
	suite.mockEncoderService.On("FromInventoryItemView", mockView).
		Return(nil, mockErr)
	suite.mockResponseFactory.On("CreateFromError", mockErr).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.ReadDetails(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestReadDetails_WhenEncoderServicePasses_ShouldReturnAsExpected() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}
	requestFixture := &http.Request{
		PathParam: pathParamFixture,
	}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockID := entity.ID(101)
	mockView := &inventory.ViewVO{Name: "some.name"}
	mockJson := []byte("some.json")
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(mockID, nil)
	suite.mockInventoryService.On("ReadDetails", mockID).
		Return(mockView, nil)
	suite.mockEncoderService.On("FromInventoryItemView", mockView).
		Return(mockJson, nil)
	suite.mockResponseFactory.On("CreateJSON", uint(200), mockJson).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.ReadDetails(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestUpdate_WhenParameterConverterFails_ShouldFail() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}
	requestFixture := &http.Request{
		PathParam: pathParamFixture,
	}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockErr := fmt.Errorf("some.error")
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(entity.InvalidID, mockErr)
	suite.mockResponseFactory.On("CreateFromError", mockErr).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.Update(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestUpdate_WhenDecoderServiceFails_ShouldFail() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}
	bodyFixture := []byte("some.body")
	requestFixture := &http.Request{
		PathParam: pathParamFixture,
		Body:      bodyFixture,
	}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockErr := fmt.Errorf("some.error")
	mockID := entity.ID(101)
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(mockID, nil)
	suite.mockDecoderService.On("ToInventoryUpdateItemVo", bodyFixture).
		Return(nil, mockErr)
	suite.mockResponseFactory.On("CreateFromError", mockErr).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.Update(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestUpdate_WhenInventoryServiceFails_ShouldFail() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}
	bodyFixture := []byte("some.body")
	requestFixture := &http.Request{
		PathParam: pathParamFixture,
		Body:      bodyFixture,
	}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockErr := fmt.Errorf("some.error")
	mockID := entity.ID(101)
	mockVo := &inventory.UpdateItemVO{Name: "some.name"}
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(mockID, nil)
	suite.mockDecoderService.On("ToInventoryUpdateItemVo", bodyFixture).
		Return(mockVo, nil)
	suite.mockInventoryService.On("Update", mockID, mockVo).
		Return(mockErr)
	suite.mockResponseFactory.On("CreateFromError", mockErr).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.Update(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestUpdate_WhenInventoryServicePasses_ShouldReturnAsExpected() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}
	bodyFixture := []byte("some.body")
	requestFixture := &http.Request{
		PathParam: pathParamFixture,
		Body:      bodyFixture,
	}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockID := entity.ID(101)
	mockVo := &inventory.UpdateItemVO{Name: "some.name"}
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(mockID, nil)
	suite.mockDecoderService.On("ToInventoryUpdateItemVo", bodyFixture).
		Return(mockVo, nil)
	suite.mockInventoryService.On("Update", mockID, mockVo).
		Return(nil)
	suite.mockResponseFactory.On("CreateEmpty", uint(204)).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.Update(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestDelete_WhenParameterConverterFails_ShouldFail() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}
	requestFixture := &http.Request{
		PathParam: pathParamFixture,
	}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockErr := fmt.Errorf("some.error")
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(entity.InvalidID, mockErr)
	suite.mockResponseFactory.On("CreateFromError", mockErr).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.Delete(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestDelete_WhenInventoryServiceFails_ShouldFail() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}
	requestFixture := &http.Request{
		PathParam: pathParamFixture,
	}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockID := entity.ID(101)
	mockErr := fmt.Errorf("some.error")
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(mockID, nil)
	suite.mockInventoryService.On("Delete", mockID).
		Return(mockErr)
	suite.mockResponseFactory.On("CreateFromError", mockErr).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.Delete(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestDelete_WhenInventoryServicePasses_ShouldReturnAsExpected() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}
	requestFixture := &http.Request{
		PathParam: pathParamFixture,
	}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockID := entity.ID(101)
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(mockID, nil)
	suite.mockInventoryService.On("Delete", mockID).
		Return(nil)
	suite.mockResponseFactory.On("CreateEmpty", uint(204)).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.Delete(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestCheckout_WhenParameterConverterFails_ShouldFail() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}
	requestFixture := &http.Request{
		PathParam: pathParamFixture,
	}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockErr := fmt.Errorf("some.error")
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(entity.InvalidID, mockErr)
	suite.mockResponseFactory.On("CreateFromError", mockErr).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.Checkout(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestCheckout_WhenInventoryServiceFails_ShouldFail() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}
	requestFixture := &http.Request{
		PathParam: pathParamFixture,
	}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockID := entity.ID(101)
	mockErr := fmt.Errorf("some.error")
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(mockID, nil)
	suite.mockInventoryService.On("Checkout", mockID).
		Return(mockErr)
	suite.mockResponseFactory.On("CreateFromError", mockErr).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.Checkout(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestCheckout_WhenInventoryServicePasses_ShouldReturnAsExpected() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}
	requestFixture := &http.Request{
		PathParam: pathParamFixture,
	}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockID := entity.ID(101)
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(mockID, nil)
	suite.mockInventoryService.On("Checkout", mockID).
		Return(nil)
	suite.mockResponseFactory.On("CreateEmpty", uint(204)).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.Checkout(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestCheckIn_WhenParameterConverterFails_ShouldFail() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}
	requestFixture := &http.Request{
		PathParam: pathParamFixture,
	}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockErr := fmt.Errorf("some.error")
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(entity.InvalidID, mockErr)
	suite.mockResponseFactory.On("CreateFromError", mockErr).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.CheckIn(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestCheckIn_WhenInventoryServiceFails_ShouldFail() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}
	requestFixture := &http.Request{
		PathParam: pathParamFixture,
	}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockID := entity.ID(101)
	mockErr := fmt.Errorf("some.error")
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(mockID, nil)
	suite.mockInventoryService.On("CheckIn", mockID).
		Return(mockErr)
	suite.mockResponseFactory.On("CreateFromError", mockErr).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.CheckIn(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

func (suite *InventoryControllerTestSuite) TestCheckIn_WhenInventoryServicePasses_ShouldReturnAsExpected() {
	// Setup fixture
	pathParamFixture := map[string]string{"some": "param"}
	requestFixture := &http.Request{
		PathParam: pathParamFixture,
	}

	// Setup expectations
	expected := &http.Response{
		StatusCode: 101,
		Body:       []byte("some.response"),
	}

	// Setup mocks
	mockID := entity.ID(101)
	suite.mockParameterConverter.On("ToEntityID", pathParamFixture, "id").
		Return(mockID, nil)
	suite.mockInventoryService.On("CheckIn", mockID).
		Return(nil)
	suite.mockResponseFactory.On("CreateEmpty", uint(204)).
		Return(expected)

	// Exercise SUT
	actual := suite.sut.CheckIn(requestFixture)

	// Verify results
	suite.Equal(expected, actual)
}

// EqualKeys matches the keys of a map
func equalKeys(expected []http.HandlerPattern, actual map[http.HandlerPattern]http.Handler) error {
	if len(actual) != len(expected) {
		return fmt.Errorf("Maps are different lengths. Expected: %d, Actual: %d", len(expected), len(actual))
	}
	for _, k := range expected {
		if _, ok := actual[k]; !ok {
			return fmt.Errorf("Key mismatch. Expected: %v, but was not in Actual.", k)
		}
	}
	return nil
}
