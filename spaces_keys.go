package godo

import (
	"context"
)

const spacesKeysBasePath = "v2/spaces/keys"

// SpacesKeysService is an interface for managing Spaces keys with the DigitalOcean API.
type SpacesKeysService interface {
	List(context.Context, *ListOptions) ([]*SpacesKey, *Response, error)
	Update(context.Context, string, *SpacesKeyUpdateRequest) (*SpacesKey, *Response, error)
	Create(context.Context, *SpacesKeyCreateRequest) (*SpacesKey, *Response, error)
	Delete(context.Context, string) (*Response, error)
	Get(context.Context, string) (*SpacesKey, *Response, error)
}

// SpacesKeysServiceOp handles communication with the Spaces key related methods of the
// DigitalOcean API.
type SpacesKeysServiceOp struct {
	client *Client
}

var _ SpacesKeysService = &SpacesKeysServiceOp{}

// SpacesKeyPermission represents a permission for a Spaces grant
type SpacesKeyPermission string

const (
	// SpacesKeyRead grants read-only access to the Spaces bucket
	SpacesKeyRead SpacesKeyPermission = "read"
	// SpacesKeyReadWrite grants read and write access to the Spaces bucket
	SpacesKeyReadWrite SpacesKeyPermission = "readwrite"
	// SpacesKeyFullAccess grants full access to the Spaces bucket
	SpacesKeyFullAccess SpacesKeyPermission = "fullaccess"
)

// Grant represents a Grant for a Spaces key
type Grant struct {
	Bucket     string              `json:"bucket"`
	Permission SpacesKeyPermission `json:"permission"`
}

// SpacesKey represents a DigitalOcean Spaces key
type SpacesKey struct {
	Name      string   `json:"name"`
	AccessKey string   `json:"access_key"`
	SecretKey string   `json:"secret_key"`
	Grants    []*Grant `json:"grants"`
	CreatedAt string   `json:"created_at"`
}

// SpacesKeyRoot represents a response from the DigitalOcean API
type spacesKeyRoot struct {
	Key *SpacesKey `json:"key"`
}

// SpacesKeyCreateRequest represents a request to create a Spaces key.
type SpacesKeyCreateRequest struct {
	Name   string   `json:"name"`
	Grants []*Grant `json:"grants"`
}

// SpacesKeyUpdateRequest represents a request to update a Spaces key.
type SpacesKeyUpdateRequest struct {
	Name   string   `json:"name"`
	Grants []*Grant `json:"grants"`
}

// spacesListKeysRoot represents a response from the DigitalOcean API
type spacesListKeysRoot struct {
	Keys  []*SpacesKey `json:"keys,omitempty"`
	Links *Links       `json:"links,omitempty"`
	Meta  *Meta        `json:"meta"`
}

// Create creates a new Spaces key.
func (s *SpacesKeysServiceOp) Create(ctx context.Context, createRequest *SpacesKeyCreateRequest) (*SpacesKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete deletes a Spaces key.
func (s *SpacesKeysServiceOp) Delete(ctx context.Context, accessKey string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update updates a Spaces key.
func (s *SpacesKeysServiceOp) Update(ctx context.Context, accessKey string, updateRequest *SpacesKeyUpdateRequest) (*SpacesKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// List returns a list of Spaces keys.
func (s *SpacesKeysServiceOp) List(ctx context.Context, opts *ListOptions) ([]*SpacesKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get retrieves a Spaces key.
func (s *SpacesKeysServiceOp) Get(ctx context.Context, accessKey string) (*SpacesKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
