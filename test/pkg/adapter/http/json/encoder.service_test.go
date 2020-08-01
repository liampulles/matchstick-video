package json_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/liampulles/matchstick-video/pkg/adapter/http/json"
	"github.com/liampulles/matchstick-video/pkg/usecase/inventory"
)

type EncoderServiceImplTestSuite struct {
	suite.Suite
	sut *json.EncoderServiceImpl
}

func TestEncoderServiceImplTestSuite(t *testing.T) {
	suite.Run(t, new(EncoderServiceImplTestSuite))
}

func (suite *EncoderServiceImplTestSuite) SetupTest() {
	suite.sut = json.NewEncoderServiceImpl()
}

func (suite *EncoderServiceImplTestSuite) TestFromInventoryItemView_WhenMarshalPasses_ShouldPass() {
	// Setup fixture
	fixture := &inventory.ViewVO{
		ID:        101,
		Name:      "some.name",
		Location:  "some.location",
		Available: true,
	}

	// Setup expectations
	expected := "{\"id\":101,\"name\":\"some.name\",\"location\":\"some.location\",\"available\":true}"

	// Exercise SUT
	actual, err := suite.sut.FromInventoryItemView(fixture)

	// Verify results
	suite.NoError(err)
	suite.Equal(expected, string(actual))
}

func (suite *EncoderServiceImplTestSuite) TestFromInventoryItemThinViews_WhenMarshalPasses_ShouldPass() {
	// Setup fixture
	fixture := []inventory.ThinViewVO{
		inventory.ThinViewVO{
			ID:   101,
			Name: "some.name.1",
		},
		inventory.ThinViewVO{
			ID:   102,
			Name: "some.name.2",
		},
	}

	// Setup expectations
	expected := "[{\"id\":101,\"name\":\"some.name.1\"},{\"id\":102,\"name\":\"some.name.2\"}]"

	// Exercise SUT
	actual, err := suite.sut.FromInventoryItemThinViews(fixture)

	// Verify results
	suite.NoError(err)
	suite.Equal(expected, string(actual))
}

func (suite *EncoderServiceImplTestSuite) TestFromInventoryItemThinViews_GivenNilInput_WhenMarshalPasses_ShouldReturnEmptyJsonArray() {
	// Setup expectations
	expected := "[]"

	// Exercise SUT
	actual, err := suite.sut.FromInventoryItemThinViews(nil)

	// Verify results
	suite.NoError(err)
	suite.Equal(expected, string(actual))
}
