package godo

import (
	"context"
)

const nfsBasePath = "v2/nfs"

const nfsSnapshotsBasePath = "v2/nfs/snapshots"

// NfsShareStatus represents the status of an NFS share.
type NfsShareStatus string

// Possible states for an NFS share.
const (
	NfsShareCreating = NfsShareStatus("CREATING")
	NfsShareActive   = NfsShareStatus("ACTIVE")
	NfsShareFailed   = NfsShareStatus("FAILED")
	NfsShareDeleted  = NfsShareStatus("DELETED")
)

// NfsSnapshotStatus represents the status of an NFS snapshot.
type NfsSnapshotStatus string

// Possible states for an NFS snapshot.
const (
	NfsSnapshotUnknown  = NfsSnapshotStatus("UNKNOWN")
	NfsSnapshotCreating = NfsSnapshotStatus("CREATING")
	NfsSnapshotActive   = NfsSnapshotStatus("ACTIVE")
	NfsSnapshotFailed   = NfsSnapshotStatus("FAILED")
	NfsSnapshotDeleted  = NfsSnapshotStatus("DELETED")
)

type NfsService interface {
	// List retrieves a list of NFS shares filtered by region
	List(ctx context.Context, opts *ListOptions, region string) ([]*Nfs, *Response, error)
	// Create creates a new NFS share with the provided configuration
	Create(ctx context.Context, nfsCreateRequest *NfsCreateRequest) (*Nfs, *Response, error)
	// Delete removes an NFS share by its ID and region
	Delete(ctx context.Context, nfsShareId string, region string) (*Response, error)
	// Get retrieves a specific NFS share by its ID and region
	Get(ctx context.Context, nfsShareId string, region string) (*Nfs, *Response, error)
	// List retrieves a list of NFS snapshots filtered by an optional share ID and region
	ListSnapshots(ctx context.Context, opts *ListOptions, nfsShareId string, region string) ([]*NfsSnapshot, *Response, error)
	// Get retrieves a specific NFS snapshot by its ID and region
	GetSnapshot(ctx context.Context, nfsSnapshotID string, region string) (*NfsSnapshot, *Response, error)
	// Delete removes an NFS snapshot by its ID and region
	DeleteSnapshot(ctx context.Context, nfsSnapshotID string, region string) (*Response, error)
}

// NfsServiceOp handles communication with the NFS related methods of the
// DigitalOcean API.
type NfsServiceOp struct {
	client *Client
}

var _ NfsService = &NfsServiceOp{}

// Nfs represents a DigitalOcean NFS share
type Nfs struct {
	// ID is the unique identifier for the NFS share
	ID string `json:"id"`
	// Name is the human-readable name for the NFS share
	Name string `json:"name"`
	// SizeGib is the size of the NFS share in gibibytes
	SizeGib int `json:"size_gib"`
	// Region is the datacenter region where the NFS share is located
	Region string `json:"region"`
	// Status represents the current state of the NFS share
	Status NfsShareStatus `json:"status"`
	// CreatedAt is the timestamp when the NFS share was created
	CreatedAt string `json:"created_at"`
	// VpcIDs is a list of VPC IDs that have access to the NFS share
	VpcIDs []string `json:"vpc_ids"`
	// Host is the IP address of the NFS server accessible from the associated VPC
	Host string `json:"host"`
	// MountPath is the path at which the share will be available
	MountPath string `json:"mount_path"`
	//PerformanceTier is the performance tier of the NFS share
	PerformanceTier string `json:"performance_tier"`
}

type NfsSnapshot struct {
	// ID is the unique identifier for the NFS snapshot
	ID string `json:"id"`
	// Name is the human-readable name for the NFS snapshot
	Name string `json:"name"`
	// SizeGib is the size of the NFS snapshot in gibibytes
	SizeGib int `json:"size_gib"`
	// Region is the datacenter region where the NFS snapshot is located
	Region string `json:"region"`
	// Status represents the current status of the NFS snapshot
	Status NfsSnapshotStatus `json:"status"`
	// CreatedAt is the timestamp when the NFS snapshot was created
	CreatedAt string `json:"created_at"`
	// ShareID is the unique identifier of the share from which this snapshot was created.
	ShareID string `json:"share_id"`
}

// NfsCreateRequest represents a request to create an NFS share.
type NfsCreateRequest struct {
	Name            string   `json:"name"`
	SizeGib         int      `json:"size_gib"`
	Region          string   `json:"region"`
	VpcIDs          []string `json:"vpc_ids,omitempty"`
	PerformanceTier string   `json:"performance_tier,omitempty"`
}

// nfsRoot represents a response from the DigitalOcean API
type nfsRoot struct {
	Share *Nfs `json:"share"`
}

// nfsListRoot represents a response from the DigitalOcean API
type nfsListRoot struct {
	Shares []*Nfs `json:"shares,omitempty"`
	Links  *Links `json:"links,omitempty"`
	Meta   *Meta  `json:"meta"`
}

// nfsSnapshotRoot represents a response from the DigitalOcean API
type nfsSnapshotRoot struct {
	Snapshot *NfsSnapshot `json:"snapshot"`
}

// nfsSnapshotListRoot represents a response from the DigitalOcean API
type nfsSnapshotListRoot struct {
	Snapshots []*NfsSnapshot `json:"snapshots,omitempty"`
	Links     *Links         `json:"links,omitempty"`
	Meta      *Meta          `json:"meta"`
}

// nfsOptions represents the query param options for NFS operations
type nfsOptions struct {
	// Region is the datacenter region where the NFS share/shapshot is located
	Region string `url:"region"`
	// ShareID is the unique identifier of the share from which this snapshot was created.
	ShareID string `url:"share_id,omitempty"`
}

// Create creates a new NFS share.
func (s *NfsServiceOp) Create(ctx context.Context, createRequest *NfsCreateRequest) (*Nfs, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get retrieves an NFS share by ID and region.
func (s *NfsServiceOp) Get(ctx context.Context, nfsShareId string, region string) (*Nfs, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// List returns a list of NFS shares.
func (s *NfsServiceOp) List(ctx context.Context, opts *ListOptions, region string) ([]*Nfs, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete deletes an NFS share by ID and region.
func (s *NfsServiceOp) Delete(ctx context.Context, nfsShareId string, region string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get retrieves an NFS snapshot by ID and region.
func (s *NfsServiceOp) GetSnapshot(ctx context.Context, nfsSnapshotID string, region string) (*NfsSnapshot, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// List returns a list of NFS snapshots.
func (s *NfsServiceOp) ListSnapshots(ctx context.Context, opts *ListOptions, nfsShareId, region string) ([]*NfsSnapshot, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete deletes an NFS snapshot by ID and region.
func (s *NfsServiceOp) DeleteSnapshot(ctx context.Context, nfsSnapshotID string, region string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
