package godo

import (
	"context"
)

// StorageActionsService is an interface for interfacing with the
// storage actions endpoints of the Digital Ocean API.
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Block-Storage-Actions
type StorageActionsService interface {
	Attach(ctx context.Context, volumeID string, dropletID int) (*Action, *Response, error)
	DetachByDropletID(ctx context.Context, volumeID string, dropletID int) (*Action, *Response, error)
	Get(ctx context.Context, volumeID string, actionID int) (*Action, *Response, error)
	List(ctx context.Context, volumeID string, opt *ListOptions) ([]Action, *Response, error)
	Resize(ctx context.Context, volumeID string, sizeGigabytes int, regionSlug string) (*Action, *Response, error)
}

// StorageActionsServiceOp handles communication with the storage volumes
// action related methods of the DigitalOcean API.
type StorageActionsServiceOp struct {
	client *Client
}

// StorageAttachment represents the attachment of a block storage
// volume to a specific Droplet under the device name.
type StorageAttachment struct {
	DropletID int `json:"droplet_id"`
}

// Attach a storage volume to a Droplet.
func (s *StorageActionsServiceOp) Attach(ctx context.Context, volumeID string, dropletID int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// DetachByDropletID a storage volume from a Droplet by Droplet ID.
func (s *StorageActionsServiceOp) DetachByDropletID(ctx context.Context, volumeID string, dropletID int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get an action for a particular storage volume by id.
func (s *StorageActionsServiceOp) Get(ctx context.Context, volumeID string, actionID int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// List the actions for a particular storage volume.
func (s *StorageActionsServiceOp) List(ctx context.Context, volumeID string, opt *ListOptions) ([]Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Resize a storage volume.
func (s *StorageActionsServiceOp) Resize(ctx context.Context, volumeID string, sizeGigabytes int, regionSlug string) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *StorageActionsServiceOp) doAction(ctx context.Context, volumeID string, request *ActionRequest) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *StorageActionsServiceOp) get(ctx context.Context, path string) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *StorageActionsServiceOp) list(ctx context.Context, path string) ([]Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func storageAllocationActionPath(volumeID string) string { _ = "STUB: not implemented"; return "" }
