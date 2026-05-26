package godo

import (
	"context"
	"time"
)

const (
	storageBasePath  = "v2"
	storageAllocPath = storageBasePath + "/volumes"
	storageSnapPath  = storageBasePath + "/snapshots"
)

// StorageService is an interface for interfacing with the storage
// endpoints of the Digital Ocean API.
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Block-Storage
type StorageService interface {
	ListVolumes(context.Context, *ListVolumeParams) ([]Volume, *Response, error)
	GetVolume(context.Context, string) (*Volume, *Response, error)
	CreateVolume(context.Context, *VolumeCreateRequest) (*Volume, *Response, error)
	DeleteVolume(context.Context, string) (*Response, error)
	ListSnapshots(ctx context.Context, volumeID string, opts *ListOptions) ([]Snapshot, *Response, error)
	GetSnapshot(context.Context, string) (*Snapshot, *Response, error)
	CreateSnapshot(context.Context, *SnapshotCreateRequest) (*Snapshot, *Response, error)
	DeleteSnapshot(context.Context, string) (*Response, error)
}

// StorageServiceOp handles communication with the storage volumes related methods of the
// DigitalOcean API.
type StorageServiceOp struct {
	client *Client
}

// ListVolumeParams stores the options you can set for a ListVolumeCall
type ListVolumeParams struct {
	Region      string       `json:"region"`
	Name        string       `json:"name"`
	ListOptions *ListOptions `json:"list_options,omitempty"`
}

var _ StorageService = &StorageServiceOp{}

// Volume represents a Digital Ocean block store volume.
type Volume struct {
	ID              string    `json:"id"`
	Region          *Region   `json:"region"`
	Name            string    `json:"name"`
	SizeGigaBytes   int64     `json:"size_gigabytes"`
	Description     string    `json:"description"`
	DropletIDs      []int     `json:"droplet_ids"`
	CreatedAt       time.Time `json:"created_at"`
	FilesystemType  string    `json:"filesystem_type"`
	FilesystemLabel string    `json:"filesystem_label"`
	Tags            []string  `json:"tags"`
}

func (f Volume) String() string { _ = "STUB: not implemented"; return "" }

// URN returns the volume ID as a valid DO API URN
func (f Volume) URN() string { _ = "STUB: not implemented"; return "" }

type storageVolumesRoot struct {
	Volumes []Volume `json:"volumes"`
	Links   *Links   `json:"links"`
	Meta    *Meta    `json:"meta"`
}

type storageVolumeRoot struct {
	Volume *Volume `json:"volume"`
	Links  *Links  `json:"links,omitempty"`
}

// VolumeCreateRequest represents a request to create a block store
// volume.
type VolumeCreateRequest struct {
	Region          string   `json:"region"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	SizeGigaBytes   int64    `json:"size_gigabytes"`
	SnapshotID      string   `json:"snapshot_id"`
	FilesystemType  string   `json:"filesystem_type"`
	FilesystemLabel string   `json:"filesystem_label"`
	Tags            []string `json:"tags"`
}

// ListVolumes lists all storage volumes.
func (svc *StorageServiceOp) ListVolumes(ctx context.Context, params *ListVolumeParams) ([]Volume, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// CreateVolume creates a storage volume. The name must be unique.
func (svc *StorageServiceOp) CreateVolume(ctx context.Context, createRequest *VolumeCreateRequest) (*Volume, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetVolume retrieves an individual storage volume.
func (svc *StorageServiceOp) GetVolume(ctx context.Context, id string) (*Volume, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// DeleteVolume deletes a storage volume.
func (svc *StorageServiceOp) DeleteVolume(ctx context.Context, id string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SnapshotCreateRequest represents a request to create a block store
// volume.
type SnapshotCreateRequest struct {
	VolumeID    string   `json:"volume_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

// ListSnapshots lists all snapshots related to a storage volume.
func (svc *StorageServiceOp) ListSnapshots(ctx context.Context, volumeID string, opt *ListOptions) ([]Snapshot, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// CreateSnapshot creates a snapshot of a storage volume.
func (svc *StorageServiceOp) CreateSnapshot(ctx context.Context, createRequest *SnapshotCreateRequest) (*Snapshot, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetSnapshot retrieves an individual snapshot.
func (svc *StorageServiceOp) GetSnapshot(ctx context.Context, id string) (*Snapshot, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// DeleteSnapshot deletes a snapshot.
func (svc *StorageServiceOp) DeleteSnapshot(ctx context.Context, id string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
