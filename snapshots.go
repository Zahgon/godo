package godo

import (
	"context"
)

const snapshotBasePath = "v2/snapshots"

// SnapshotsService is an interface for interfacing with the snapshots
// endpoints of the DigitalOcean API
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Snapshots
type SnapshotsService interface {
	List(context.Context, *ListOptions) ([]Snapshot, *Response, error)
	ListVolume(context.Context, *ListOptions) ([]Snapshot, *Response, error)
	ListVolumeSnapshotByRegion(context.Context, string, *ListOptions) ([]Snapshot, *Response, error)
	ListDroplet(context.Context, *ListOptions) ([]Snapshot, *Response, error)
	Get(context.Context, string) (*Snapshot, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// SnapshotsServiceOp handles communication with the snapshot related methods of the
// DigitalOcean API.
type SnapshotsServiceOp struct {
	client *Client
}

var _ SnapshotsService = &SnapshotsServiceOp{}

// Snapshot represents a DigitalOcean Snapshot
type Snapshot struct {
	ID            string   `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	ResourceID    string   `json:"resource_id,omitempty"`
	ResourceType  string   `json:"resource_type,omitempty"`
	Regions       []string `json:"regions,omitempty"`
	MinDiskSize   int      `json:"min_disk_size,omitempty"`
	SizeGigaBytes float64  `json:"size_gigabytes,omitempty"`
	Created       string   `json:"created_at,omitempty"`
	Tags          []string `json:"tags,omitempty"`
}

type snapshotRoot struct {
	Snapshot *Snapshot `json:"snapshot"`
}

type snapshotsRoot struct {
	Snapshots []Snapshot `json:"snapshots"`
	Links     *Links     `json:"links,omitempty"`
	Meta      *Meta      `json:"meta,omitempty"`
}

type listSnapshotOptions struct {
	ResourceType string `url:"resource_type,omitempty"`
	Region       string `url:"region,omitempty"`
}

func (s Snapshot) String() string { _ = "STUB: not implemented"; return "" }

// List lists all the snapshots available.
func (s *SnapshotsServiceOp) List(ctx context.Context, opt *ListOptions) ([]Snapshot, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil,

		// ListDroplet lists all the Droplet snapshots.
		nil
}

func (s *SnapshotsServiceOp) ListDroplet(ctx context.Context, opt *ListOptions) ([]Snapshot, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ListVolume lists all the volume snapshots.
func (s *SnapshotsServiceOp) ListVolume(ctx context.Context, opt *ListOptions) ([]Snapshot, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ListVolumeSnapshotByRegion lists all the volume snapshot for given region
func (s *SnapshotsServiceOp) ListVolumeSnapshotByRegion(ctx context.Context, region string, opt *ListOptions) ([]Snapshot, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get retrieves a snapshot by id.
func (s *SnapshotsServiceOp) Get(ctx context.Context, snapshotID string) (*Snapshot, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil,

		// Delete an snapshot.
		nil
}

func (s *SnapshotsServiceOp) Delete(ctx context.Context, snapshotID string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Helper method for getting an individual snapshot
func (s *SnapshotsServiceOp) get(ctx context.Context, ID string) (*Snapshot, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Helper method for listing snapshots
func (s *SnapshotsServiceOp) list(ctx context.Context, opt *ListOptions, listOpt *listSnapshotOptions) ([]Snapshot, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
