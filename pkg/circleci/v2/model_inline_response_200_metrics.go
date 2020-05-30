// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// InlineResponse200Metrics Metrics relating to a workflow's runs.
type InlineResponse200Metrics struct {
	// The ratio of successful runs / total runs.
	SuccessRate float32 `json:"success_rate"`
	// The total number of runs.
	TotalRuns int64 `json:"total_runs"`
	// The number of failed runs.
	FailedRuns int64 `json:"failed_runs"`
	// The number of successful runs.
	SuccessfulRuns int64 `json:"successful_runs"`
	// The average number of workflow runs per day.
	Throughput float32 `json:"throughput"`
	// The mean time to recovery (mean time between failures and their next success) in seconds.
	Mttr int64 `json:"mttr"`
	// The total credits consumed by the workflow in the aggregation window.
	TotalCreditsUsed int64                                   `json:"total_credits_used"`
	DurationMetrics  InlineResponse200MetricsDurationMetrics `json:"duration_metrics"`
}

// NewInlineResponse200Metrics instantiates a new InlineResponse200Metrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200Metrics(successRate float32, totalRuns int64, failedRuns int64, successfulRuns int64, throughput float32, mttr int64, totalCreditsUsed int64, durationMetrics InlineResponse200MetricsDurationMetrics) *InlineResponse200Metrics {
	this := InlineResponse200Metrics{}
	this.SuccessRate = successRate
	this.TotalRuns = totalRuns
	this.FailedRuns = failedRuns
	this.SuccessfulRuns = successfulRuns
	this.Throughput = throughput
	this.Mttr = mttr
	this.TotalCreditsUsed = totalCreditsUsed
	this.DurationMetrics = durationMetrics
	return &this
}

// NewInlineResponse200MetricsWithDefaults instantiates a new InlineResponse200Metrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200MetricsWithDefaults() *InlineResponse200Metrics {
	this := InlineResponse200Metrics{}
	return &this
}

// GetSuccessRate returns the SuccessRate field value
func (o *InlineResponse200Metrics) GetSuccessRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SuccessRate
}

// GetSuccessRateOk returns a tuple with the SuccessRate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200Metrics) GetSuccessRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessRate, true
}

// SetSuccessRate sets field value
func (o *InlineResponse200Metrics) SetSuccessRate(v float32) {
	o.SuccessRate = v
}

// GetTotalRuns returns the TotalRuns field value
func (o *InlineResponse200Metrics) GetTotalRuns() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalRuns
}

// GetTotalRunsOk returns a tuple with the TotalRuns field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200Metrics) GetTotalRunsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRuns, true
}

// SetTotalRuns sets field value
func (o *InlineResponse200Metrics) SetTotalRuns(v int64) {
	o.TotalRuns = v
}

// GetFailedRuns returns the FailedRuns field value
func (o *InlineResponse200Metrics) GetFailedRuns() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FailedRuns
}

// GetFailedRunsOk returns a tuple with the FailedRuns field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200Metrics) GetFailedRunsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailedRuns, true
}

// SetFailedRuns sets field value
func (o *InlineResponse200Metrics) SetFailedRuns(v int64) {
	o.FailedRuns = v
}

// GetSuccessfulRuns returns the SuccessfulRuns field value
func (o *InlineResponse200Metrics) GetSuccessfulRuns() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SuccessfulRuns
}

// GetSuccessfulRunsOk returns a tuple with the SuccessfulRuns field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200Metrics) GetSuccessfulRunsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessfulRuns, true
}

// SetSuccessfulRuns sets field value
func (o *InlineResponse200Metrics) SetSuccessfulRuns(v int64) {
	o.SuccessfulRuns = v
}

// GetThroughput returns the Throughput field value
func (o *InlineResponse200Metrics) GetThroughput() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Throughput
}

// GetThroughputOk returns a tuple with the Throughput field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200Metrics) GetThroughputOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Throughput, true
}

// SetThroughput sets field value
func (o *InlineResponse200Metrics) SetThroughput(v float32) {
	o.Throughput = v
}

// GetMttr returns the Mttr field value
func (o *InlineResponse200Metrics) GetMttr() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Mttr
}

// GetMttrOk returns a tuple with the Mttr field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200Metrics) GetMttrOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mttr, true
}

// SetMttr sets field value
func (o *InlineResponse200Metrics) SetMttr(v int64) {
	o.Mttr = v
}

// GetTotalCreditsUsed returns the TotalCreditsUsed field value
func (o *InlineResponse200Metrics) GetTotalCreditsUsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalCreditsUsed
}

// GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200Metrics) GetTotalCreditsUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCreditsUsed, true
}

// SetTotalCreditsUsed sets field value
func (o *InlineResponse200Metrics) SetTotalCreditsUsed(v int64) {
	o.TotalCreditsUsed = v
}

// GetDurationMetrics returns the DurationMetrics field value
func (o *InlineResponse200Metrics) GetDurationMetrics() InlineResponse200MetricsDurationMetrics {
	if o == nil {
		var ret InlineResponse200MetricsDurationMetrics
		return ret
	}

	return o.DurationMetrics
}

// GetDurationMetricsOk returns a tuple with the DurationMetrics field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200Metrics) GetDurationMetricsOk() (*InlineResponse200MetricsDurationMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationMetrics, true
}

// SetDurationMetrics sets field value
func (o *InlineResponse200Metrics) SetDurationMetrics(v InlineResponse200MetricsDurationMetrics) {
	o.DurationMetrics = v
}

func (o InlineResponse200Metrics) MarshalJSON() ([]byte, error) {
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
		toSerialize["mttr"] = o.Mttr
	}
	if true {
		toSerialize["total_credits_used"] = o.TotalCreditsUsed
	}
	if true {
		toSerialize["duration_metrics"] = o.DurationMetrics
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200Metrics struct {
	value *InlineResponse200Metrics
	isSet bool
}

func (v NullableInlineResponse200Metrics) Get() *InlineResponse200Metrics {
	return v.value
}

func (v *NullableInlineResponse200Metrics) Set(val *InlineResponse200Metrics) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200Metrics) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200Metrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200Metrics(val *InlineResponse200Metrics) *NullableInlineResponse200Metrics {
	return &NullableInlineResponse200Metrics{value: val, isSet: true}
}

func (v NullableInlineResponse200Metrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200Metrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
