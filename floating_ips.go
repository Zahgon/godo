package godo

import (
	"context"
)

const floatingBasePath = "v2/floating_ips"

// FloatingIPsService is an interface for interfacing with the floating IPs
// endpoints of the Digital Ocean API.
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Floating-IPs
type FloatingIPsService interface {
	List(context.Context, *ListOptions) ([]FloatingIP, *Response, error)
	Get(context.Context, string) (*FloatingIP, *Response, error)
	Create(context.Context, *FloatingIPCreateRequest) (*FloatingIP, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// FloatingIPsServiceOp handles communication with the floating IPs related methods of the
// DigitalOcean API.
type FloatingIPsServiceOp struct {
	client *Client
}

var _ FloatingIPsService = &FloatingIPsServiceOp{}

// FloatingIP represents a Digital Ocean floating IP.
type FloatingIP struct {
	Region    *Region  `json:"region"`
	Droplet   *Droplet `json:"droplet"`
	IP        string   `json:"ip"`
	ProjectID string   `json:"project_id"`
	Locked    bool     `json:"locked"`
}

func (f FloatingIP) String() string { _ = "STUB: not implemented"; return "" }

// URN returns the floating IP in a valid DO API URN form.
func (f FloatingIP) URN() string { _ = "STUB: not implemented"; return "" }

type floatingIPsRoot struct {
	FloatingIPs []FloatingIP `json:"floating_ips"`
	Links       *Links       `json:"links"`
	Meta        *Meta        `json:"meta"`
}

type floatingIPRoot struct {
	FloatingIP *FloatingIP `json:"floating_ip"`
	Links      *Links      `json:"links,omitempty"`
}

// FloatingIPCreateRequest represents a request to create a floating IP.
// Specify DropletID to assign the floating IP to a Droplet or Region
// to reserve it to the region.
type FloatingIPCreateRequest struct {
	Region    string `json:"region,omitempty"`
	DropletID int    `json:"droplet_id,omitempty"`
	ProjectID string `json:"project_id,omitempty"`
}

// List all floating IPs.
func (f *FloatingIPsServiceOp) List(ctx context.Context, opt *ListOptions) ([]FloatingIP, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get an individual floating IP.
func (f *FloatingIPsServiceOp) Get(ctx context.Context, ip string) (*FloatingIP, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create a floating IP. If the DropletID field of the request is not empty,
// the floating IP will also be assigned to the droplet.
func (f *FloatingIPsServiceOp) Create(ctx context.Context, createRequest *FloatingIPCreateRequest) (*FloatingIP, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete a floating IP.
func (f *FloatingIPsServiceOp) Delete(ctx context.Context, ip string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
