package godo

import (
	"context"
	"time"
)

const (
	dropletAutoscaleBasePath = "/v2/droplets/autoscale"
)

// DropletAutoscaleService defines an interface for managing droplet autoscale pools through DigitalOcean API
type DropletAutoscaleService interface {
	Create(context.Context, *DropletAutoscalePoolRequest) (*DropletAutoscalePool, *Response, error)
	Get(context.Context, string) (*DropletAutoscalePool, *Response, error)
	List(context.Context, *ListOptions) ([]*DropletAutoscalePool, *Response, error)
	ListMembers(context.Context, string, *ListOptions) ([]*DropletAutoscaleResource, *Response, error)
	ListHistory(context.Context, string, *ListOptions) ([]*DropletAutoscaleHistoryEvent, *Response, error)
	Update(context.Context, string, *DropletAutoscalePoolRequest) (*DropletAutoscalePool, *Response, error)
	Delete(context.Context, string) (*Response, error)
	DeleteDangerous(context.Context, string) (*Response, error)
}

// DropletAutoscalePool represents a DigitalOcean droplet autoscale pool
type DropletAutoscalePool struct {
	ID                 string                               `json:"id"`
	Name               string                               `json:"name"`
	Config             *DropletAutoscaleConfiguration       `json:"config"`
	DropletTemplate    *DropletAutoscaleResourceTemplate    `json:"droplet_template"`
	CreatedAt          time.Time                            `json:"created_at"`
	UpdatedAt          time.Time                            `json:"updated_at"`
	CurrentUtilization *DropletAutoscaleResourceUtilization `json:"current_utilization,omitempty"`
	Status             string                               `json:"status"`
}

// DropletAutoscaleConfiguration represents a DigitalOcean droplet autoscale pool configuration
type DropletAutoscaleConfiguration struct {
	MinInstances            uint64  `json:"min_instances,omitempty"`
	MaxInstances            uint64  `json:"max_instances,omitempty"`
	TargetCPUUtilization    float64 `json:"target_cpu_utilization,omitempty"`
	TargetMemoryUtilization float64 `json:"target_memory_utilization,omitempty"`
	CooldownMinutes         uint32  `json:"cooldown_minutes,omitempty"`
	TargetNumberInstances   uint64  `json:"target_number_instances,omitempty"`
}

// DropletAutoscaleResourceTemplate represents a DigitalOcean droplet autoscale pool resource template
type DropletAutoscaleResourceTemplate struct {
	Size             string   `json:"size"`
	Region           string   `json:"region"`
	Image            string   `json:"image"`
	Tags             []string `json:"tags"`
	SSHKeys          []string `json:"ssh_keys"`
	VpcUUID          string   `json:"vpc_uuid"`
	WithDropletAgent bool     `json:"with_droplet_agent"`
	ProjectID        string   `json:"project_id"`
	IPV6             bool     `json:"ipv6"`
	UserData         string   `json:"user_data"`
	PublicNetworking *bool    `json:"public_networking"`
}

// DropletAutoscaleResourceUtilization represents a DigitalOcean droplet autoscale pool resource utilization
type DropletAutoscaleResourceUtilization struct {
	Memory float64 `json:"memory,omitempty"`
	CPU    float64 `json:"cpu,omitempty"`
}

// DropletAutoscaleResource represents a DigitalOcean droplet autoscale pool resource
type DropletAutoscaleResource struct {
	DropletID          uint64                               `json:"droplet_id"`
	CreatedAt          time.Time                            `json:"created_at"`
	UpdatedAt          time.Time                            `json:"updated_at"`
	HealthStatus       string                               `json:"health_status"`
	UnhealthyReason    string                               `json:"unhealthy_reason,omitempty"`
	Status             string                               `json:"status"`
	CurrentUtilization *DropletAutoscaleResourceUtilization `json:"current_utilization,omitempty"`
}

// DropletAutoscaleHistoryEvent represents a DigitalOcean droplet autoscale pool history event
type DropletAutoscaleHistoryEvent struct {
	HistoryEventID       string    `json:"history_event_id"`
	CurrentInstanceCount uint64    `json:"current_instance_count"`
	DesiredInstanceCount uint64    `json:"desired_instance_count"`
	Reason               string    `json:"reason"`
	Status               string    `json:"status"`
	ErrorReason          string    `json:"error_reason,omitempty"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

// DropletAutoscalePoolRequest represents a DigitalOcean droplet autoscale pool create/update request
type DropletAutoscalePoolRequest struct {
	Name            string                            `json:"name"`
	Config          *DropletAutoscaleConfiguration    `json:"config"`
	DropletTemplate *DropletAutoscaleResourceTemplate `json:"droplet_template"`
}

type dropletAutoscalePoolRoot struct {
	AutoscalePool *DropletAutoscalePool `json:"autoscale_pool"`
}

type dropletAutoscalePoolsRoot struct {
	AutoscalePools []*DropletAutoscalePool `json:"autoscale_pools"`
	Links          *Links                  `json:"links"`
	Meta           *Meta                   `json:"meta"`
}

type dropletAutoscaleMembersRoot struct {
	Droplets []*DropletAutoscaleResource `json:"droplets"`
	Links    *Links                      `json:"links"`
	Meta     *Meta                       `json:"meta"`
}

type dropletAutoscaleHistoryEventsRoot struct {
	History []*DropletAutoscaleHistoryEvent `json:"history"`
	Links   *Links                          `json:"links"`
	Meta    *Meta                           `json:"meta"`
}

// DropletAutoscaleServiceOp handles communication with droplet autoscale-related methods of the DigitalOcean API
type DropletAutoscaleServiceOp struct {
	client *Client
}

var _ DropletAutoscaleService = &DropletAutoscaleServiceOp{}

// Create a new droplet autoscale pool
func (d *DropletAutoscaleServiceOp) Create(ctx context.Context, createReq *DropletAutoscalePoolRequest) (*DropletAutoscalePool, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get an existing droplet autoscale pool
func (d *DropletAutoscaleServiceOp) Get(ctx context.Context, id string) (*DropletAutoscalePool, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// List all existing droplet autoscale pools
func (d *DropletAutoscaleServiceOp) List(ctx context.Context, opts *ListOptions) ([]*DropletAutoscalePool, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ListMembers all members for an existing droplet autoscale pool
func (d *DropletAutoscaleServiceOp) ListMembers(ctx context.Context, id string, opts *ListOptions) ([]*DropletAutoscaleResource, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ListHistory all history events for an existing droplet autoscale pool
func (d *DropletAutoscaleServiceOp) ListHistory(ctx context.Context, id string, opts *ListOptions) ([]*DropletAutoscaleHistoryEvent, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Update an existing autoscale pool
func (d *DropletAutoscaleServiceOp) Update(ctx context.Context, id string, updateReq *DropletAutoscalePoolRequest) (*DropletAutoscalePool, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete an existing autoscale pool
func (d *DropletAutoscaleServiceOp) Delete(ctx context.Context, id string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteDangerous deletes an existing autoscale pool with all underlying resources
func (d *DropletAutoscaleServiceOp) DeleteDangerous(ctx context.Context, id string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
