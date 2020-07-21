package mux

import (
	"net/http"

	adapterHttp "github.com/liampulles/matchstick-video/pkg/adapter/http"
)

// IOMapper maps requests and responses between mux's format
// and the format needed by the adapter layer.
type IOMapper interface {
	MapRequest(*http.Request) (*adapterHttp.Request, error)
	MapResponse(*adapterHttp.Response, http.ResponseWriter)
}
