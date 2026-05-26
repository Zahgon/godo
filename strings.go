package godo

import (
	"io"
	"reflect"
)

var timestampType = reflect.TypeOf(Timestamp{})

// ResourceWithURN is an interface for interfacing with the types
// that implement the URN method.
type ResourceWithURN interface {
	URN() string
}

// ToURN converts the resource type and ID to a valid DO API URN.
func ToURN(resourceType string, id interface{}) string { _ = "STUB: not implemented"; return "" }

// Stringify attempts to create a string representation of DigitalOcean types
func Stringify(message interface{}) string { _ = "STUB: not implemented"; return "" }

// stringifyValue was graciously cargoculted from the goprotubuf library
func stringifyValue(w io.Writer, val reflect.Value) { _ = "STUB: not implemented"; return }

func stringifySlice(w io.Writer, v reflect.Value) { _ = "STUB: not implemented"; return }

func stringifyMap(w io.Writer, v reflect.Value) { _ = "STUB: not implemented"; return }

// Sort the keys so that the output is stable

func stringifyStruct(w io.Writer, v reflect.Value) { _ = "STUB: not implemented"; return }

// special handling of Timestamp values
