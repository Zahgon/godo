package godo

import (
	"context"
)

// Links manages links that are returned along with a List
type Links struct {
	Pages   *Pages       `json:"pages,omitempty"`
	Actions []LinkAction `json:"actions,omitempty"`
}

// Pages are pages specified in Links
type Pages struct {
	First string `json:"first,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Last  string `json:"last,omitempty"`
	Next  string `json:"next,omitempty"`
}

// LinkAction is a pointer to an action
type LinkAction struct {
	ID   int    `json:"id,omitempty"`
	Rel  string `json:"rel,omitempty"`
	HREF string `json:"href,omitempty"`
}

// CurrentPage is current page of the list
func (l *Links) CurrentPage() (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// NextPageToken is the page token to request the next page of the list
		nil
}

func (l *Links) NextPageToken() (string, error) { _ = "STUB: not implemented"; return "", nil }

// PrevPageToken is the page token to request the previous page of the list
func (l *Links) PrevPageToken() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (p *Pages) current() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *Pages) nextPageToken() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (p *Pages) prevPageToken() (string, error) { _ = "STUB: not implemented"; return "", nil }

// IsLastPage returns true if the current page is the last
func (l *Links) IsLastPage() bool { _ = "STUB: not implemented"; return false }

func (p *Pages) isLast() bool { _ = "STUB: not implemented"; return false }

func pageForURL(urlText string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func pageTokenFromURL(urlText string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Get a link action by id.
func (la *LinkAction) Get(ctx context.Context, client *Client) (*Action, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
