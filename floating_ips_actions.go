package godo

import (
	"context"
)

// FloatingIPActionsService is an interface for interfacing with the
// floating IPs actions endpoints of the Digital Ocean API.
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Floating-IP-Actions
type FloatingIPActionsService interface {
	Assign(ctx context.Context, ip string, dropletID int) (*Action, *Response, error)
	Unassign(ctx context.Context, ip string) (*Action, *Response, error)
	Get(ctx context.Context, ip string, actionID int) (*Action, *Response, error)
	List(ctx context.Context, ip string, opt *ListOptions) ([]Action, *Response, error)
}

// FloatingIPActionsServiceOp handles communication with the floating IPs
// action related methods of the DigitalOcean API.
type FloatingIPActionsServiceOp struct {
	client *Client
}

// Assign a floating IP to a droplet.
func (s *FloatingIPActionsServiceOp) Assign(ctx context.Context, ip string, dropletID int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Unassign a floating IP from the droplet it is currently assigned to.
func (s *FloatingIPActionsServiceOp) Unassign(ctx context.Context, ip string) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get an action for a particular floating IP by id.
func (s *FloatingIPActionsServiceOp) Get(ctx context.Context, ip string, actionID int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// List the actions for a particular floating IP.
func (s *FloatingIPActionsServiceOp) List(ctx context.Context, ip string, opt *ListOptions) ([]Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *FloatingIPActionsServiceOp) doAction(ctx context.Context, ip string, request *ActionRequest) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *FloatingIPActionsServiceOp) get(ctx context.Context, path string) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *FloatingIPActionsServiceOp) list(ctx context.Context, path string) ([]Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func floatingIPActionPath(ip string) string { _ = "STUB: not implemented"; return "" }
