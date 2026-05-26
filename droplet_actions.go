package godo

import (
	"context"
)

// ActionRequest represents DigitalOcean Action Request
type ActionRequest map[string]interface{}

// DropletActionsService is an interface for interfacing with the Droplet actions
// endpoints of the DigitalOcean API
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Droplet-Actions
type DropletActionsService interface {
	Shutdown(context.Context, int) (*Action, *Response, error)
	ShutdownByTag(context.Context, string) ([]Action, *Response, error)
	PowerOff(context.Context, int) (*Action, *Response, error)
	PowerOffByTag(context.Context, string) ([]Action, *Response, error)
	PowerOn(context.Context, int) (*Action, *Response, error)
	PowerOnByTag(context.Context, string) ([]Action, *Response, error)
	PowerCycle(context.Context, int) (*Action, *Response, error)
	PowerCycleByTag(context.Context, string) ([]Action, *Response, error)
	Reboot(context.Context, int) (*Action, *Response, error)
	Restore(context.Context, int, int) (*Action, *Response, error)
	Resize(context.Context, int, string, bool) (*Action, *Response, error)
	Rename(context.Context, int, string) (*Action, *Response, error)
	Snapshot(context.Context, int, string) (*Action, *Response, error)
	SnapshotByTag(context.Context, string, string) ([]Action, *Response, error)
	EnableBackups(context.Context, int) (*Action, *Response, error)
	EnableBackupsByTag(context.Context, string) ([]Action, *Response, error)
	EnableBackupsWithPolicy(context.Context, int, *DropletBackupPolicyRequest) (*Action, *Response, error)
	ChangeBackupPolicy(context.Context, int, *DropletBackupPolicyRequest) (*Action, *Response, error)
	DisableBackups(context.Context, int) (*Action, *Response, error)
	DisableBackupsByTag(context.Context, string) ([]Action, *Response, error)
	PasswordReset(context.Context, int) (*Action, *Response, error)
	RebuildByImageID(context.Context, int, int) (*Action, *Response, error)
	RebuildByImageSlug(context.Context, int, string) (*Action, *Response, error)
	ChangeKernel(context.Context, int, int) (*Action, *Response, error)
	EnableIPv6(context.Context, int) (*Action, *Response, error)
	EnableIPv6ByTag(context.Context, string) ([]Action, *Response, error)
	EnablePrivateNetworking(context.Context, int) (*Action, *Response, error)
	EnablePrivateNetworkingByTag(context.Context, string) ([]Action, *Response, error)
	Get(context.Context, int, int) (*Action, *Response, error)
	GetByURI(context.Context, string) (*Action, *Response, error)
}

// DropletActionsServiceOp handles communication with the Droplet action related
// methods of the DigitalOcean API.
type DropletActionsServiceOp struct {
	client *Client
}

var _ DropletActionsService = &DropletActionsServiceOp{}

// Shutdown a Droplet
func (s *DropletActionsServiceOp) Shutdown(ctx context.Context, id int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ShutdownByTag shuts down Droplets matched by a Tag.
func (s *DropletActionsServiceOp) ShutdownByTag(ctx context.Context, tag string) ([]Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// PowerOff a Droplet
func (s *DropletActionsServiceOp) PowerOff(ctx context.Context, id int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// PowerOffByTag powers off Droplets matched by a Tag.
func (s *DropletActionsServiceOp) PowerOffByTag(ctx context.Context, tag string) ([]Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// PowerOn a Droplet
func (s *DropletActionsServiceOp) PowerOn(ctx context.Context, id int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// PowerOnByTag powers on Droplets matched by a Tag.
func (s *DropletActionsServiceOp) PowerOnByTag(ctx context.Context, tag string) ([]Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// PowerCycle a Droplet
func (s *DropletActionsServiceOp) PowerCycle(ctx context.Context, id int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// PowerCycleByTag power cycles Droplets matched by a Tag.
func (s *DropletActionsServiceOp) PowerCycleByTag(ctx context.Context, tag string) ([]Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Reboot a Droplet
func (s *DropletActionsServiceOp) Reboot(ctx context.Context, id int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Restore an image to a Droplet
func (s *DropletActionsServiceOp) Restore(ctx context.Context, id, imageID int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Resize a Droplet
func (s *DropletActionsServiceOp) Resize(ctx context.Context, id int, sizeSlug string, resizeDisk bool) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Rename a Droplet
func (s *DropletActionsServiceOp) Rename(ctx context.Context, id int, name string) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Snapshot a Droplet.
func (s *DropletActionsServiceOp) Snapshot(ctx context.Context, id int, name string) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// SnapshotByTag snapshots Droplets matched by a Tag.
func (s *DropletActionsServiceOp) SnapshotByTag(ctx context.Context, tag string, name string) ([]Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// EnableBackups enables backups for a Droplet.
func (s *DropletActionsServiceOp) EnableBackups(ctx context.Context, id int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// EnableBackupsByTag enables backups for Droplets matched by a Tag.
func (s *DropletActionsServiceOp) EnableBackupsByTag(ctx context.Context, tag string) ([]Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// EnableBackupsWithPolicy enables droplet's backup with a backup policy applied.
func (s *DropletActionsServiceOp) EnableBackupsWithPolicy(ctx context.Context, id int, policy *DropletBackupPolicyRequest) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ChangeBackupPolicy updates a backup policy when backups are enabled.
func (s *DropletActionsServiceOp) ChangeBackupPolicy(ctx context.Context, id int, policy *DropletBackupPolicyRequest) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// DisableBackups disables backups for a Droplet.
func (s *DropletActionsServiceOp) DisableBackups(ctx context.Context, id int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// DisableBackupsByTag disables backups for Droplet matched by a Tag.
func (s *DropletActionsServiceOp) DisableBackupsByTag(ctx context.Context, tag string) ([]Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// PasswordReset resets the password for a Droplet.
func (s *DropletActionsServiceOp) PasswordReset(ctx context.Context, id int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// RebuildByImageID rebuilds a Droplet from an image with a given id.
func (s *DropletActionsServiceOp) RebuildByImageID(ctx context.Context, id, imageID int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// RebuildByImageSlug rebuilds a Droplet from an Image matched by a given Slug.
func (s *DropletActionsServiceOp) RebuildByImageSlug(ctx context.Context, id int, slug string) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ChangeKernel changes the kernel for a Droplet.
func (s *DropletActionsServiceOp) ChangeKernel(ctx context.Context, id, kernelID int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// EnableIPv6 enables IPv6 for a Droplet.
func (s *DropletActionsServiceOp) EnableIPv6(ctx context.Context, id int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// EnableIPv6ByTag enables IPv6 for Droplets matched by a Tag.
func (s *DropletActionsServiceOp) EnableIPv6ByTag(ctx context.Context, tag string) ([]Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// EnablePrivateNetworking enables private networking for a Droplet.
func (s *DropletActionsServiceOp) EnablePrivateNetworking(ctx context.Context, id int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// EnablePrivateNetworkingByTag enables private networking for Droplets matched by a Tag.
func (s *DropletActionsServiceOp) EnablePrivateNetworkingByTag(ctx context.Context, tag string) ([]Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *DropletActionsServiceOp) doAction(ctx context.Context, id int, request *ActionRequest) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *DropletActionsServiceOp) doActionByTag(ctx context.Context, tag string, request *ActionRequest) ([]Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get an action for a particular Droplet by id.
func (s *DropletActionsServiceOp) Get(ctx context.Context, dropletID, actionID int) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetByURI gets an action for a particular Droplet by URI.
func (s *DropletActionsServiceOp) GetByURI(ctx context.Context, rawurl string) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *DropletActionsServiceOp) get(ctx context.Context, path string) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func dropletActionPath(dropletID int) string { _ = "STUB: not implemented"; return "" }

func dropletActionPathByTag(tag string) string { _ = "STUB: not implemented"; return "" }
