package godo

import (
	"context"
)

const tagsBasePath = "v2/tags"

// TagsService is an interface for interfacing with the tags
// endpoints of the DigitalOcean API
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Tags
type TagsService interface {
	List(context.Context, *ListOptions) ([]Tag, *Response, error)
	Get(context.Context, string) (*Tag, *Response, error)
	Create(context.Context, *TagCreateRequest) (*Tag, *Response, error)
	Delete(context.Context, string) (*Response, error)

	TagResources(context.Context, string, *TagResourcesRequest) (*Response, error)
	UntagResources(context.Context, string, *UntagResourcesRequest) (*Response, error)
}

// TagsServiceOp handles communication with tag related method of the
// DigitalOcean API.
type TagsServiceOp struct {
	client *Client
}

var _ TagsService = &TagsServiceOp{}

// ResourceType represents a class of resource, currently only droplet are supported
type ResourceType string

const (
	// DropletResourceType holds the string representing our ResourceType of Droplet.
	DropletResourceType ResourceType = "droplet"
	// ImageResourceType holds the string representing our ResourceType of Image.
	ImageResourceType ResourceType = "image"
	// VolumeResourceType holds the string representing our ResourceType of Volume.
	VolumeResourceType ResourceType = "volume"
	// LoadBalancerResourceType holds the string representing our ResourceType of LoadBalancer.
	LoadBalancerResourceType ResourceType = "load_balancer"
	// VolumeSnapshotResourceType holds the string representing our ResourceType for storage Snapshots.
	VolumeSnapshotResourceType ResourceType = "volume_snapshot"
	// DatabaseResourceType holds the string representing our ResourceType of Database.
	DatabaseResourceType ResourceType = "database"
)

// Resource represent a single resource for associating/disassociating with tags
type Resource struct {
	ID   string       `json:"resource_id,omitempty"`
	Type ResourceType `json:"resource_type,omitempty"`
}

// TaggedResources represent the set of resources a tag is attached to
type TaggedResources struct {
	Count           int                             `json:"count"`
	LastTaggedURI   string                          `json:"last_tagged_uri,omitempty"`
	Droplets        *TaggedDropletsResources        `json:"droplets,omitempty"`
	Images          *TaggedImagesResources          `json:"images"`
	Volumes         *TaggedVolumesResources         `json:"volumes"`
	VolumeSnapshots *TaggedVolumeSnapshotsResources `json:"volume_snapshots"`
	Databases       *TaggedDatabasesResources       `json:"databases"`
}

// TaggedDropletsResources represent the droplet resources a tag is attached to
type TaggedDropletsResources struct {
	Count         int      `json:"count,float64,omitempty"`
	LastTagged    *Droplet `json:"last_tagged,omitempty"`
	LastTaggedURI string   `json:"last_tagged_uri,omitempty"`
}

// TaggedResourcesData represent the generic resources a tag is attached to
type TaggedResourcesData struct {
	Count         int    `json:"count,float64,omitempty"`
	LastTaggedURI string `json:"last_tagged_uri,omitempty"`
}

// TaggedImagesResources represent the image resources a tag is attached to
type TaggedImagesResources TaggedResourcesData

// TaggedVolumesResources represent the volume resources a tag is attached to
type TaggedVolumesResources TaggedResourcesData

// TaggedVolumeSnapshotsResources represent the volume snapshot resources a tag is attached to
type TaggedVolumeSnapshotsResources TaggedResourcesData

// TaggedDatabasesResources represent the database resources a tag is attached to
type TaggedDatabasesResources TaggedResourcesData

// Tag represent DigitalOcean tag
type Tag struct {
	Name      string           `json:"name,omitempty"`
	Resources *TaggedResources `json:"resources,omitempty"`
}

// TagCreateRequest represents the JSON structure of a request of that type.
type TagCreateRequest struct {
	Name string `json:"name"`
}

// TagResourcesRequest represents the JSON structure of a request of that type.
type TagResourcesRequest struct {
	Resources []Resource `json:"resources"`
}

// UntagResourcesRequest represents the JSON structure of a request of that type.
type UntagResourcesRequest struct {
	Resources []Resource `json:"resources"`
}

type tagsRoot struct {
	Tags  []Tag  `json:"tags"`
	Links *Links `json:"links"`
	Meta  *Meta  `json:"meta"`
}

type tagRoot struct {
	Tag *Tag `json:"tag"`
}

// List all tags
func (s *TagsServiceOp) List(ctx context.Context, opt *ListOptions) ([]Tag, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get a single tag
func (s *TagsServiceOp) Get(ctx context.Context, name string) (*Tag, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create a new tag
func (s *TagsServiceOp) Create(ctx context.Context, createRequest *TagCreateRequest) (*Tag, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete an existing tag
func (s *TagsServiceOp) Delete(ctx context.Context, name string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TagResources associates resources with a given Tag.
func (s *TagsServiceOp) TagResources(ctx context.Context, name string, tagRequest *TagResourcesRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UntagResources dissociates resources with a given Tag.
func (s *TagsServiceOp) UntagResources(ctx context.Context, name string, untagRequest *UntagResourcesRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
