// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	"time"

	json "github.com/goccy/go-json"
)

// InlineResponse2002Items struct for InlineResponse2002Items
type InlineResponse2002Items struct {
	// The name of the job.
	Name string `json:"name"`
	// The start of the aggregation window for job metrics.
	WindowStart time.Time `json:"window_start"`
	// The end of the aggregation window for job metrics.
	WindowEnd time.Time                 `json:"window_end"`
	Metrics   InlineResponse2002Metrics `json:"metrics"`
}

// NewInlineResponse2002Items instantiates a new InlineResponse2002Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2002Items(name string, windowStart time.Time, windowEnd time.Time, metrics InlineResponse2002Metrics) *InlineResponse2002Items {
	this := InlineResponse2002Items{}
	this.Name = name
	this.WindowStart = windowStart
	this.WindowEnd = windowEnd
	this.Metrics = metrics
	return &this
}

// NewInlineResponse2002ItemsWithDefaults instantiates a new InlineResponse2002Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2002ItemsWithDefaults() *InlineResponse2002Items {
	this := InlineResponse2002Items{}
	return &this
}

// GetName returns the Name field value
func (o *InlineResponse2002Items) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Items) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineResponse2002Items) SetName(v string) {
	o.Name = v
}

// GetWindowStart returns the WindowStart field value
func (o *InlineResponse2002Items) GetWindowStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.WindowStart
}

// GetWindowStartOk returns a tuple with the WindowStart field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Items) GetWindowStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowStart, true
}

// SetWindowStart sets field value
func (o *InlineResponse2002Items) SetWindowStart(v time.Time) {
	o.WindowStart = v
}

// GetWindowEnd returns the WindowEnd field value
func (o *InlineResponse2002Items) GetWindowEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.WindowEnd
}

// GetWindowEndOk returns a tuple with the WindowEnd field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Items) GetWindowEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowEnd, true
}

// SetWindowEnd sets field value
func (o *InlineResponse2002Items) SetWindowEnd(v time.Time) {
	o.WindowEnd = v
}

// GetMetrics returns the Metrics field value
func (o *InlineResponse2002Items) GetMetrics() InlineResponse2002Metrics {
	if o == nil {
		var ret InlineResponse2002Metrics
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Items) GetMetricsOk() (*InlineResponse2002Metrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value
func (o *InlineResponse2002Items) SetMetrics(v InlineResponse2002Metrics) {
	o.Metrics = v
}

func (o InlineResponse2002Items) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse2002Items struct {
	value *InlineResponse2002Items
	isSet bool
}

func (v NullableInlineResponse2002Items) Get() *InlineResponse2002Items {
	return v.value
}

func (v *NullableInlineResponse2002Items) Set(val *InlineResponse2002Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2002Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2002Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2002Items(val *InlineResponse2002Items) *NullableInlineResponse2002Items {
	return &NullableInlineResponse2002Items{value: val, isSet: true}
}

func (v NullableInlineResponse2002Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2002Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
