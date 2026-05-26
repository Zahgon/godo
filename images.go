package godo

import (
	"context"
)

const imageBasePath = "v2/images"

// ImagesService is an interface for interfacing with the images
// endpoints of the DigitalOcean API
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Images
type ImagesService interface {
	List(context.Context, *ListOptions) ([]Image, *Response, error)
	ListDistribution(ctx context.Context, opt *ListOptions) ([]Image, *Response, error)
	ListApplication(ctx context.Context, opt *ListOptions) ([]Image, *Response, error)
	ListUser(ctx context.Context, opt *ListOptions) ([]Image, *Response, error)
	ListByTag(ctx context.Context, tag string, opt *ListOptions) ([]Image, *Response, error)
	GetByID(context.Context, int) (*Image, *Response, error)
	GetBySlug(context.Context, string) (*Image, *Response, error)
	Create(context.Context, *CustomImageCreateRequest) (*Image, *Response, error)
	Update(context.Context, int, *ImageUpdateRequest) (*Image, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// ImagesServiceOp handles communication with the image related methods of the
// DigitalOcean API.
type ImagesServiceOp struct {
	client *Client
}

var _ ImagesService = &ImagesServiceOp{}

// Image represents a DigitalOcean Image
type Image struct {
	ID            int      `json:"id,float64,omitempty"`
	Name          string   `json:"name,omitempty"`
	Type          string   `json:"type,omitempty"`
	Distribution  string   `json:"distribution,omitempty"`
	Slug          string   `json:"slug,omitempty"`
	Public        bool     `json:"public,omitempty"`
	Regions       []string `json:"regions,omitempty"`
	MinDiskSize   int      `json:"min_disk_size,omitempty"`
	SizeGigaBytes float64  `json:"size_gigabytes,omitempty"`
	Created       string   `json:"created_at,omitempty"`
	Description   string   `json:"description,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Status        string   `json:"status,omitempty"`
	ErrorMessage  string   `json:"error_message,omitempty"`
}

// ImageUpdateRequest represents a request to update an image.
type ImageUpdateRequest struct {
	Name         string `json:"name,omitempty"`
	Distribution string `json:"distribution,omitempty"`
	Description  string `json:"description,omitempty"`
}

// CustomImageCreateRequest represents a request to create a custom image.
type CustomImageCreateRequest struct {
	Name         string   `json:"name"`
	Url          string   `json:"url"`
	Region       string   `json:"region"`
	Distribution string   `json:"distribution,omitempty"`
	Description  string   `json:"description,omitempty"`
	Tags         []string `json:"tags,omitempty"`
}

type imageRoot struct {
	Image *Image
}

type imagesRoot struct {
	Images []Image
	Links  *Links `json:"links"`
	Meta   *Meta  `json:"meta"`
}

type listImageOptions struct {
	Private bool   `url:"private,omitempty"`
	Type    string `url:"type,omitempty"`
	Tag     string `url:"tag_name,omitempty"`
}

func (i Image) String() string { _ = "STUB: not implemented"; return "" }

// List lists all the images available.
func (s *ImagesServiceOp) List(ctx context.Context, opt *ListOptions) ([]Image, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil,

		// ListDistribution lists all the distribution images.
		nil
}

func (s *ImagesServiceOp) ListDistribution(ctx context.Context, opt *ListOptions) ([]Image, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ListApplication lists all the application images.
func (s *ImagesServiceOp) ListApplication(ctx context.Context, opt *ListOptions) ([]Image, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ListUser lists all the user images.
func (s *ImagesServiceOp) ListUser(ctx context.Context, opt *ListOptions) ([]Image, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ListByTag lists all images with a specific tag applied.
func (s *ImagesServiceOp) ListByTag(ctx context.Context, tag string, opt *ListOptions) ([]Image, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetByID retrieves an image by id.
func (s *ImagesServiceOp) GetByID(ctx context.Context, imageID int) (*Image, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetBySlug retrieves an image by slug.
func (s *ImagesServiceOp) GetBySlug(ctx context.Context, slug string) (*Image, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create a new image
func (s *ImagesServiceOp) Create(ctx context.Context, createRequest *CustomImageCreateRequest) (*Image, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Update an image name.
func (s *ImagesServiceOp) Update(ctx context.Context, imageID int, updateRequest *ImageUpdateRequest) (*Image, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete an image.
func (s *ImagesServiceOp) Delete(ctx context.Context, imageID int) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Helper method for getting an individual image
func (s *ImagesServiceOp) get(ctx context.Context, ID interface{}) (*Image, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Helper method for listing images
func (s *ImagesServiceOp) list(ctx context.Context, opt *ListOptions, listOpt *listImageOptions) ([]Image, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
