package godo

import (
	"context"
	"time"
)

const (
	vectorDatabaseBasePath        = "/v2/vector-databases"
	vectorDatabaseSinglePath      = vectorDatabaseBasePath + "/%s"
	vectorDatabaseBackupsPath     = vectorDatabaseBasePath + "/%s/backups"
	vectorDatabaseRestorePath     = vectorDatabaseBasePath + "/%s/backups/%s/restore"
	vectorDatabaseCredentialsPath = vectorDatabaseBasePath + "/%s/credentials"
	vectorDatabaseResizePath      = vectorDatabaseBasePath + "/%s/resize"
	vectorDatabaseTagsPath        = vectorDatabaseBasePath + "/%s/tags"
)

// VectorDBsService provides access to the DigitalOcean Vector DB API.
//
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Vector-Databases
type VectorDBsService interface {
	List(context.Context, *ListOptions) ([]VectorDB, *Response, error)
	Get(context.Context, string) (*VectorDB, *Response, error)
	Create(context.Context, *VectorDBCreateRequest) (*VectorDB, *Response, error)
	Update(context.Context, string, *VectorDBUpdateRequest) (*VectorDB, *Response, error)
	Delete(context.Context, string) (*Response, error)
	Resize(context.Context, string, *VectorDBResizeRequest) (*VectorDB, *Response, error)
	UpdateTags(context.Context, string, *VectorDBUpdateTagsRequest) (*VectorDB, *Response, error)
	GetCredentials(context.Context, string) (*VectorDBAdminCredentials, *Response, error)
	ListBackups(context.Context, string) ([]VectorDBBackup, *Response, error)
	RestoreBackup(context.Context, string, string, *VectorDBRestoreBackupRequest) (*VectorDBRestoreBackupResponse, *Response, error)
	GetRestoreStatus(context.Context, string, string) (*VectorDBRestoreStatus, *Response, error)
}

// VectorDBsServiceOp handles communication with the Vector DB related methods
// of the DigitalOcean API.
type VectorDBsServiceOp struct {
	client *Client
}

var _ VectorDBsService = &VectorDBsServiceOp{}

// VectorDB represents a provisioned vector database instance.
type VectorDB struct {
	ID        string             `json:"id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Region    string             `json:"region,omitempty"`
	OwnerUUID string             `json:"owner_uuid,omitempty"`
	Status    string             `json:"status,omitempty"`
	Config    *VectorDBConfig    `json:"config,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
	Endpoints *VectorDBEndpoints `json:"endpoints,omitempty"`
	Size      string             `json:"size,omitempty"`
	Tags      []string           `json:"tags,omitempty"`
}

// VectorDBConfig holds optional, advanced cluster settings.
type VectorDBConfig struct {
	DefaultQuantization string `json:"default_quantization,omitempty"`
	EnableAutoSchema    bool   `json:"enable_auto_schema,omitempty"`
	WeaviateVersion     string `json:"weaviate_version,omitempty"`
}

// VectorDBEndpoints contains the connection endpoints for a vector database instance.
type VectorDBEndpoints struct {
	HTTP string `json:"http,omitempty"`
	GRPC string `json:"grpc,omitempty"`
}

// VectorDBBackup represents a single backup of a vector database.
type VectorDBBackup struct {
	BackupID    string    `json:"backup_id,omitempty"`
	Status      string    `json:"status,omitempty"`
	StartedAt   time.Time `json:"started_at,omitempty"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
}

// VectorDBCreateRequest represents a request to create a vector database.
type VectorDBCreateRequest struct {
	Name   string   `json:"name,omitempty"`
	Region string   `json:"region,omitempty"`
	Size   string   `json:"size,omitempty"`
	Tags   []string `json:"tags,omitempty"`
}

// VectorDBUpdateRequest represents a request to update a vector database.
type VectorDBUpdateRequest struct {
	ID     string          `json:"id,omitempty"`
	Config *VectorDBConfig `json:"config,omitempty"`
}

// VectorDBResizeRequest represents a request to resize a vector database.
type VectorDBResizeRequest struct {
	ID   string `json:"id,omitempty"`
	Size string `json:"size,omitempty"`
}

// VectorDBUpdateTagsRequest represents a request to update tags on a vector database.
type VectorDBUpdateTagsRequest struct {
	ID   string   `json:"id,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

// VectorDBAdminCredentials represents admin credentials for a vector database.
type VectorDBAdminCredentials struct {
	UserID   string `json:"user_id,omitempty"`
	APIToken string `json:"api_token,omitempty"`
}

// VectorDBRestoreBackupRequest represents a request to restore a vector database from a backup.
type VectorDBRestoreBackupRequest struct {
	ID       string `json:"id,omitempty"`
	BackupID string `json:"backup_id,omitempty"`
}

// VectorDBRestoreBackupResponse represents the initial response from a restore operation.
type VectorDBRestoreBackupResponse struct {
	BackupID string `json:"backup_id,omitempty"`
	Status   string `json:"status,omitempty"`
}

// VectorDBRestoreStatus represents the current status of a restore operation.
type VectorDBRestoreStatus struct {
	BackupID string `json:"backup_id,omitempty"`
	Status   string `json:"status,omitempty"`
	Error    string `json:"error,omitempty"`
}

type vectorDBRoot struct {
	VectorDB *VectorDB `json:"vector_db"`
}

type vectorDBsRoot struct {
	VectorDBs []VectorDB `json:"vector_dbs"`
	Total     int64      `json:"total"`
}

type vectorDBBackupsRoot struct {
	Backups []VectorDBBackup `json:"backups"`
}

type vectorDBAdminCredentialsRoot struct {
	UserID   string `json:"user_id"`
	APIToken string `json:"api_token"`
}

type vectorDBRestoreBackupRoot struct {
	BackupID string `json:"backup_id"`
	Status   string `json:"status"`
}

type vectorDBRestoreStatusRoot struct {
	BackupID string `json:"backup_id"`
	Status   string `json:"status"`
	Error    string `json:"error,omitempty"`
}

// List returns a list of vector databases visible with the caller's API token.
func (svc *VectorDBsServiceOp) List(ctx context.Context, opts *ListOptions) ([]VectorDB, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get retrieves the details of a vector database.
func (svc *VectorDBsServiceOp) Get(ctx context.Context, id string) (*VectorDB, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create creates a vector database.
func (svc *VectorDBsServiceOp) Create(ctx context.Context, create *VectorDBCreateRequest) (*VectorDB, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Update updates a vector database's configuration.
func (svc *VectorDBsServiceOp) Update(ctx context.Context, id string, update *VectorDBUpdateRequest) (*VectorDB, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Delete deletes a vector database. There is no way to recover an instance once
// it has been destroyed.
func (svc *VectorDBsServiceOp) Delete(ctx context.Context, id string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Resize resizes a vector database to a new resource tier.
func (svc *VectorDBsServiceOp) Resize(ctx context.Context, id string, resize *VectorDBResizeRequest) (*VectorDB, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// UpdateTags replaces all tags on a vector database.
func (svc *VectorDBsServiceOp) UpdateTags(ctx context.Context, id string, update *VectorDBUpdateTagsRequest) (*VectorDB, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetCredentials retrieves admin credentials for a vector database.
func (svc *VectorDBsServiceOp) GetCredentials(ctx context.Context, id string) (*VectorDBAdminCredentials, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ListBackups returns the available backups for a vector database.
// Only backups with status SUCCESS are returned.
func (svc *VectorDBsServiceOp) ListBackups(ctx context.Context, id string) ([]VectorDBBackup, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// RestoreBackup initiates a restore from a backup.
// The restore is performed asynchronously; use GetRestoreStatus to monitor progress.
func (svc *VectorDBsServiceOp) RestoreBackup(ctx context.Context, id, backupID string, restore *VectorDBRestoreBackupRequest) (*VectorDBRestoreBackupResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetRestoreStatus returns the current status of a restore operation.
func (svc *VectorDBsServiceOp) GetRestoreStatus(ctx context.Context, id, backupID string) (*VectorDBRestoreStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
