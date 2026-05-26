package godo

import (
	"context"
)

// ReservedIPActionsService is an interface for interfacing with the
// reserved IPs actions endpoints of the Digital Ocean API.
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Reserved-IP-Actions
type ReservedIPV6ActionsService interface {
	Assign(ctx context.Context, ip string, dropletID int) (*Action, *Response, error)
	Unassign(ctx context.Context, ip string) (*Action, *Response, error)
}

// ReservedIPActionsServiceOp handles communication with the reserved IPs
// action related methods of the DigitalOcean API.
type ReservedIPV6ActionsServiceOp struct {
	client *Client
}

// Assign a reserved IP to a droplet.
func (s *ReservedIPV6ActionsServiceOp) Assign(ctx context.Context, ip string, dropletID int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Unassign a rerserved IP from the droplet it is currently assigned to.
func (s *ReservedIPV6ActionsServiceOp) Unassign(ctx context.Context, ip string) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *ReservedIPV6ActionsServiceOp) doV6Action(ctx context.Context, ip string, request *ActionRequest) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func reservedIPV6ActionPath(ip string) string { _ = "STUB: not implemented"; return "" }
