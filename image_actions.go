package godo

import (
	"context"
)

// ImageActionsService is an interface for interfacing with the image actions
// endpoints of the DigitalOcean API
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Image-Actions
type ImageActionsService interface {
	Get(context.Context, int, int) (*Action, *Response, error)
	GetByURI(context.Context, string) (*Action, *Response, error)
	Transfer(context.Context, int, *ActionRequest) (*Action, *Response, error)
	Convert(context.Context, int) (*Action, *Response, error)
}

// ImageActionsServiceOp handles communication with the image action related methods of the
// DigitalOcean API.
type ImageActionsServiceOp struct {
	client *Client
}

var _ ImageActionsService = &ImageActionsServiceOp{}

// Transfer an image
func (i *ImageActionsServiceOp) Transfer(ctx context.Context, imageID int, transferRequest *ActionRequest) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Convert an image to a snapshot
func (i *ImageActionsServiceOp) Convert(ctx context.Context, imageID int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get an action for a particular image by id.
func (i *ImageActionsServiceOp) Get(ctx context.Context, imageID, actionID int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetByURI gets an action for a particular image by URI.
func (i *ImageActionsServiceOp) GetByURI(ctx context.Context, rawurl string) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (i *ImageActionsServiceOp) get(ctx context.Context, path string) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
