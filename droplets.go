package godo

import (
	"context"
	"errors"
)

const dropletBasePath = "v2/droplets"

var errNoNetworks = errors.New("no networks have been defined")

// DropletsService is an interface for interfacing with the Droplet
// endpoints of the DigitalOcean API
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Droplets
type DropletsService interface {
	List(context.Context, *ListOptions) ([]Droplet, *Response, error)
	ListWithGPUs(context.Context, *ListOptions) ([]Droplet, *Response, error)
	ListByName(context.Context, string, *ListOptions) ([]Droplet, *Response, error)
	ListByTag(context.Context, string, *ListOptions) ([]Droplet, *Response, error)
	Get(context.Context, int) (*Droplet, *Response, error)
	Create(context.Context, *DropletCreateRequest) (*Droplet, *Response, error)
	CreateMultiple(context.Context, *DropletMultiCreateRequest) ([]Droplet, *Response, error)
	Delete(context.Context, int) (*Response, error)
	DeleteByTag(context.Context, string) (*Response, error)
	Kernels(context.Context, int, *ListOptions) ([]Kernel, *Response, error)
	Snapshots(context.Context, int, *ListOptions) ([]Image, *Response, error)
	Backups(context.Context, int, *ListOptions) ([]Image, *Response, error)
	Actions(context.Context, int, *ListOptions) ([]Action, *Response, error)
	Neighbors(context.Context, int) ([]Droplet, *Response, error)
	GetBackupPolicy(context.Context, int) (*DropletBackupPolicy, *Response, error)
	ListBackupPolicies(context.Context, *ListOptions) (map[int]*DropletBackupPolicy, *Response, error)
	ListSupportedBackupPolicies(context.Context) ([]*SupportedBackupPolicy, *Response, error)
	ListAssociatedResourcesForDeletion(context.Context, int) (*DropletAssociatedResources, *Response, error)
}

// DropletsServiceOp handles communication with the Droplet related methods of the
// DigitalOcean API.
type DropletsServiceOp struct {
	client *Client
}

var _ DropletsService = &DropletsServiceOp{}

// Droplet represents a DigitalOcean Droplet
type Droplet struct {
	ID               int           `json:"id,float64,omitempty"`
	Name             string        `json:"name,omitempty"`
	Memory           int           `json:"memory,omitempty"`
	Vcpus            int           `json:"vcpus,omitempty"`
	Disk             int           `json:"disk,omitempty"`
	Region           *Region       `json:"region,omitempty"`
	Image            *Image        `json:"image,omitempty"`
	Size             *Size         `json:"size,omitempty"`
	SizeSlug         string        `json:"size_slug,omitempty"`
	BackupIDs        []int         `json:"backup_ids,omitempty"`
	NextBackupWindow *BackupWindow `json:"next_backup_window,omitempty"`
	SnapshotIDs      []int         `json:"snapshot_ids,omitempty"`
	Features         []string      `json:"features,omitempty"`
	Locked           bool          `json:"locked,bool,omitempty"`
	Status           string        `json:"status,omitempty"`
	Networks         *Networks     `json:"networks,omitempty"`
	Created          string        `json:"created_at,omitempty"`
	Kernel           *Kernel       `json:"kernel,omitempty"`
	Tags             []string      `json:"tags,omitempty"`
	VolumeIDs        []string      `json:"volume_ids"`
	VPCUUID          string        `json:"vpc_uuid,omitempty"`
}

// PublicIPv4 returns the public IPv4 address for the Droplet.
func (d *Droplet) PublicIPv4() (string, error) { _ = "STUB: not implemented"; return "", nil }

// PrivateIPv4 returns the private IPv4 address for the Droplet.
func (d *Droplet) PrivateIPv4() (string, error) { _ = "STUB: not implemented"; return "", nil }

// PublicIPv6 returns the public IPv6 address for the Droplet.
func (d *Droplet) PublicIPv6() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Kernel object
type Kernel struct {
	ID      int    `json:"id,float64,omitempty"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

// BackupWindow object
type BackupWindow struct {
	Start *Timestamp `json:"start,omitempty"`
	End   *Timestamp `json:"end,omitempty"`
}

// Convert Droplet to a string
func (d Droplet) String() string { _ = "STUB: not implemented"; return "" }

// URN returns the droplet ID in a valid DO API URN form.
func (d Droplet) URN() string { _ = "STUB: not implemented"; return "" }

// DropletRoot represents a Droplet root
type dropletRoot struct {
	Droplet *Droplet `json:"droplet"`
	Links   *Links   `json:"links,omitempty"`
}

type dropletsRoot struct {
	Droplets []Droplet `json:"droplets"`
	Links    *Links    `json:"links"`
	Meta     *Meta     `json:"meta"`
}

type kernelsRoot struct {
	Kernels []Kernel `json:"kernels,omitempty"`
	Links   *Links   `json:"links"`
	Meta    *Meta    `json:"meta"`
}

type dropletSnapshotsRoot struct {
	Snapshots []Image `json:"snapshots,omitempty"`
	Links     *Links  `json:"links"`
	Meta      *Meta   `json:"meta"`
}

type backupsRoot struct {
	Backups []Image `json:"backups,omitempty"`
	Links   *Links  `json:"links"`
	Meta    *Meta   `json:"meta"`
}

// DropletCreateImage identifies an image for the create request. It prefers slug over ID.
type DropletCreateImage struct {
	ID   int
	Slug string
}

// MarshalJSON returns either the slug or id of the image. It returns the id
// if the slug is empty.
func (d DropletCreateImage) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DropletCreateVolume identifies a volume to attach for the create request.
type DropletCreateVolume struct {
	ID string
	// Deprecated: You must pass the volume's ID when creating a Droplet.
	Name string
}

// MarshalJSON returns an object with either the ID or name of the volume. It
// prefers the ID over the name.
func (d DropletCreateVolume) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DropletCreateSSHKey identifies a SSH Key for the create request. It prefers fingerprint over ID.
type DropletCreateSSHKey struct {
	ID          int
	Fingerprint string
}

// MarshalJSON returns either the fingerprint or id of the ssh key. It returns
// the id if the fingerprint is empty.
func (d DropletCreateSSHKey) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DropletCreateRequest represents a request to create a Droplet.
type DropletCreateRequest struct {
	Name              string                      `json:"name"`
	Region            string                      `json:"region"`
	Size              string                      `json:"size"`
	Image             DropletCreateImage          `json:"image"`
	SSHKeys           []DropletCreateSSHKey       `json:"ssh_keys"`
	Backups           bool                        `json:"backups"`
	IPv6              bool                        `json:"ipv6"`
	PrivateNetworking bool                        `json:"private_networking"`
	Monitoring        bool                        `json:"monitoring"`
	UserData          string                      `json:"user_data,omitempty"`
	Volumes           []DropletCreateVolume       `json:"volumes,omitempty"`
	Tags              []string                    `json:"tags"`
	VPCUUID           string                      `json:"vpc_uuid,omitempty"`
	WithDropletAgent  *bool                       `json:"with_droplet_agent,omitempty"`
	BackupPolicy      *DropletBackupPolicyRequest `json:"backup_policy,omitempty"`
	PublicNetworking  *bool                       `json:"public_networking,omitempty"`
}

// DropletMultiCreateRequest is a request to create multiple Droplets.
type DropletMultiCreateRequest struct {
	Names             []string                    `json:"names"`
	Region            string                      `json:"region"`
	Size              string                      `json:"size"`
	Image             DropletCreateImage          `json:"image"`
	SSHKeys           []DropletCreateSSHKey       `json:"ssh_keys"`
	Backups           bool                        `json:"backups"`
	IPv6              bool                        `json:"ipv6"`
	PrivateNetworking bool                        `json:"private_networking"`
	Monitoring        bool                        `json:"monitoring"`
	UserData          string                      `json:"user_data,omitempty"`
	Tags              []string                    `json:"tags"`
	VPCUUID           string                      `json:"vpc_uuid,omitempty"`
	WithDropletAgent  *bool                       `json:"with_droplet_agent,omitempty"`
	BackupPolicy      *DropletBackupPolicyRequest `json:"backup_policy,omitempty"`
	PublicNetworking  *bool                       `json:"public_networking,omitempty"`
}

// DropletBackupPolicyRequest defines the backup policy when creating a Droplet.
type DropletBackupPolicyRequest struct {
	Plan    string `json:"plan,omitempty"`
	Weekday string `json:"weekday,omitempty"`
	Hour    *int   `json:"hour,omitempty"`
}

// DropletAssociatedResource represents a billable resource associated with a Droplet.
type DropletAssociatedResource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Cost string `json:"cost"`
}

// DropletAssociatedResources represents the associated billable resources that can be destroyed along with a Droplet.
type DropletAssociatedResources struct {
	ReservedIPs     []*DropletAssociatedResource `json:"reserved_ips"`
	FloatingIPs     []*DropletAssociatedResource `json:"floating_ips"`
	Snapshots       []*DropletAssociatedResource `json:"snapshots"`
	Volumes         []*DropletAssociatedResource `json:"volumes"`
	VolumeSnapshots []*DropletAssociatedResource `json:"volume_snapshots"`
}

func (a DropletAssociatedResources) String() string { _ = "STUB: not implemented"; return "" }

func (d DropletCreateRequest) String() string { _ = "STUB: not implemented"; return "" }

func (d DropletMultiCreateRequest) String() string { _ = "STUB: not implemented"; return "" }

// Networks represents the Droplet's Networks.
type Networks struct {
	V4 []NetworkV4 `json:"v4,omitempty"`
	V6 []NetworkV6 `json:"v6,omitempty"`
}

// NetworkV4 represents a DigitalOcean IPv4 Network.
type NetworkV4 struct {
	IPAddress string `json:"ip_address,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Gateway   string `json:"gateway,omitempty"`
	Type      string `json:"type,omitempty"`
}

func (n NetworkV4) String() string { _ = "STUB: not implemented"; return "" }

// NetworkV6 represents a DigitalOcean IPv6 network.
type NetworkV6 struct {
	IPAddress string `json:"ip_address,omitempty"`
	Netmask   int    `json:"netmask,omitempty"`
	Gateway   string `json:"gateway,omitempty"`
	Type      string `json:"type,omitempty"`
}

func (n NetworkV6) String() string { _ = "STUB: not implemented"; return "" }

// Performs a list request given a path.
func (s *DropletsServiceOp) list(ctx context.Context, path string) ([]Droplet, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// List all Droplets.
func (s *DropletsServiceOp) List(ctx context.Context, opt *ListOptions) ([]Droplet, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ListWithGPUs lists all Droplets with GPUs.
func (s *DropletsServiceOp) ListWithGPUs(ctx context.Context, opt *ListOptions) ([]Droplet, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ListByName lists all Droplets filtered by name returning only exact matches.
// It is case-insensitive
func (s *DropletsServiceOp) ListByName(ctx context.Context, name string, opt *ListOptions) ([]Droplet, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ListByTag lists all Droplets matched by a Tag.
func (s *DropletsServiceOp) ListByTag(ctx context.Context, tag string, opt *ListOptions) ([]Droplet, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get individual Droplet.
func (s *DropletsServiceOp) Get(ctx context.Context, dropletID int) (*Droplet, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create Droplet
func (s *DropletsServiceOp) Create(ctx context.Context, createRequest *DropletCreateRequest) (*Droplet, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// CreateMultiple creates multiple Droplets.
func (s *DropletsServiceOp) CreateMultiple(ctx context.Context, createRequest *DropletMultiCreateRequest) ([]Droplet, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Performs a delete request given a path
func (s *DropletsServiceOp) delete(ctx context.Context, path string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete Droplet.
func (s *DropletsServiceOp) Delete(ctx context.Context, dropletID int) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteByTag deletes Droplets matched by a Tag.
func (s *DropletsServiceOp) DeleteByTag(ctx context.Context, tag string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Kernels lists kernels available for a Droplet.
func (s *DropletsServiceOp) Kernels(ctx context.Context, dropletID int, opt *ListOptions) ([]Kernel, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Actions lists the actions for a Droplet.
func (s *DropletsServiceOp) Actions(ctx context.Context, dropletID int, opt *ListOptions) ([]Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Backups lists the backups for a Droplet.
func (s *DropletsServiceOp) Backups(ctx context.Context, dropletID int, opt *ListOptions) ([]Image, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Snapshots lists the snapshots available for a Droplet.
func (s *DropletsServiceOp) Snapshots(ctx context.Context, dropletID int, opt *ListOptions) ([]Image, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Neighbors lists the neighbors for a Droplet.
func (s *DropletsServiceOp) Neighbors(ctx context.Context, dropletID int) ([]Droplet, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ListAssociatedResourcesForDeletion lists a Droplet's associated resources that can be destroyed along with the Droplet.
// Associated resources include reserved IPs, floating IPs, snapshots, volumes, and volume snapshots.
func (s *DropletsServiceOp) ListAssociatedResourcesForDeletion(ctx context.Context, dropletID int) (*DropletAssociatedResources, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *DropletsServiceOp) dropletActionStatus(ctx context.Context, uri string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// DropletBackupPolicy defines the information about a droplet's backup policy.
type DropletBackupPolicy struct {
	DropletID        int                        `json:"droplet_id,omitempty"`
	BackupEnabled    bool                       `json:"backup_enabled,omitempty"`
	BackupPolicy     *DropletBackupPolicyConfig `json:"backup_policy,omitempty"`
	NextBackupWindow *BackupWindow              `json:"next_backup_window,omitempty"`
}

// DropletBackupPolicyConfig defines the backup policy for a Droplet.
type DropletBackupPolicyConfig struct {
	Plan                string `json:"plan,omitempty"`
	Weekday             string `json:"weekday,omitempty"`
	Hour                int    `json:"hour,omitempty"`
	WindowLengthHours   int    `json:"window_length_hours,omitempty"`
	RetentionPeriodDays int    `json:"retention_period_days,omitempty"`
}

// dropletBackupPolicyRoot represents a DropletBackupPolicy root
type dropletBackupPolicyRoot struct {
	DropletBackupPolicy *DropletBackupPolicy `json:"policy,omitempty"`
}

type dropletBackupPoliciesRoot struct {
	DropletBackupPolicies map[int]*DropletBackupPolicy `json:"policies,omitempty"`
	Links                 *Links                       `json:"links,omitempty"`
	Meta                  *Meta                        `json:"meta"`
}

// Get individual droplet backup policy.
func (s *DropletsServiceOp) GetBackupPolicy(ctx context.Context, dropletID int) (*DropletBackupPolicy, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// List all droplet backup policies.
func (s *DropletsServiceOp) ListBackupPolicies(ctx context.Context, opt *ListOptions) (map[int]*DropletBackupPolicy, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type SupportedBackupPolicy struct {
	Name                 string   `json:"name,omitempty"`
	PossibleWindowStarts []int    `json:"possible_window_starts,omitempty"`
	WindowLengthHours    int      `json:"window_length_hours,omitempty"`
	RetentionPeriodDays  int      `json:"retention_period_days,omitempty"`
	PossibleDays         []string `json:"possible_days,omitempty"`
}

type dropletSupportedBackupPoliciesRoot struct {
	SupportedBackupPolicies []*SupportedBackupPolicy `json:"supported_policies,omitempty"`
}

// List supported droplet backup policies.
func (s *DropletsServiceOp) ListSupportedBackupPolicies(ctx context.Context) ([]*SupportedBackupPolicy, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
