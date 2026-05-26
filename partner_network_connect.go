package godo

import (
	"context"
	"time"
)

const partnerNetworkConnectBasePath = "/v2/partner_network_connect/attachments"

// PartnerAttachmentService is an interface for managing Partner Attachments with the
// DigitalOcean API.
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/PartnerNetworkConnect
type PartnerAttachmentService interface {
	List(context.Context, *ListOptions) ([]*PartnerAttachment, *Response, error)
	Create(context.Context, *PartnerAttachmentCreateRequest) (*PartnerAttachment, *Response, error)
	Get(context.Context, string) (*PartnerAttachment, *Response, error)
	Update(context.Context, string, *PartnerAttachmentUpdateRequest) (*PartnerAttachment, *Response, error)
	Delete(context.Context, string) (*Response, error)
	GetServiceKey(context.Context, string) (*ServiceKey, *Response, error)
	ListRoutes(context.Context, string, *ListOptions) ([]*RemoteRoute, *Response, error)
	GetBGPAuthKey(ctx context.Context, iaID string) (*BgpAuthKey, *Response, error)
	RegenerateServiceKey(ctx context.Context, iaID string) (*RegenerateServiceKey, *Response, error)
}

var _ PartnerAttachmentService = &PartnerAttachmentServiceOp{}

// PartnerAttachmentServiceOp interfaces with the Partner Attachment endpoints in the DigitalOcean API.
type PartnerAttachmentServiceOp struct {
	client *Client
}

// PartnerAttachmentCreateRequest represents a request to create a Partner Attachment.
type PartnerAttachmentCreateRequest struct {
	// Name is the name of the Partner Attachment
	Name string `json:"name,omitempty"`
	// ConnectionBandwidthInMbps is the bandwidth of the connection in Mbps
	ConnectionBandwidthInMbps int `json:"connection_bandwidth_in_mbps,omitempty"`
	// Region is the region where the Partner Attachment is created
	Region string `json:"region,omitempty"`
	// NaaSProvider is the name of the Network as a Service provider
	NaaSProvider string `json:"naas_provider,omitempty"`
	// VPCIDs is the IDs of the VPCs to which the Partner Attachment is connected to
	VPCIDs []string `json:"vpc_ids,omitempty"`
	// BGP is the BGP configuration of the Partner Attachment
	BGP BGP `json:"bgp,omitempty"`
	// RedundancyZone is the redundancy zone of the Partner Attachment
	RedundancyZone string `json:"redundancy_zone,omitempty"`
	// ParentUuid is the uuid of the parent Partner Attachment if there is one
	ParentUuid string `json:"parent_uuid,omitempty"`
}

type partnerAttachmentRequestBody struct {
	// Name is the name of the Partner Attachment
	Name string `json:"name,omitempty"`
	// ConnectionBandwidthInMbps is the bandwidth of the connection in Mbps
	ConnectionBandwidthInMbps int `json:"connection_bandwidth_in_mbps,omitempty"`
	// Region is the region where the Partner Attachment is created
	Region string `json:"region,omitempty"`
	// NaaSProvider is the name of the Network as a Service provider
	NaaSProvider string `json:"naas_provider,omitempty"`
	// VPCIDs is the IDs of the VPCs to which the Partner Attachment is connected to
	VPCIDs []string `json:"vpc_ids,omitempty"`
	// BGP is the BGP configuration of the Partner Attachment
	BGP *BGPInput `json:"bgp,omitempty"`
	// RedundancyZone is the redundancy zone of the Partner Attachment
	RedundancyZone string `json:"redundancy_zone,omitempty"`
	// ParentUuid is the uuid of the parent Partner Attachment if there is one
	ParentUuid string `json:"parent_uuid,omitempty"`
}

func (req *PartnerAttachmentCreateRequest) buildReq() *partnerAttachmentRequestBody {
	_ = "STUB: not implemented"
	return nil
}

// PartnerAttachmentUpdateRequest represents a request to update a Partner Attachment.
type PartnerAttachmentUpdateRequest struct {
	// Name is the name of the Partner Attachment
	Name string `json:"name,omitempty"`
	//VPCIDs is the IDs of the VPCs to which the Partner Attachment is connected to
	VPCIDs []string `json:"vpc_ids,omitempty"`
}

// BGP represents the BGP configuration of a Partner Attachment.
type BGP struct {
	// LocalASN is the local ASN
	LocalASN int `json:"local_asn,omitempty"`
	// LocalRouterIP is the local router IP
	LocalRouterIP string `json:"local_router_ip,omitempty"`
	// PeerASN is the peer ASN
	PeerASN int `json:"peer_asn,omitempty"`
	// PeerRouterIP is the peer router IP
	PeerRouterIP string `json:"peer_router_ip,omitempty"`
	// AuthKey is the authentication key
	AuthKey string `json:"auth_key,omitempty"`
}

func (b *BGP) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// BGPInput represents the BGP configuration of a Partner Attachment.
type BGPInput struct {
	// LocalASN is the local ASN
	LocalASN int `json:"local_router_asn,omitempty"`
	// LocalRouterIP is the local router IP
	LocalRouterIP string `json:"local_router_ip,omitempty"`
	// PeerASN is the peer ASN
	PeerASN int `json:"peer_router_asn,omitempty"`
	// PeerRouterIP is the peer router IP
	PeerRouterIP string `json:"peer_router_ip,omitempty"`
	// AuthKey is the authentication key
	AuthKey string `json:"auth_key,omitempty"`
}

// ServiceKey represents the service key of a Partner Attachment.
type ServiceKey struct {
	Value     string    `json:"value,omitempty"`
	State     string    `json:"state,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// RemoteRoute represents a route for a Partner Attachment.
type RemoteRoute struct {
	// Cidr is the CIDR of the route
	Cidr string `json:"cidr,omitempty"`
}

// PartnerAttachment represents a DigitalOcean Partner Attachment.
type PartnerAttachment struct {
	// ID is the generated ID of the Partner Attachment
	ID string `json:"id,omitempty"`
	// Name is the name of the Partner Attachment
	Name string `json:"name,omitempty"`
	// State is the state of the Partner Attachment
	State string `json:"state,omitempty"`
	// ConnectionBandwidthInMbps is the bandwidth of the connection in Mbps
	ConnectionBandwidthInMbps int `json:"connection_bandwidth_in_mbps,omitempty"`
	// Region is the region where the Partner Attachment is created
	Region string `json:"region,omitempty"`
	// NaaSProvider is the name of the Network as a Service provider
	NaaSProvider string `json:"naas_provider,omitempty"`
	// VPCIDs is the IDs of the VPCs to which the Partner Attachment is connected to
	VPCIDs []string `json:"vpc_ids,omitempty"`
	// BGP is the BGP configuration of the Partner Attachment
	BGP BGP `json:"bgp,omitempty"`
	// RedundancyZone is the redundancy zone of the Partner Attachment
	RedundancyZone string `json:"redundancy_zone,omitempty"`
	// CreatedAt is time when this Partner Attachment was first created
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ParentUuid is the uuid of the parent Partner Attachment if there is one
	ParentUuid string `json:"parent_uuid,omitempty"`
	// Children is a slice of children Partner Attachment UUIDs associated
	// with the given Partner Attachment if there are any
	Children []string `json:"children,omitempty"`
}

type partnerNetworkConnectAttachmentRoot struct {
	PartnerAttachment *PartnerAttachment `json:"partner_attachment"`
}

type partnerNetworkConnectAttachmentsRoot struct {
	PartnerAttachments []*PartnerAttachment `json:"partner_attachments"`
	Links              *Links               `json:"links"`
	Meta               *Meta                `json:"meta"`
}

type serviceKeyRoot struct {
	ServiceKey *ServiceKey `json:"service_key"`
}

type remoteRoutesRoot struct {
	RemoteRoutes []*RemoteRoute `json:"remote_routes"`
	Links        *Links         `json:"links"`
	Meta         *Meta          `json:"meta"`
}

type BgpAuthKey struct {
	Value string `json:"value"`
}

type bgpAuthKeyRoot struct {
	BgpAuthKey *BgpAuthKey `json:"bgp_auth_key"`
}

type RegenerateServiceKey struct {
}

type regenerateServiceKeyRoot struct {
	RegenerateServiceKey *RegenerateServiceKey `json:"-"`
}

// List returns a list of all Partner Attachment, with optional pagination.
func (s *PartnerAttachmentServiceOp) List(ctx context.Context, opt *ListOptions) ([]*PartnerAttachment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create creates a new Partner Attachment.
func (s *PartnerAttachmentServiceOp) Create(ctx context.Context, create *PartnerAttachmentCreateRequest) (*PartnerAttachment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get returns the details of a Partner Attachment.
func (s *PartnerAttachmentServiceOp) Get(ctx context.Context, id string) (*PartnerAttachment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Update updates a Partner Attachment properties.
func (s *PartnerAttachmentServiceOp) Update(ctx context.Context, id string, update *PartnerAttachmentUpdateRequest) (*PartnerAttachment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete deletes a Partner Attachment.
func (s *PartnerAttachmentServiceOp) Delete(ctx context.Context, id string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *PartnerAttachmentServiceOp) GetServiceKey(ctx context.Context, id string) (*ServiceKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ListRoutes lists all remote routes for a Partner Attachment.
func (s *PartnerAttachmentServiceOp) ListRoutes(ctx context.Context, id string, opt *ListOptions) ([]*RemoteRoute, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetBGPAuthKey returns Partner Attachment bgp auth key
func (s *PartnerAttachmentServiceOp) GetBGPAuthKey(ctx context.Context, iaID string) (*BgpAuthKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// RegenerateServiceKey regenerates the service key of a Partner Attachment.
func (s *PartnerAttachmentServiceOp) RegenerateServiceKey(ctx context.Context, iaID string) (*RegenerateServiceKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
