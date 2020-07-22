package json_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/liampulles/matchstick-video/pkg/adapter/http/json"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

type DecoderServiceImplTestSuite struct {
	suite.Suite
	sut *json.DecoderServiceImpl
}

func TestDecoderServiceImplTestSuite(t *testing.T) {
	suite.Run(t, new(DecoderServiceImplTestSuite))
}

func (suite *DecoderServiceImplTestSuite) SetupTest() {
	suite.sut = json.NewDecoderServiceImpl()
}

func (suite *DecoderServiceImplTestSuite) TestToInventoryCreateItemVo_WhenUnmarshalFails_ShouldFail() {
	// Setup fixture
	fixture := []byte("not.json")

	// Setup expectations
	expectedErr := "could not unmarshal to inventory create item vo: invalid character 'o' in literal null (expecting 'u')"

	// Exercise SUT
	actual, err := suite.sut.ToInventoryCreateItemVo(fixture)

	// Verify results
	suite.Nil(actual)
	suite.EqualError(err, expectedErr)
}

func (suite *DecoderServiceImplTestSuite) TestToInventoryCreateItemVo_WhenUnmarshalPasses_ShouldPass() {
	// Setup fixture
	fixture := []byte("{\"name\": \"some.name\", \"location\": \"some.location\"}")

	// Setup expectations
	expected := &inventory.CreateItemVO{
		Name:     "some.name",
		Location: "some.location",
	}

	// Exercise SUT
	actual, err := suite.sut.ToInventoryCreateItemVo(fixture)

	// Verify results
	suite.NoError(err)
	suite.Equal(expected, actual)
}
