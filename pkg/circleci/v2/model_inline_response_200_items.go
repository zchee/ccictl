// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	"time"

	json "github.com/goccy/go-json"
)

// InlineResponse200Items struct for InlineResponse200Items
type InlineResponse200Items struct {
	// The name of the workflow.
	Name string `json:"name"`
	// The start of the aggregation window for workflow metrics.
	WindowStart time.Time `json:"window_start"`
	// The end of the aggregation window for workflow metrics.
	WindowEnd time.Time                `json:"window_end"`
	Metrics   InlineResponse200Metrics `json:"metrics"`
}

// NewInlineResponse200Items instantiates a new InlineResponse200Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200Items(name string, windowStart time.Time, windowEnd time.Time, metrics InlineResponse200Metrics) *InlineResponse200Items {
	this := InlineResponse200Items{}
	this.Name = name
	this.WindowStart = windowStart
	this.WindowEnd = windowEnd
	this.Metrics = metrics
	return &this
}

// NewInlineResponse200ItemsWithDefaults instantiates a new InlineResponse200Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200ItemsWithDefaults() *InlineResponse200Items {
	this := InlineResponse200Items{}
	return &this
}

// GetName returns the Name field value
func (o *InlineResponse200Items) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200Items) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineResponse200Items) SetName(v string) {
	o.Name = v
}

// GetWindowStart returns the WindowStart field value
func (o *InlineResponse200Items) GetWindowStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.WindowStart
}

// GetWindowStartOk returns a tuple with the WindowStart field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200Items) GetWindowStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowStart, true
}

// SetWindowStart sets field value
func (o *InlineResponse200Items) SetWindowStart(v time.Time) {
	o.WindowStart = v
}

// GetWindowEnd returns the WindowEnd field value
func (o *InlineResponse200Items) GetWindowEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.WindowEnd
}

// GetWindowEndOk returns a tuple with the WindowEnd field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200Items) GetWindowEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowEnd, true
}

// SetWindowEnd sets field value
func (o *InlineResponse200Items) SetWindowEnd(v time.Time) {
	o.WindowEnd = v
}

// GetMetrics returns the Metrics field value
func (o *InlineResponse200Items) GetMetrics() InlineResponse200Metrics {
	if o == nil {
		var ret InlineResponse200Metrics
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200Items) GetMetricsOk() (*InlineResponse200Metrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value
func (o *InlineResponse200Items) SetMetrics(v InlineResponse200Metrics) {
	o.Metrics = v
}

func (o InlineResponse200Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["window_start"] = o.WindowStart
	}
	if true {
		toSerialize["window_end"] = o.WindowEnd
	}
	if true {
		toSerialize["metrics"] = o.Metrics
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200Items struct {
	value *InlineResponse200Items
	isSet bool
}

func (v NullableInlineResponse200Items) Get() *InlineResponse200Items {
	return v.value
}

func (v *NullableInlineResponse200Items) Set(val *InlineResponse200Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200Items(val *InlineResponse200Items) *NullableInlineResponse200Items {
	return &NullableInlineResponse200Items{value: val, isSet: true}
}

func (v NullableInlineResponse200Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
