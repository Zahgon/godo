package godo

import (
	"context"
	"time"
)

const vpcsBasePath = "/v2/vpcs"

// VPCsService is an interface for managing Virtual Private Cloud configurations with the
// DigitalOcean API.
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/VPCs
type VPCsService interface {
	Create(context.Context, *VPCCreateRequest) (*VPC, *Response, error)
	Get(context.Context, string) (*VPC, *Response, error)
	List(context.Context, *ListOptions) ([]*VPC, *Response, error)
	ListMembers(context.Context, string, *VPCListMembersRequest, *ListOptions) ([]*VPCMember, *Response, error)
	Update(context.Context, string, *VPCUpdateRequest) (*VPC, *Response, error)
	Set(context.Context, string, ...VPCSetField) (*VPC, *Response, error)
	Delete(context.Context, string) (*Response, error)
	CreateVPCPeering(context.Context, *VPCPeeringCreateRequest) (*VPCPeering, *Response, error)
	GetVPCPeering(context.Context, string) (*VPCPeering, *Response, error)
	ListVPCPeerings(context.Context, *ListOptions) ([]*VPCPeering, *Response, error)
	UpdateVPCPeering(context.Context, string, *VPCPeeringUpdateRequest) (*VPCPeering, *Response, error)
	DeleteVPCPeering(context.Context, string) (*Response, error)
	CreateVPCPeeringByVPCID(context.Context, string, *VPCPeeringCreateRequestByVPCID) (*VPCPeering, *Response, error)
	ListVPCPeeringsByVPCID(context.Context, string, *ListOptions) ([]*VPCPeering, *Response, error)
	UpdateVPCPeeringByVPCID(context.Context, string, string, *VPCPeeringUpdateRequest) (*VPCPeering, *Response, error)
}

var _ VPCsService = &VPCsServiceOp{}

// VPCsServiceOp interfaces with VPC endpoints in the DigitalOcean API.
type VPCsServiceOp struct {
	client *Client
}

// VPCCreateRequest represents a request to create a Virtual Private Cloud.
type VPCCreateRequest struct {
	Name        string `json:"name,omitempty"`
	RegionSlug  string `json:"region,omitempty"`
	Description string `json:"description,omitempty"`
	IPRange     string `json:"ip_range,omitempty"`
}

// VPCUpdateRequest represents a request to update a Virtual Private Cloud.
type VPCUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Default     *bool  `json:"default,omitempty"`
}

// VPCSetField allows one to set individual fields within a VPC configuration.
type VPCSetField interface {
	vpcSetField(map[string]interface{})
}

// VPCSetName is used when one want to set the `name` field of a VPC.
// Ex.: VPCs.Set(..., VPCSetName("new-name"))
type VPCSetName string

// VPCSetDescription is used when one want to set the `description` field of a VPC.
// Ex.: VPCs.Set(..., VPCSetDescription("vpc description"))
type VPCSetDescription string

// VPCSetDefault is used when one wants to enable the `default` field of a VPC, to
// set a VPC as the default one in the region
// Ex.: VPCs.Set(..., VPCSetDefault())
func VPCSetDefault() VPCSetField {
	_ = "STUB: not implemented"
	return *

	// vpcSetDefault satisfies the VPCSetField interface
	new(VPCSetField)
}

type vpcSetDefault struct{}

// VPC represents a DigitalOcean Virtual Private Cloud configuration.
type VPC struct {
	ID          string    `json:"id,omitempty"`
	URN         string    `json:"urn"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	IPRange     string    `json:"ip_range,omitempty"`
	RegionSlug  string    `json:"region,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Default     bool      `json:"default,omitempty"`
}

type VPCListMembersRequest struct {
	ResourceType string `url:"resource_type,omitempty"`
}

type VPCMember struct {
	URN       string    `json:"urn,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type vpcRoot struct {
	VPC *VPC `json:"vpc"`
}

type vpcsRoot struct {
	VPCs  []*VPC `json:"vpcs"`
	Links *Links `json:"links"`
	Meta  *Meta  `json:"meta"`
}

type vpcMembersRoot struct {
	Members []*VPCMember `json:"members"`
	Links   *Links       `json:"links"`
	Meta    *Meta        `json:"meta"`
}

// Get returns the details of a Virtual Private Cloud.
func (v *VPCsServiceOp) Get(ctx context.Context, id string) (*VPC, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create creates a new Virtual Private Cloud.
func (v *VPCsServiceOp) Create(ctx context.Context, create *VPCCreateRequest) (*VPC, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// List returns a list of the caller's VPCs, with optional pagination.
func (v *VPCsServiceOp) List(ctx context.Context, opt *ListOptions) ([]*VPC, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Update updates a Virtual Private Cloud's properties.
func (v *VPCsServiceOp) Update(ctx context.Context, id string, update *VPCUpdateRequest) (*VPC, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (n VPCSetName) vpcSetField(in map[string]interface{}) { _ = "STUB: not implemented"; return }

func (n VPCSetDescription) vpcSetField(in map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func (*vpcSetDefault) vpcSetField(in map[string]interface{}) { _ = "STUB: not implemented"; return }

// Set updates specific properties of a Virtual Private Cloud.
func (v *VPCsServiceOp) Set(ctx context.Context, id string, fields ...VPCSetField) (*VPC, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete deletes a Virtual Private Cloud. There is no way to recover a VPC once it has been
// destroyed.
func (v *VPCsServiceOp) Delete(ctx context.Context, id string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *VPCsServiceOp) ListMembers(ctx context.Context, id string, request *VPCListMembersRequest, opt *ListOptions) ([]*VPCMember, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
