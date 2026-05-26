package godo

import (
	"context"
)

const firewallsBasePath = "/v2/firewalls"

// FirewallsService is an interface for managing Firewalls with the DigitalOcean API.
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Firewalls
type FirewallsService interface {
	Get(context.Context, string) (*Firewall, *Response, error)
	Create(context.Context, *FirewallRequest) (*Firewall, *Response, error)
	Update(context.Context, string, *FirewallRequest) (*Firewall, *Response, error)
	Delete(context.Context, string) (*Response, error)
	List(context.Context, *ListOptions) ([]Firewall, *Response, error)
	ListByDroplet(context.Context, int, *ListOptions) ([]Firewall, *Response, error)
	AddDroplets(context.Context, string, ...int) (*Response, error)
	RemoveDroplets(context.Context, string, ...int) (*Response, error)
	AddTags(context.Context, string, ...string) (*Response, error)
	RemoveTags(context.Context, string, ...string) (*Response, error)
	AddRules(context.Context, string, *FirewallRulesRequest) (*Response, error)
	RemoveRules(context.Context, string, *FirewallRulesRequest) (*Response, error)
}

// FirewallsServiceOp handles communication with Firewalls methods of the DigitalOcean API.
type FirewallsServiceOp struct {
	client *Client
}

// Firewall represents a DigitalOcean Firewall configuration.
type Firewall struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	Status         string          `json:"status"`
	InboundRules   []InboundRule   `json:"inbound_rules"`
	OutboundRules  []OutboundRule  `json:"outbound_rules"`
	DropletIDs     []int           `json:"droplet_ids"`
	Tags           []string        `json:"tags"`
	Created        string          `json:"created_at"`
	PendingChanges []PendingChange `json:"pending_changes"`
}

// String creates a human-readable description of a Firewall.
func (fw Firewall) String() string { _ = "STUB: not implemented"; return "" }

// URN returns the firewall name in a valid DO API URN form.
func (fw Firewall) URN() string { _ = "STUB: not implemented"; return "" }

// FirewallRequest represents the configuration to be applied to an existing or a new Firewall.
type FirewallRequest struct {
	Name          string         `json:"name"`
	InboundRules  []InboundRule  `json:"inbound_rules"`
	OutboundRules []OutboundRule `json:"outbound_rules"`
	DropletIDs    []int          `json:"droplet_ids"`
	Tags          []string       `json:"tags"`
}

// FirewallRulesRequest represents rules configuration to be applied to an existing Firewall.
type FirewallRulesRequest struct {
	InboundRules  []InboundRule  `json:"inbound_rules"`
	OutboundRules []OutboundRule `json:"outbound_rules"`
}

// InboundRule represents a DigitalOcean Firewall inbound rule.
type InboundRule struct {
	Protocol  string   `json:"protocol,omitempty"`
	PortRange string   `json:"ports,omitempty"`
	Sources   *Sources `json:"sources"`
}

// OutboundRule represents a DigitalOcean Firewall outbound rule.
type OutboundRule struct {
	Protocol     string        `json:"protocol,omitempty"`
	PortRange    string        `json:"ports,omitempty"`
	Destinations *Destinations `json:"destinations"`
}

// Sources represents a DigitalOcean Firewall InboundRule sources.
type Sources struct {
	Addresses        []string `json:"addresses,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	DropletIDs       []int    `json:"droplet_ids,omitempty"`
	LoadBalancerUIDs []string `json:"load_balancer_uids,omitempty"`
	KubernetesIDs    []string `json:"kubernetes_ids,omitempty"`
}

// PendingChange represents a DigitalOcean Firewall status details.
type PendingChange struct {
	DropletID int    `json:"droplet_id,omitempty"`
	Removing  bool   `json:"removing,omitempty"`
	Status    string `json:"status,omitempty"`
}

// Destinations represents a DigitalOcean Firewall OutboundRule destinations.
type Destinations struct {
	Addresses        []string `json:"addresses,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	DropletIDs       []int    `json:"droplet_ids,omitempty"`
	LoadBalancerUIDs []string `json:"load_balancer_uids,omitempty"`
	KubernetesIDs    []string `json:"kubernetes_ids,omitempty"`
}

var _ FirewallsService = &FirewallsServiceOp{}

// Get an existing Firewall by its identifier.
func (fw *FirewallsServiceOp) Get(ctx context.Context, fID string) (*Firewall, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create a new Firewall with a given configuration.
func (fw *FirewallsServiceOp) Create(ctx context.Context, fr *FirewallRequest) (*Firewall, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Update an existing Firewall with new configuration.
func (fw *FirewallsServiceOp) Update(ctx context.Context, fID string, fr *FirewallRequest) (*Firewall, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete a Firewall by its identifier.
func (fw *FirewallsServiceOp) Delete(ctx context.Context, fID string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List Firewalls.
func (fw *FirewallsServiceOp) List(ctx context.Context, opt *ListOptions) ([]Firewall, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ListByDroplet Firewalls.
func (fw *FirewallsServiceOp) ListByDroplet(ctx context.Context, dID int, opt *ListOptions) ([]Firewall, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// AddDroplets to a Firewall.
func (fw *FirewallsServiceOp) AddDroplets(ctx context.Context, fID string, dropletIDs ...int) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveDroplets from a Firewall.
func (fw *FirewallsServiceOp) RemoveDroplets(ctx context.Context, fID string, dropletIDs ...int) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddTags to a Firewall.
func (fw *FirewallsServiceOp) AddTags(ctx context.Context, fID string, tags ...string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveTags from a Firewall.
func (fw *FirewallsServiceOp) RemoveTags(ctx context.Context, fID string, tags ...string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddRules to a Firewall.
func (fw *FirewallsServiceOp) AddRules(ctx context.Context, fID string, rr *FirewallRulesRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveRules from a Firewall.
func (fw *FirewallsServiceOp) RemoveRules(ctx context.Context, fID string, rr *FirewallRulesRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type dropletsRequest struct {
	IDs []int `json:"droplet_ids"`
}

type tagsRequest struct {
	Tags []string `json:"tags"`
}

type firewallRoot struct {
	Firewall *Firewall `json:"firewall"`
}

type firewallsRoot struct {
	Firewalls []Firewall `json:"firewalls"`
	Links     *Links     `json:"links"`
	Meta      *Meta      `json:"meta"`
}

func (fw *FirewallsServiceOp) createAndDoReq(ctx context.Context, method, path string, v interface{}) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fw *FirewallsServiceOp) listHelper(ctx context.Context, path string) ([]Firewall, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
