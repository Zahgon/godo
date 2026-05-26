package util

import (
	"context"

	"github.com/digitalocean/godo"
)

const (
	// availableFailure is the amount of times we can fail before deciding
	// the check for available is a total failure. This can help account
	// for servers randomly not answering.
	availableFailure = 3
)

// WaitForAvailable waits for a image to become available
func WaitForAvailable(ctx context.Context, client *godo.Client, monitorURI string) error {
	_ = "STUB: not implemented"
	return nil
}
