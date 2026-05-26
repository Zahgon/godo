package godo

import (
	"context"
)

// RegionsService is an interface for interfacing with the regions
// endpoints of the DigitalOcean API
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Regions
type RegionsService interface {
	List(context.Context, *ListOptions) ([]Region, *Response, error)
}

// RegionsServiceOp handles communication with the region related methods of the
// DigitalOcean API.
type RegionsServiceOp struct {
	client *Client
}

var _ RegionsService = &RegionsServiceOp{}

// Region represents a DigitalOcean Region
type Region struct {
	Slug      string   `json:"slug,omitempty"`
	Name      string   `json:"name,omitempty"`
	Sizes     []string `json:"sizes,omitempty"`
	Available bool     `json:"available,omitempty"`
	Features  []string `json:"features,omitempty"`
}

type regionsRoot struct {
	Regions []Region
	Links   *Links `json:"links"`
	Meta    *Meta  `json:"meta"`
}

func (r Region) String() string { _ = "STUB: not implemented"; return "" }

// List all regions
func (s *RegionsServiceOp) List(ctx context.Context, opt *ListOptions) ([]Region, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
