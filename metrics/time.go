// Copyright 2013 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import (
	"math"
	"time"
)

const (
	// MinimumTick is the minimum supported time resolution. This has to be
	// at least time.Second in order for the code below to work.
	minimumTick = time.Millisecond
	// second is the Time duration equivalent to one second.
	second = int64(time.Second / minimumTick)
	// The number of nanoseconds per minimum tick.
	nanosPerTick = int64(minimumTick / time.Nanosecond)

	// Earliest is the earliest Time representable. Handy for
	// initializing a high watermark.
	Earliest = Time(math.MinInt64)
	// Latest is the latest Time representable. Handy for initializing
	// a low watermark.
	Latest = Time(math.MaxInt64)
)

// Time is the number of milliseconds since the epoch
// (1970-01-01 00:00 UTC) excluding leap seconds.
type Time int64

// Interval describes an interval between two timestamps.
type Interval struct {
	Start, End Time
}

// Now returns the current time as a Time.
func Now() Time { _ = "STUB: not implemented"; return *new(Time) }

// TimeFromUnix returns the Time equivalent to the Unix Time t
// provided in seconds.
func TimeFromUnix(t int64) Time {
	_ = "STUB: not implemented"
	return *

	// TimeFromUnixNano returns the Time equivalent to the Unix Time
	// t provided in nanoseconds.
	new(Time)
}

func TimeFromUnixNano(t int64) Time { _ = "STUB: not implemented"; return *new(Time) }

// Equal reports whether two Times represent the same instant.
func (t Time) Equal(o Time) bool {
	_ = "STUB: not implemented"

	// Before reports whether the Time t is before o.
	return false
}

func (t Time) Before(o Time) bool {
	_ = "STUB: not implemented"

	// After reports whether the Time t is after o.
	return false
}

func (t Time) After(o Time) bool {
	_ = "STUB: not implemented"

	// Add returns the Time t + d.
	return false
}

func (t Time) Add(d time.Duration) Time { _ = "STUB: not implemented"; return *new(Time) }

// Sub returns the Duration t - o.
func (t Time) Sub(o Time) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// Time returns the time.Time representation of t.
func (t Time) Time() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Unix returns t as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC.
func (t Time) Unix() int64 { _ = "STUB: not implemented"; return 0 }

// UnixNano returns t as a Unix time, the number of nanoseconds elapsed
// since January 1, 1970 UTC.
func (t Time) UnixNano() int64 { _ = "STUB: not implemented"; return 0 }

// The number of digits after the dot.
var dotPrecision = int(math.Log10(float64(second)))

// String returns a string representation of the Time.
func (t Time) String() string { _ = "STUB: not implemented"; return "" }

// MarshalJSON implements the json.Marshaler interface.
func (t Time) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Time) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// If the value was something like -0.1 the negative is lost in the
// parsing because of the leading zero, this ensures that we capture it.
