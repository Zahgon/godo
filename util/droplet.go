package util

import (
	"context"

	"github.com/digitalocean/godo"
)

const (
	// activeFailure is the amount of times we can fail before deciding
	// the check for active is a total failure. This can help account
	// for servers randomly not answering.
	activeFailure = 3
)

// WaitForActive waits for a droplet to become active
func WaitForActive(ctx context.Context, client *godo.Client, monitorURI string) error {
	_ = "STUB: not implemented"
	return nil
}
