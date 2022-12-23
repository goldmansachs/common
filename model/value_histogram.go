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

package model

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type IntString uint64

func (v IntString) String() string {
	return fmt.Sprintf("%d", v)
}

func (v IntString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

func (v *IntString) UnmarshalJSON(b []byte) error {
	if len(b) < 2 || b[0] != '"' || b[len(b)-1] != '"' {
		return fmt.Errorf("int value must be a quoted string")
	}
	i, err := strconv.ParseInt(string(b[1:len(b)-1]), 10, 64)
	if err != nil {
		return err
	}
	*v = IntString(i)
	return nil
}

func (v IntString) Equal(o IntString) bool {
	return v == o
}

type FloatString float64

func (v FloatString) String() string {
	return strconv.FormatFloat(float64(v), 'f', -1, 64)
}

func (v FloatString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

func (v *FloatString) UnmarshalJSON(b []byte) error {
	if len(b) < 2 || b[0] != '"' || b[len(b)-1] != '"' {
		return fmt.Errorf("float value must be a quoted string")
	}
	f, err := strconv.ParseFloat(string(b[1:len(b)-1]), 64)
	if err != nil {
		return err
	}
	*v = FloatString(f)
	return nil
}

func (v FloatString) Equal(o FloatString) bool {
	return v == o
}

type HistogramBucket struct {
	Boundaries int
	Lower      FloatString
	Upper      FloatString
	Count      IntString
}

func (s HistogramBucket) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(s.Boundaries)
	if err != nil {
		return nil, err
	}
	l, err := json.Marshal(s.Lower)
	if err != nil {
		return nil, err
	}
	u, err := json.Marshal(s.Upper)
	if err != nil {
		return nil, err
	}
	c, err := json.Marshal(s.Count)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("[%s,%s,%s,%s]", b, l, u, c)), nil
}

func (s *HistogramBucket) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&s.Boundaries, &s.Lower, &s.Upper, &s.Count}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields: %d != %d", g, e)
	}
	return nil
}

func (s *HistogramBucket) Equal(o *HistogramBucket) bool {
	return s == o || (s.Boundaries == o.Boundaries && s.Lower.Equal(o.Lower) && s.Upper.Equal(o.Upper) && s.Count.Equal(o.Count))
}

type HistogramBuckets []*HistogramBucket

func (s HistogramBuckets) Equal(o HistogramBuckets) bool {
	if len(s) != len(o) {
		return false
	}

	for i, bucket := range s {
		if !bucket.Equal(o[i]) {
			return false
		}
	}
	return true
}

type SampleHistogram struct {
	Count   IntString        `json:"count"`
	Sum     FloatString      `json:"sum"`
	Buckets HistogramBuckets `json:"buckets"`
}

func (s SampleHistogram) String() string {
	return fmt.Sprintf("Count: %d, Sum: %f, Buckets: [redacted]", s.Count, s.Sum)
}

func (s *SampleHistogram) Equal(o *SampleHistogram) bool {
	return s == o || (s.Count.Equal(o.Count) && s.Sum.Equal(o.Sum) && s.Buckets.Equal(o.Buckets))
}

type SampleHistogramPair struct {
	Timestamp Time
	Histogram SampleHistogram
}

func (s SampleHistogramPair) MarshalJSON() ([]byte, error) {
	t, err := json.Marshal(s.Timestamp)
	if err != nil {
		return nil, err
	}
	v, err := json.Marshal(s.Histogram)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("[%s,%s]", t, v)), nil
}

func (s *SampleHistogramPair) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&s.Timestamp, &s.Histogram}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields: %d != %d", g, e)
	}
	return nil
}

func (s SampleHistogramPair) String() string {
	return fmt.Sprintf("%s @[%s]", s.Histogram, s.Timestamp)
}

func (s *SampleHistogramPair) Equal(o *SampleHistogramPair) bool {
	return s == o || (s.Histogram.Equal(&o.Histogram) && s.Timestamp.Equal(o.Timestamp))
}
