package godo

import (
	"context"
	"time"
)

const byoipsBasePath = "/v2/byoip_prefixes"

// BYOIPsService is an interface for interacting with the BYOIPs
// endpoints of the Digital Ocean API.

type BYOIPPrefixesService interface {
	Create(context.Context, *BYOIPPrefixCreateReq) (*BYOIPPrefixCreateResp, *Response, error)
	List(context.Context, *ListOptions) ([]*BYOIPPrefix, *Response, error)
	Get(context.Context, string) (*BYOIPPrefix, *Response, error)
	GetResources(context.Context, string, *ListOptions) ([]BYOIPPrefixResource, *Response, error)
	Delete(context.Context, string) (*Response, error)
	Update(context.Context, string, *BYOIPPrefixUpdateReq) (*BYOIPPrefix, *Response, error)
}

// BYOIPPrefixServiceOp handles communication with the BYOIP Prefix related methods of the
// DigitalOcean API.
type BYOIPPrefixServiceOp struct {
	client *Client
}

var _ BYOIPPrefixesService = (*BYOIPPrefixServiceOp)(nil)

type BYOIPPrefix struct {
	Prefix        string `json:"prefix"`
	Status        string `json:"status"`
	UUID          string `json:"uuid"`
	Region        string `json:"region"`
	Validations   []any  `json:"validations"`
	FailureReason string `json:"failure_reason"`
	ProjectID     string `json:"project_id"`
	Advertised    bool   `json:"advertised"`
	Locked        bool   `json:"locked"`
}

// BYOIPPrefixCreateReq represents a request to create a BYOIP prefix.
type BYOIPPrefixCreateReq struct {
	Prefix    string `json:"prefix"`
	Signature string `json:"signature"`
	Region    string `json:"region"`
}

// BYOIPPrefixCreateResp represents the response from creating a BYOIP prefix.
type BYOIPPrefixCreateResp struct {
	UUID   string `json:"uuid"`
	Region string `json:"region"`
	Status string `json:"status"`
}

type byoipPrefixCreateRoot struct {
	BYOIPPrefixCreate *BYOIPPrefixCreateResp `json:"byoip_prefix"`
}

// BYOIPPrefixResource represents a BYOIP resource allocations
type BYOIPPrefixResource struct {
	ID         uint64    `json:"id"`
	BYOIP      string    `json:"byoip"`
	Resource   string    `json:"resource"`
	Region     string    `json:"region"`
	AssignedAt time.Time `json:"assigned_at"`
}

type byoipPrefixRoot struct {
	BYOIPPrefix *BYOIPPrefix `json:"byoip_prefix"`
}

type byoipsRoot struct {
	BYOIPs []*BYOIPPrefix `json:"byoip_prefixes"`
	Links  *Links         `json:"links"`
	Meta   *Meta          `json:"meta"`
}

type byoipResourcesRoot struct {
	Resources []BYOIPPrefixResource `json:"ips"`
	Links     *Links                `json:"links"`
	Meta      *Meta                 `json:"meta"`
}

type BYOIPPrefixUpdateReq struct {
	Advertise *bool `json:"advertise"`
}

type byoipPrefixUpdateRoot struct {
	BYOIPPrefix *BYOIPPrefix `json:"byoip_prefix"`
}

func (r BYOIPPrefix) String() string { _ = "STUB: not implemented"; return "" }

// List all BYOIP prefixes.
func (r *BYOIPPrefixServiceOp) List(ctx context.Context, opt *ListOptions) ([]*BYOIPPrefix, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get an individual BYOIP prefix details.
func (r *BYOIPPrefixServiceOp) Get(ctx context.Context, uuid string) (*BYOIPPrefix, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetResources return all existing BYOIP allocations for given BYOIP prefix id.
func (r *BYOIPPrefixServiceOp) GetResources(ctx context.Context, uuid string, opt *ListOptions) ([]BYOIPPrefixResource, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create a BYOIP prefix
func (r *BYOIPPrefixServiceOp) Create(ctx context.Context, byoipPrefix *BYOIPPrefixCreateReq) (*BYOIPPrefixCreateResp, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (r *BYOIPPrefixServiceOp) Delete(ctx context.Context, uuid string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update a BYOIP prefix
func (r *BYOIPPrefixServiceOp) Update(ctx context.Context, prefixUUID string, updateReq *BYOIPPrefixUpdateReq) (*BYOIPPrefix, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
