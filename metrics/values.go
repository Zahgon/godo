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

// A SampleValue is a representation of a value for a given sample at a given time.
type SampleValue float64

// UnmarshalJSON implements json.Unmarshaler.
func (v *SampleValue) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
func (v SampleValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (v SampleValue) String() string { _ = "STUB: not implemented"; return "" }

// Equal returns true if the value of v and o is equal or if both are NaN. Note
// that v==o is false if both are NaN. If you want the conventional float
// behavior, use == to compare two SampleValues.
func (v SampleValue) Equal(o SampleValue) bool { _ = "STUB: not implemented"; return false }

// SamplePair pairs a SampleValue with a Timestamp.
type SamplePair struct {
	Timestamp Time
	Value     SampleValue
}

func (s SamplePair) String() string { _ = "STUB: not implemented"; return "" }

// UnmarshalJSON implements json.Unmarshaler.
func (s *SamplePair) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
func (s SamplePair) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SampleStream is a stream of Values belonging to an attached COWMetric.
type SampleStream struct {
	Metric Metric       `json:"metric"`
	Values []SamplePair `json:"values"`
}

func (ss SampleStream) String() string { _ = "STUB: not implemented"; return "" }
