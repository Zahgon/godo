package godo

import (
	"context"
	"time"
)

const resourceV6Type = "ReservedIPv6"
const reservedIPV6sBasePath = "v2/reserved_ipv6"

// ReservedIPV6sService is an interface for interfacing with the reserved IPV6s
// endpoints of the Digital Ocean API.
type ReservedIPV6sService interface {
	List(context.Context, *ListOptions) ([]ReservedIPV6, *Response, error)
	Get(context.Context, string) (*ReservedIPV6, *Response, error)
	Create(context.Context, *ReservedIPV6CreateRequest) (*ReservedIPV6, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// ReservedIPV6sServiceOp handles communication with the reserved IPs related methods of the
// DigitalOcean API.
type ReservedIPV6sServiceOp struct {
	client *Client
}

var _ ReservedIPV6sService = (*ReservedIPV6sServiceOp)(nil)

// ReservedIPV6 represents a Digital Ocean reserved IP.
type ReservedIPV6 struct {
	RegionSlug string    `json:"region_slug"`
	IP         string    `json:"ip"`
	ReservedAt time.Time `json:"reserved_at"`
	Droplet    *Droplet  `json:"droplet,omitempty"`
}
type reservedIPV6Root struct {
	ReservedIPV6 *ReservedIPV6 `json:"reserved_ipv6"`
}

type reservedIPV6sRoot struct {
	ReservedIPV6s []ReservedIPV6 `json:"reserved_ipv6s"`
	Links         *Links         `json:"links"`
	Meta          *Meta          `json:"meta"`
}

func (f ReservedIPV6) String() string { _ = "STUB: not implemented"; return "" }

// URN returns the reserved IP in a valid DO API URN form.
func (f ReservedIPV6) URN() string { _ = "STUB: not implemented"; return "" }

// ReservedIPV6CreateRequest represents a request to reserve a reserved IP.
type ReservedIPV6CreateRequest struct {
	Region string `json:"region_slug,omitempty"`
}

// List all reserved IPV6s.
func (r *ReservedIPV6sServiceOp) List(ctx context.Context, opt *ListOptions) ([]ReservedIPV6, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get an individual reserved IPv6.
func (r *ReservedIPV6sServiceOp) Get(ctx context.Context, ip string) (*ReservedIPV6, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create a new IPv6
func (r *ReservedIPV6sServiceOp) Create(ctx context.Context, reserveRequest *ReservedIPV6CreateRequest) (*ReservedIPV6, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete a reserved IPv6.
func (r *ReservedIPV6sServiceOp) Delete(ctx context.Context, ip string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
