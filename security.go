package godo

import (
	"context"
)

const (
	securityScansBasePath     = "v2/security/scans"
	securityScansFindingsPath = "findings"
	securityAffectedResources = "affected_resources"
)

// SecurityService is an interface for interacting with the CSPM endpoints of
// the DigitalOcean API.
type SecurityService interface {
	CreateScan(context.Context, *CreateScanRequest) (*Scan, *Response, error)
	ListScans(context.Context, *ListOptions) ([]*Scan, *Response, error)
	GetScan(context.Context, string, *ScanFindingsOptions) (*Scan, *Response, error)
	GetLatestScan(context.Context, *ScanFindingsOptions) (*Scan, *Response, error)
	ListFindingAffectedResources(context.Context, *ListFindingAffectedResourcesRequest, *ListOptions) ([]*AffectedResource, *Response, error)
}

// SecurityServiceOp handles communication with security scan related methods of the DigitalOcean API.
type SecurityServiceOp struct {
	client *Client
}

var _ SecurityService = &SecurityServiceOp{}

// CreateScanRequest contains the request payload to create a scan.
type CreateScanRequest struct {
	Resources []string `json:"resources,omitempty"`
}

// ScanFindingsOptions contains the query parameters for paginating and
// filtering scan findings.
type ScanFindingsOptions struct {
	ListOptions
	Type     string `url:"type,omitempty"`
	Severity string `url:"severity,omitempty"`
}

// Scan represents a CSPM scan.
type Scan struct {
	ID        string         `json:"id,omitempty"`
	Status    string         `json:"status,omitempty"`
	CreatedAt string         `json:"created_at,omitempty"`
	Findings  []*ScanFinding `json:"findings,omitempty"`
}

func (s *Scan) Completed() bool { _ = "STUB: not implemented"; return false }

// ScanFinding represents a finding within a scan.
type ScanFinding struct {
	RuleUUID               string                `json:"rule_uuid,omitempty"`
	Name                   string                `json:"name,omitempty"`
	Details                string                `json:"details,omitempty"`
	FoundAt                string                `json:"found_at,omitempty"`
	Severity               string                `json:"severity,omitempty"`
	BusinessImpact         string                `json:"business_impact,omitempty"`
	TechnicalDetails       string                `json:"technical_details,omitempty"`
	MitigationSteps        []*ScanMitigationStep `json:"mitigation_steps,omitempty"`
	AffectedResourcesCount int                   `json:"affected_resources_count,omitempty"`
}

// ScanMitigationStep represents a mitigation step for a scan finding.
type ScanMitigationStep struct {
	Step        int    `json:"step,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

// An AffectedResource represents a resource affected by a scan finding.
type AffectedResource struct {
	URN  string `json:"urn,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type scanRoot struct {
	Scan *Scan `json:"scan"`
}

type scansRoot struct {
	Scans []*Scan `json:"scans"`
	Links *Links  `json:"links"`
	Meta  *Meta   `json:"meta"`
}

type affectedResourcesRoot struct {
	AffectedResources []*AffectedResource `json:"affected_resources"`
	Links             *Links              `json:"links"`
	Meta              *Meta               `json:"meta"`
}

// CreateScan initiates a new CSPM scan.
func (s *SecurityServiceOp) CreateScan(ctx context.Context, createScanRequest *CreateScanRequest) (*Scan, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ListScans lists all CSPM scans.
func (s *SecurityServiceOp) ListScans(ctx context.Context, opts *ListOptions) ([]*Scan, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetScan retrieves a scan by its UUID with optional findings filters.
func (s *SecurityServiceOp) GetScan(ctx context.Context, scanUUID string, opts *ScanFindingsOptions) (*Scan, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetLatestScan retrieves the latest scan with optional findings filters.
func (s *SecurityServiceOp) GetLatestScan(ctx context.Context, opts *ScanFindingsOptions) (*Scan, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ListFindingAffectedResourcesRequest contains the fields to list the
// affected resources for a scan finding.
type ListFindingAffectedResourcesRequest struct {
	ScanUUID    string
	FindingUUID string
}

// ListFindingAffectedResources lists the affected resources for a scan
// finding.
func (s *SecurityServiceOp) ListFindingAffectedResources(ctx context.Context, listFindingAffectedResourcesRequest *ListFindingAffectedResourcesRequest, opts *ListOptions) ([]*AffectedResource, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
