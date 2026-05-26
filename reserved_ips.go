package godo

import (
	"context"
)

const resourceType = "ReservedIP"
const reservedIPsBasePath = "v2/reserved_ips"

// ReservedIPsService is an interface for interfacing with the reserved IPs
// endpoints of the Digital Ocean API.
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Reserved-IPs
type ReservedIPsService interface {
	List(context.Context, *ListOptions) ([]ReservedIP, *Response, error)
	Get(context.Context, string) (*ReservedIP, *Response, error)
	Create(context.Context, *ReservedIPCreateRequest) (*ReservedIP, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// ReservedIPsServiceOp handles communication with the reserved IPs related methods of the
// DigitalOcean API.
type ReservedIPsServiceOp struct {
	client *Client
}

var _ ReservedIPsService = &ReservedIPsServiceOp{}

// ReservedIP represents a Digital Ocean reserved IP.
type ReservedIP struct {
	Region    *Region  `json:"region"`
	Droplet   *Droplet `json:"droplet"`
	IP        string   `json:"ip"`
	ProjectID string   `json:"project_id"`
	Locked    bool     `json:"locked"`
}

func (f ReservedIP) String() string { _ = "STUB: not implemented"; return "" }

// URN returns the reserved IP in a valid DO API URN form.
func (f ReservedIP) URN() string { _ = "STUB: not implemented"; return "" }

type reservedIPsRoot struct {
	ReservedIPs []ReservedIP `json:"reserved_ips"`
	Links       *Links       `json:"links"`
	Meta        *Meta        `json:"meta"`
}

type reservedIPRoot struct {
	ReservedIP *ReservedIP `json:"reserved_ip"`
	Links      *Links      `json:"links,omitempty"`
}

// ReservedIPCreateRequest represents a request to create a reserved IP.
// Specify DropletID to assign the reserved IP to a Droplet or Region
// to reserve it to the region.
type ReservedIPCreateRequest struct {
	Region    string `json:"region,omitempty"`
	DropletID int    `json:"droplet_id,omitempty"`
	ProjectID string `json:"project_id,omitempty"`
}

// List all reserved IPs.
func (r *ReservedIPsServiceOp) List(ctx context.Context, opt *ListOptions) ([]ReservedIP, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get an individual reserved IP.
func (r *ReservedIPsServiceOp) Get(ctx context.Context, ip string) (*ReservedIP, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create a reserved IP. If the DropletID field of the request is not empty,
// the reserved IP will also be assigned to the droplet.
func (r *ReservedIPsServiceOp) Create(ctx context.Context, createRequest *ReservedIPCreateRequest) (*ReservedIP, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete a reserved IP.
func (r *ReservedIPsServiceOp) Delete(ctx context.Context, ip string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
