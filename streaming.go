package godo

import (
	"bufio"
	"io"
)

// SSEEvent is a dispatched Server-Sent Events event. Data is the joined
// value of all "data:" fields, separated by '\n', and is only valid until
// the next call to SSEReader.Next.
type SSEEvent struct {
	Event string
	Data  []byte
	ID    string
	Retry int
}

// SSEReader parses a "text/event-stream" byte stream into SSEEvents.
// Bare-CR line terminators are not supported. Not safe for concurrent use.
type SSEReader struct {
	r       *bufio.Reader
	lastID  string
	scratch []byte
	err     error
}

// NewSSEReader returns an SSEReader that reads from r.
func NewSSEReader(r io.Reader) *SSEReader { _ = "STUB: not implemented"; return nil }

// Next returns the next dispatched event, or io.EOF when the stream
// ends. A final event without a trailing blank line is dispatched on
// the call that hits EOF; the subsequent call returns io.EOF.
func (s *SSEReader) Next() (*SSEEvent, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *SSEReader) makeEvent(eventTyp string, retry int) *SSEEvent {
	_ = "STUB: not implemented"
	return nil
}

func (s *SSEReader) readLine() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func splitField(line []byte) (string, []byte) { _ = "STUB: not implemented"; return "", nil }
