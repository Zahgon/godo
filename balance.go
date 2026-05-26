package godo

import (
	"context"
	"time"
)

// BalanceService is an interface for interfacing with the Balance
// endpoints of the DigitalOcean API
// See: https://docs.digitalocean.com/reference/api/api-reference/#operation/balance_get
type BalanceService interface {
	Get(context.Context) (*Balance, *Response, error)
}

// BalanceServiceOp handles communication with the Balance related methods of
// the DigitalOcean API.
type BalanceServiceOp struct {
	client *Client
}

var _ BalanceService = &BalanceServiceOp{}

// Balance represents a DigitalOcean Balance
type Balance struct {
	MonthToDateBalance string    `json:"month_to_date_balance"`
	AccountBalance     string    `json:"account_balance"`
	MonthToDateUsage   string    `json:"month_to_date_usage"`
	GeneratedAt        time.Time `json:"generated_at"`
}

func (r Balance) String() string { _ = "STUB: not implemented"; return "" }

// Get DigitalOcean balance info
func (s *BalanceServiceOp) Get(ctx context.Context) (*Balance, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
