package godo

import (
	"context"
	"time"
)

const billingHistoryBasePath = "v2/customers/my/billing_history"

// BillingHistoryService is an interface for interfacing with the BillingHistory
// endpoints of the DigitalOcean API
// See: https://docs.digitalocean.com/reference/api/api-reference/#operation/billingHistory_list
type BillingHistoryService interface {
	List(context.Context, *ListOptions) (*BillingHistory, *Response, error)
}

// BillingHistoryServiceOp handles communication with the BillingHistory related methods of
// the DigitalOcean API.
type BillingHistoryServiceOp struct {
	client *Client
}

var _ BillingHistoryService = &BillingHistoryServiceOp{}

// BillingHistory represents a DigitalOcean Billing History
type BillingHistory struct {
	BillingHistory []BillingHistoryEntry `json:"billing_history"`
	Links          *Links                `json:"links"`
	Meta           *Meta                 `json:"meta"`
}

// BillingHistoryEntry represents an entry in a customer's Billing History
type BillingHistoryEntry struct {
	Description string    `json:"description"`
	Amount      string    `json:"amount"`
	InvoiceID   *string   `json:"invoice_id"`
	InvoiceUUID *string   `json:"invoice_uuid"`
	Date        time.Time `json:"date"`
	Type        string    `json:"type"`
}

func (b BillingHistory) String() string { _ = "STUB: not implemented"; return "" }

// List the Billing History for a customer
func (s *BillingHistoryServiceOp) List(ctx context.Context, opt *ListOptions) (*BillingHistory, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
