package json

import "github.com/liampulles/matchstick-video/pkg/usecase/inventory"

// DecoderService converts JSON to structs
type DecoderService interface {
	ToInventoryCreateItemVo(json []byte) (*inventory.CreateItemVO, error)
}
