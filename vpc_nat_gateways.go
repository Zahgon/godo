package godo

import (
	"context"
	"time"
)

const (
	vpcNatGatewaysBasePath = "/v2/vpc_nat_gateways"
)

// VPCNATGatewaysService defines an interface for managing VPC NAT Gateways through the DigitalOcean API
type VPCNATGatewaysService interface {
	Create(context.Context, *VPCNATGatewayRequest) (*VPCNATGateway, *Response, error)
	Get(context.Context, string) (*VPCNATGateway, *Response, error)
	List(context.Context, *VPCNATGatewaysListOptions) ([]*VPCNATGateway, *Response, error)
	Update(context.Context, string, *VPCNATGatewayRequest) (*VPCNATGateway, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// VPCNATGatewayRequest represents a DigitalOcean VPC NAT Gateway create/update request
type VPCNATGatewayRequest struct {
	Name               string        `json:"name"`
	Type               string        `json:"type"`
	Region             string        `json:"region"`
	Size               uint32        `json:"size"`
	VPCs               []*IngressVPC `json:"vpcs"`
	UDPTimeoutSeconds  uint32        `json:"udp_timeout_seconds,omitempty"`
	ICMPTimeoutSeconds uint32        `json:"icmp_timeout_seconds,omitempty"`
	TCPTimeoutSeconds  uint32        `json:"tcp_timeout_seconds,omitempty"`
	ProjectID          string        `json:"project_id,omitempty"`
}

// VPCNATGateway represents a DigitalOcean VPC NAT Gateway resource
type VPCNATGateway struct {
	ID                 string        `json:"id"`
	Name               string        `json:"name"`
	Type               string        `json:"type"`
	State              string        `json:"state"`
	Region             string        `json:"region"`
	Size               uint32        `json:"size"`
	VPCs               []*IngressVPC `json:"vpcs"`
	Egresses           *Egresses     `json:"egresses,omitempty"`
	UDPTimeoutSeconds  uint32        `json:"udp_timeout_seconds,omitempty"`
	ICMPTimeoutSeconds uint32        `json:"icmp_timeout_seconds,omitempty"`
	TCPTimeoutSeconds  uint32        `json:"tcp_timeout_seconds,omitempty"`
	CreatedAt          time.Time     `json:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at"`
	ProjectID          string        `json:"project_id,omitempty"`
}

// IngressVPC defines the ingress configs supported by a VPC NAT Gateway
type IngressVPC struct {
	VpcUUID        string `json:"vpc_uuid"`
	GatewayIP      string `json:"gateway_ip,omitempty"`
	DefaultGateway bool   `json:"default_gateway,omitempty"`
}

// Egresses define egress routes supported by a VPC NAT Gateway
type Egresses struct {
	PublicGateways []*PublicGateway `json:"public_gateways,omitempty"`
}

// PublicGateway defines the public egress supported by a VPC NAT Gateway
type PublicGateway struct {
	IPv4 string `json:"ipv4"`
}

// VPCNATGatewaysListOptions define custom options for listing VPC NAT Gateways
type VPCNATGatewaysListOptions struct {
	ListOptions
	State  []string `url:"state,omitempty"`
	Region []string `url:"region,omitempty"`
	Type   []string `url:"type,omitempty"`
	Name   []string `url:"name,omitempty"`
}

type vpcNatGatewayRoot struct {
	VPCNATGateway *VPCNATGateway `json:"vpc_nat_gateway"`
}

type vpcNatGatewaysRoot struct {
	VPCNATGateways []*VPCNATGateway `json:"vpc_nat_gateways"`
	Links          *Links           `json:"links"`
	Meta           *Meta            `json:"meta"`
}

// VPCNATGatewaysServiceOp handles communication with VPC NAT Gateway methods of the DigitalOcean API
type VPCNATGatewaysServiceOp struct {
	client *Client
}

var _ VPCNATGatewaysService = &VPCNATGatewaysServiceOp{}

// Create a new VPC NAT Gateway
func (n *VPCNATGatewaysServiceOp) Create(ctx context.Context, createReq *VPCNATGatewayRequest) (*VPCNATGateway, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get an existing VPC NAT Gateway
func (n *VPCNATGatewaysServiceOp) Get(ctx context.Context, id string) (*VPCNATGateway, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// List all active VPC NAT Gateways
func (n *VPCNATGatewaysServiceOp) List(ctx context.Context, opts *VPCNATGatewaysListOptions) ([]*VPCNATGateway, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Update an existing VPC NAT Gateway
func (n *VPCNATGatewaysServiceOp) Update(ctx context.Context, id string, updateReq *VPCNATGatewayRequest) (*VPCNATGateway, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete an existing VPC NAT Gateway
func (n *VPCNATGatewaysServiceOp) Delete(ctx context.Context, id string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
