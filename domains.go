package godo

import (
	"context"
)

const domainsBasePath = "v2/domains"

// DomainsService is an interface for managing DNS with the DigitalOcean API.
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Domains and
// https://docs.digitalocean.com/reference/api/api-reference/#tag/Domain-Records
type DomainsService interface {
	List(context.Context, *ListOptions) ([]Domain, *Response, error)
	Get(context.Context, string) (*Domain, *Response, error)
	Create(context.Context, *DomainCreateRequest) (*Domain, *Response, error)
	Delete(context.Context, string) (*Response, error)

	Records(context.Context, string, *ListOptions) ([]DomainRecord, *Response, error)
	RecordsByType(context.Context, string, string, *ListOptions) ([]DomainRecord, *Response, error)
	RecordsByName(context.Context, string, string, *ListOptions) ([]DomainRecord, *Response, error)
	RecordsByTypeAndName(context.Context, string, string, string, *ListOptions) ([]DomainRecord, *Response, error)
	Record(context.Context, string, int) (*DomainRecord, *Response, error)
	DeleteRecord(context.Context, string, int) (*Response, error)
	EditRecord(context.Context, string, int, *DomainRecordEditRequest) (*DomainRecord, *Response, error)
	CreateRecord(context.Context, string, *DomainRecordEditRequest) (*DomainRecord, *Response, error)
}

// DomainsServiceOp handles communication with the domain related methods of the
// DigitalOcean API.
type DomainsServiceOp struct {
	client *Client
}

var _ DomainsService = &DomainsServiceOp{}

// Domain represents a DigitalOcean domain
type Domain struct {
	Name     string `json:"name"`
	TTL      int    `json:"ttl"`
	ZoneFile string `json:"zone_file"`
}

// domainRoot represents a response from the DigitalOcean API
type domainRoot struct {
	Domain *Domain `json:"domain"`
}

type domainsRoot struct {
	Domains []Domain `json:"domains"`
	Links   *Links   `json:"links"`
	Meta    *Meta    `json:"meta"`
}

// DomainCreateRequest represents a request to create a domain.
type DomainCreateRequest struct {
	Name      string `json:"name"`
	IPAddress string `json:"ip_address,omitempty"`
}

// DomainRecordRoot is the root of an individual Domain Record response
type domainRecordRoot struct {
	DomainRecord *DomainRecord `json:"domain_record"`
}

// DomainRecordsRoot is the root of a group of Domain Record responses
type domainRecordsRoot struct {
	DomainRecords []DomainRecord `json:"domain_records"`
	Links         *Links         `json:"links"`
}

// DomainRecord represents a DigitalOcean DomainRecord
type DomainRecord struct {
	ID       int    `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	Name     string `json:"name,omitempty"`
	Data     string `json:"data,omitempty"`
	Priority int    `json:"priority"`
	Port     int    `json:"port"`
	TTL      int    `json:"ttl,omitempty"`
	Weight   int    `json:"weight"`
	Flags    int    `json:"flags"`
	Tag      string `json:"tag,omitempty"`
}

// DomainRecordEditRequest represents a request to update a domain record.
type DomainRecordEditRequest struct {
	Type     string `json:"type,omitempty"`
	Name     string `json:"name,omitempty"`
	Data     string `json:"data,omitempty"`
	Priority int    `json:"priority"`
	Port     int    `json:"port"`
	TTL      int    `json:"ttl,omitempty"`
	Weight   int    `json:"weight"`
	Flags    int    `json:"flags"`
	Tag      string `json:"tag,omitempty"`
}

func (d Domain) String() string { _ = "STUB: not implemented"; return "" }

// URN returns the domain name in a valid DO API URN form.
func (d Domain) URN() string { _ = "STUB: not implemented"; return "" }

// List all domains.
func (s DomainsServiceOp) List(ctx context.Context, opt *ListOptions) ([]Domain, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get individual domain. It requires a non-empty domain name.
func (s *DomainsServiceOp) Get(ctx context.Context, name string) (*Domain, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create a new domain
func (s *DomainsServiceOp) Create(ctx context.Context, createRequest *DomainCreateRequest) (*Domain, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete domain
func (s *DomainsServiceOp) Delete(ctx context.Context, name string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Converts a DomainRecord to a string.
func (d DomainRecord) String() string { _ = "STUB: not implemented"; return "" }

// Converts a DomainRecordEditRequest to a string.
func (d DomainRecordEditRequest) String() string { _ = "STUB: not implemented"; return "" }

// Records returns a slice of DomainRecord for a domain.
func (s *DomainsServiceOp) Records(ctx context.Context, domain string, opt *ListOptions) ([]DomainRecord, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// RecordsByType returns a slice of DomainRecord for a domain matched by record type.
func (s *DomainsServiceOp) RecordsByType(ctx context.Context, domain, ofType string, opt *ListOptions) ([]DomainRecord, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// RecordsByName returns a slice of DomainRecord for a domain matched by record name.
func (s *DomainsServiceOp) RecordsByName(ctx context.Context, domain, name string, opt *ListOptions) ([]DomainRecord, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// RecordsByTypeAndName returns a slice of DomainRecord for a domain matched by record type and name.
func (s *DomainsServiceOp) RecordsByTypeAndName(ctx context.Context, domain, ofType, name string, opt *ListOptions) ([]DomainRecord, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Record returns the record id from a domain
func (s *DomainsServiceOp) Record(ctx context.Context, domain string, id int) (*DomainRecord, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// DeleteRecord deletes a record from a domain identified by id
func (s *DomainsServiceOp) DeleteRecord(ctx context.Context, domain string, id int) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EditRecord edits a record using a DomainRecordEditRequest
func (s *DomainsServiceOp) EditRecord(ctx context.Context,
	domain string,
	id int,
	editRequest *DomainRecordEditRequest,
) (*DomainRecord, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// CreateRecord creates a record using a DomainRecordEditRequest
func (s *DomainsServiceOp) CreateRecord(ctx context.Context,
	domain string,
	createRequest *DomainRecordEditRequest) (*DomainRecord, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Performs a domain records request given a path.
func (s *DomainsServiceOp) records(ctx context.Context, path string) ([]DomainRecord, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
