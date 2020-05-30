// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// JobDetailsParallelRuns Info about a status of the parallel run.
type JobDetailsParallelRuns struct {
	// Index of the parallel run.
	Index int64 `json:"index"`
	// Status of the parallel run.
	Status string `json:"status"`
}

// NewJobDetailsParallelRuns instantiates a new JobDetailsParallelRuns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobDetailsParallelRuns(index int64, status string) *JobDetailsParallelRuns {
	this := JobDetailsParallelRuns{}
	this.Index = index
	this.Status = status
	return &this
}

// NewJobDetailsParallelRunsWithDefaults instantiates a new JobDetailsParallelRuns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobDetailsParallelRunsWithDefaults() *JobDetailsParallelRuns {
	this := JobDetailsParallelRuns{}
	return &this
}

// GetIndex returns the Index field value
func (o *JobDetailsParallelRuns) GetIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *JobDetailsParallelRuns) GetIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *JobDetailsParallelRuns) SetIndex(v int64) {
	o.Index = v
}

// GetStatus returns the Status field value
func (o *JobDetailsParallelRuns) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *JobDetailsParallelRuns) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *JobDetailsParallelRuns) SetStatus(v string) {
	o.Status = v
}

func (o JobDetailsParallelRuns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableJobDetailsParallelRuns struct {
	value *JobDetailsParallelRuns
	isSet bool
}

func (v NullableJobDetailsParallelRuns) Get() *JobDetailsParallelRuns {
	return v.value
}

func (v *NullableJobDetailsParallelRuns) Set(val *JobDetailsParallelRuns) {
	v.value = val
	v.isSet = true
}

func (v NullableJobDetailsParallelRuns) IsSet() bool {
	return v.isSet
}

func (v *NullableJobDetailsParallelRuns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobDetailsParallelRuns(val *JobDetailsParallelRuns) *NullableJobDetailsParallelRuns {
	return &NullableJobDetailsParallelRuns{value: val, isSet: true}
}

func (v NullableJobDetailsParallelRuns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobDetailsParallelRuns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
