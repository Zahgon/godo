package godo

import (
	"context"
)

const certificatesBasePath = "/v2/certificates"

// CertificatesService is an interface for managing certificates with the DigitalOcean API.
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Certificates
type CertificatesService interface {
	Get(context.Context, string) (*Certificate, *Response, error)
	List(context.Context, *ListOptions) ([]Certificate, *Response, error)
	ListByName(context.Context, string, *ListOptions) ([]Certificate, *Response, error)
	Create(context.Context, *CertificateRequest) (*Certificate, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// Certificate represents a DigitalOcean certificate configuration.
type Certificate struct {
	ID              string   `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	DNSNames        []string `json:"dns_names,omitempty"`
	NotAfter        string   `json:"not_after,omitempty"`
	SHA1Fingerprint string   `json:"sha1_fingerprint,omitempty"`
	Created         string   `json:"created_at,omitempty"`
	State           string   `json:"state,omitempty"`
	Type            string   `json:"type,omitempty"`
}

// CertificateRequest represents configuration for a new certificate.
type CertificateRequest struct {
	Name             string   `json:"name,omitempty"`
	DNSNames         []string `json:"dns_names,omitempty"`
	PrivateKey       string   `json:"private_key,omitempty"`
	LeafCertificate  string   `json:"leaf_certificate,omitempty"`
	CertificateChain string   `json:"certificate_chain,omitempty"`
	Type             string   `json:"type,omitempty"`
}

type certificateRoot struct {
	Certificate *Certificate `json:"certificate"`
}

type certificatesRoot struct {
	Certificates []Certificate `json:"certificates"`
	Links        *Links        `json:"links"`
	Meta         *Meta         `json:"meta"`
}

// CertificatesServiceOp handles communication with certificates methods of the DigitalOcean API.
type CertificatesServiceOp struct {
	client *Client
}

var _ CertificatesService = &CertificatesServiceOp{}

// Get an existing certificate by its identifier.
func (c *CertificatesServiceOp) Get(ctx context.Context, cID string) (*Certificate, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// List all certificates.
func (c *CertificatesServiceOp) List(ctx context.Context, opt *ListOptions) ([]Certificate, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *CertificatesServiceOp) ListByName(ctx context.Context, name string, opt *ListOptions) ([]Certificate, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create a new certificate with provided configuration.
func (c *CertificatesServiceOp) Create(ctx context.Context, cr *CertificateRequest) (*Certificate, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete a certificate by its identifier.
func (c *CertificatesServiceOp) Delete(ctx context.Context, cID string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
