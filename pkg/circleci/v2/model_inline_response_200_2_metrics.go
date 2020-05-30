// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// InlineResponse2002Metrics Metrics relating to a workflow job's runs.
type InlineResponse2002Metrics struct {
	// The ratio of successful runs / total runs.
	SuccessRate float32 `json:"success_rate"`
	// The total number of runs.
	TotalRuns int64 `json:"total_runs"`
	// The number of failed runs.
	FailedRuns int64 `json:"failed_runs"`
	// The number of successful runs.
	SuccessfulRuns int64 `json:"successful_runs"`
	// The average number of job runs per day.
	Throughput float32 `json:"throughput"`
	// The total credits consumed by the job in the aggregation window.
	TotalCreditsUsed int64                                    `json:"total_credits_used"`
	DurationMetrics  InlineResponse2002MetricsDurationMetrics `json:"duration_metrics"`
}

// NewInlineResponse2002Metrics instantiates a new InlineResponse2002Metrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2002Metrics(successRate float32, totalRuns int64, failedRuns int64, successfulRuns int64, throughput float32, totalCreditsUsed int64, durationMetrics InlineResponse2002MetricsDurationMetrics) *InlineResponse2002Metrics {
	this := InlineResponse2002Metrics{}
	this.SuccessRate = successRate
	this.TotalRuns = totalRuns
	this.FailedRuns = failedRuns
	this.SuccessfulRuns = successfulRuns
	this.Throughput = throughput
	this.TotalCreditsUsed = totalCreditsUsed
	this.DurationMetrics = durationMetrics
	return &this
}

// NewInlineResponse2002MetricsWithDefaults instantiates a new InlineResponse2002Metrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2002MetricsWithDefaults() *InlineResponse2002Metrics {
	this := InlineResponse2002Metrics{}
	return &this
}

// GetSuccessRate returns the SuccessRate field value
func (o *InlineResponse2002Metrics) GetSuccessRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SuccessRate
}

// GetSuccessRateOk returns a tuple with the SuccessRate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Metrics) GetSuccessRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessRate, true
}

// SetSuccessRate sets field value
func (o *InlineResponse2002Metrics) SetSuccessRate(v float32) {
	o.SuccessRate = v
}

// GetTotalRuns returns the TotalRuns field value
func (o *InlineResponse2002Metrics) GetTotalRuns() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalRuns
}

// GetTotalRunsOk returns a tuple with the TotalRuns field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Metrics) GetTotalRunsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRuns, true
}

// SetTotalRuns sets field value
func (o *InlineResponse2002Metrics) SetTotalRuns(v int64) {
	o.TotalRuns = v
}

// GetFailedRuns returns the FailedRuns field value
func (o *InlineResponse2002Metrics) GetFailedRuns() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FailedRuns
}

// GetFailedRunsOk returns a tuple with the FailedRuns field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Metrics) GetFailedRunsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailedRuns, true
}

// SetFailedRuns sets field value
func (o *InlineResponse2002Metrics) SetFailedRuns(v int64) {
	o.FailedRuns = v
}

// GetSuccessfulRuns returns the SuccessfulRuns field value
func (o *InlineResponse2002Metrics) GetSuccessfulRuns() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SuccessfulRuns
}

// GetSuccessfulRunsOk returns a tuple with the SuccessfulRuns field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Metrics) GetSuccessfulRunsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessfulRuns, true
}

// SetSuccessfulRuns sets field value
func (o *InlineResponse2002Metrics) SetSuccessfulRuns(v int64) {
	o.SuccessfulRuns = v
}

// GetThroughput returns the Throughput field value
func (o *InlineResponse2002Metrics) GetThroughput() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Throughput
}

// GetThroughputOk returns a tuple with the Throughput field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Metrics) GetThroughputOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Throughput, true
}

// SetThroughput sets field value
func (o *InlineResponse2002Metrics) SetThroughput(v float32) {
	o.Throughput = v
}

// GetTotalCreditsUsed returns the TotalCreditsUsed field value
func (o *InlineResponse2002Metrics) GetTotalCreditsUsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalCreditsUsed
}

// GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Metrics) GetTotalCreditsUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCreditsUsed, true
}

// SetTotalCreditsUsed sets field value
func (o *InlineResponse2002Metrics) SetTotalCreditsUsed(v int64) {
	o.TotalCreditsUsed = v
}

// GetDurationMetrics returns the DurationMetrics field value
func (o *InlineResponse2002Metrics) GetDurationMetrics() InlineResponse2002MetricsDurationMetrics {
	if o == nil {
		var ret InlineResponse2002MetricsDurationMetrics
		return ret
	}

	return o.DurationMetrics
}

// GetDurationMetricsOk returns a tuple with the DurationMetrics field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Metrics) GetDurationMetricsOk() (*InlineResponse2002MetricsDurationMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationMetrics, true
}

// SetDurationMetrics sets field value
func (o *InlineResponse2002Metrics) SetDurationMetrics(v InlineResponse2002MetricsDurationMetrics) {
	o.DurationMetrics = v
}

func (o InlineResponse2002Metrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["success_rate"] = o.SuccessRate
	}
	if true {
		toSerialize["total_runs"] = o.TotalRuns
	}
	if true {
		toSerialize["failed_runs"] = o.FailedRuns
	}
	if true {
		toSerialize["successful_runs"] = o.SuccessfulRuns
	}
	if true {
		toSerialize["throughput"] = o.Throughput
	}
	if true {
		toSerialize["total_credits_used"] = o.TotalCreditsUsed
	}
	if true {
		toSerialize["duration_metrics"] = o.DurationMetrics
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2002Metrics struct {
	value *InlineResponse2002Metrics
	isSet bool
}

func (v NullableInlineResponse2002Metrics) Get() *InlineResponse2002Metrics {
	return v.value
}

func (v *NullableInlineResponse2002Metrics) Set(val *InlineResponse2002Metrics) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2002Metrics) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2002Metrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2002Metrics(val *InlineResponse2002Metrics) *NullableInlineResponse2002Metrics {
	return &NullableInlineResponse2002Metrics{value: val, isSet: true}
}

func (v NullableInlineResponse2002Metrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2002Metrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
