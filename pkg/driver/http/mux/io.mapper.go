package mux

import (
	"fmt"
	"io/ioutil"
	"net/http"

	adapterHttp "github.com/liampulles/matchstick-video/pkg/adapter/http"
)

// IOMapper maps requests and responses between mux's format
// and the format needed by the adapter layer.
type IOMapper interface {
	MapRequest(*http.Request) (*adapterHttp.Request, error)
	MapResponse(*adapterHttp.Response, http.ResponseWriter)
}

// IOMapperImpl implements IOMapper
type IOMapperImpl struct {
	wrapper Wrapper
}

var _ IOMapper = &IOMapperImpl{}

// NewIOMapperImpl is a constructor
func NewIOMapperImpl(wrapper Wrapper) *IOMapperImpl {
	return &IOMapperImpl{
		wrapper: wrapper,
	}
}

// MapRequest converts mux's (i.e. Go's) notion of a request to the adapter
// version
func (i *IOMapperImpl) MapRequest(req *http.Request) (*adapterHttp.Request, error) {
	pathParam := i.extractPathParam(req)
	queryParam := extractQueryParam(req)
	body, err := extractBody(req)
	if err != nil {
		return nil, err
	}

	return &adapterHttp.Request{
		PathParam:  pathParam,
		QueryParam: queryParam,
		Body:       body,
	}, nil
}

// MapResponse converts the adapter's notion of a response to mux's (i.e. Go's)
// version
func (i *IOMapperImpl) MapResponse(adapterResp *adapterHttp.Response, goResp http.ResponseWriter) {
	goResp.Header().Set("Content-Type", adapterResp.ContentType)
	goResp.WriteHeader(int(adapterResp.StatusCode))
	goResp.Write(adapterResp.Body)
}

func (i *IOMapperImpl) extractPathParam(req *http.Request) map[string]string {
	return i.wrapper.Vars(req)
}

func extractQueryParam(req *http.Request) map[string][]string {
	return req.URL.Query()
}

func extractBody(req *http.Request) ([]byte, error) {
	bytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("could not extract body: %w", err)
	}
	return bytes, nil
}
