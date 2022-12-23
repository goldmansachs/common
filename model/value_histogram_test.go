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
	"reflect"
	"testing"
)

func TestSampleHistogramPairJSON(t *testing.T) {
	input := []struct {
		plain string
		value SampleHistogramPair
	}{
		{
			plain: `[1234.567,{"count":"1","sum":"4500","buckets":[[0,"4466.7196729968955","4870.992343051145","1"]]}]`,
			value: SampleHistogramPair{
				Histogram: SampleHistogram{
					Count: 1,
					Sum:   4500,
					Buckets: HistogramBuckets{
						{
							Boundaries: 0,
							Lower:      4466.7196729968955,
							Upper:      4870.992343051145,
							Count:      1,
						},
					},
				},
				Timestamp: 1234567,
			},
		},
	}

	for _, test := range input {
		b, err := json.Marshal(test.value)
		if err != nil {
			t.Error(err)
			continue
		}

		if string(b) != test.plain {
			t.Errorf("encoding error: expected %q, got %q", test.plain, b)
			continue
		}

		var sp SampleHistogramPair
		err = json.Unmarshal(b, &sp)
		if err != nil {
			t.Error(err)
			continue
		}

		if !sp.Equal(&test.value) {
			t.Errorf("decoding error: expected %v, got %v", test.value, sp)
		}
	}
}

func TestSampleHistogramJSON(t *testing.T) {
	input := []struct {
		plain string
		value Sample
	}{
		{
			plain: `{"metric":{"__name__":"test_metric"},"histogram":[1234.567,{"count":"1","sum":"4500","buckets":[[0,"4466.7196729968955","4870.992343051145","1"]]}]}`,
			value: Sample{
				Metric: Metric{
					MetricNameLabel: "test_metric",
				},
				Histogram: SampleHistogram{
					Count: 1,
					Sum:   4500,
					Buckets: HistogramBuckets{
						{
							Boundaries: 0,
							Lower:      4466.7196729968955,
							Upper:      4870.992343051145,
							Count:      1,
						},
					},
				},
				Timestamp: 1234567,
			},
		},
	}

	for _, test := range input {
		b, err := json.Marshal(test.value)
		if err != nil {
			t.Error(err)
			continue
		}

		if string(b) != test.plain {
			t.Errorf("encoding error: expected %q, got %q", test.plain, b)
			continue
		}

		var sv Sample
		err = json.Unmarshal(b, &sv)
		if err != nil {
			t.Error(err)
			continue
		}

		if !reflect.DeepEqual(sv, test.value) {
			t.Errorf("decoding error: expected %v, got %v", test.value, sv)
		}
	}
}

func TestVectorHistogramJSON(t *testing.T) {
	input := []struct {
		plain string
		value Vector
	}{
		{
			plain: `[{"metric":{"__name__":"test_metric"},"histogram":[1234.567,{"count":"1","sum":"4500","buckets":[[0,"4466.7196729968955","4870.992343051145","1"]]}]}]`,
			value: Vector{&Sample{
				Metric: Metric{
					MetricNameLabel: "test_metric",
				},
				Histogram: SampleHistogram{
					Count: 1,
					Sum:   4500,
					Buckets: HistogramBuckets{
						{
							Boundaries: 0,
							Lower:      4466.7196729968955,
							Upper:      4870.992343051145,
							Count:      1,
						},
					},
				},
				Timestamp: 1234567,
			}},
		},
		{
			plain: `[{"metric":{"__name__":"test_metric"},"histogram":[1234.567,{"count":"1","sum":"4500","buckets":[[0,"4466.7196729968955","4870.992343051145","1"]]}]},{"metric":{"foo":"bar"},"histogram":[1.234,{"count":"1","sum":"4500","buckets":[[0,"4466.7196729968955","4870.992343051145","1"]]}]}]`,
			value: Vector{
				&Sample{
					Metric: Metric{
						MetricNameLabel: "test_metric",
					},
					Histogram: SampleHistogram{
						Count: 1,
						Sum:   4500,
						Buckets: HistogramBuckets{
							{
								Boundaries: 0,
								Lower:      4466.7196729968955,
								Upper:      4870.992343051145,
								Count:      1,
							},
						},
					},
					Timestamp: 1234567,
				},
				&Sample{
					Metric: Metric{
						"foo": "bar",
					},
					Histogram: SampleHistogram{
						Count: 1,
						Sum:   4500,
						Buckets: HistogramBuckets{
							{
								Boundaries: 0,
								Lower:      4466.7196729968955,
								Upper:      4870.992343051145,
								Count:      1,
							},
						},
					},
					Timestamp: 1234,
				},
			},
		},
	}

	for _, test := range input {
		b, err := json.Marshal(test.value)
		if err != nil {
			t.Error(err)
			continue
		}

		if string(b) != test.plain {
			t.Errorf("encoding error: expected %q, got %q", test.plain, b)
			continue
		}

		var vec Vector
		err = json.Unmarshal(b, &vec)
		if err != nil {
			t.Error(err)
			continue
		}

		if !reflect.DeepEqual(vec, test.value) {
			t.Errorf("decoding error: expected %v, got %v", test.value, vec)
		}
	}
}
