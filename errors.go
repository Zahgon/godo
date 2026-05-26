package godo

// ArgError is an error that represents an error with an input to godo. It
// identifies the argument and the cause (if possible).
type ArgError struct {
	arg    string
	reason string
}

var _ error = &ArgError{}

// NewArgError creates an InputError.
func NewArgError(arg, reason string) *ArgError { _ = "STUB: not implemented"; return nil }

func (e *ArgError) Error() string { _ = "STUB: not implemented"; return "" }
