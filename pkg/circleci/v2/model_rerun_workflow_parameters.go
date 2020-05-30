// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// RerunWorkflowParameters The information you can supply when rerunning a workflow.
type RerunWorkflowParameters struct {
	// A list of job IDs to rerun.
	Jobs *[]string `json:"jobs,omitempty"`
	// Whether to rerun the workflow from the failed job. Mutually exclusive with the jobs parameter.
	FromFailed *bool `json:"from_failed,omitempty"`
}

// NewRerunWorkflowParameters instantiates a new RerunWorkflowParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRerunWorkflowParameters() *RerunWorkflowParameters {
	this := RerunWorkflowParameters{}
	return &this
}

// NewRerunWorkflowParametersWithDefaults instantiates a new RerunWorkflowParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRerunWorkflowParametersWithDefaults() *RerunWorkflowParameters {
	this := RerunWorkflowParameters{}
	return &this
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *RerunWorkflowParameters) GetJobs() []string {
	if o == nil || o.Jobs == nil {
		var ret []string
		return ret
	}
	return *o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RerunWorkflowParameters) GetJobsOk() (*[]string, bool) {
	if o == nil || o.Jobs == nil {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *RerunWorkflowParameters) HasJobs() bool {
	if o != nil && o.Jobs != nil {
		return true
	}

	return false
}

// SetJobs gets a reference to the given []string and assigns it to the Jobs field.
func (o *RerunWorkflowParameters) SetJobs(v []string) {
	o.Jobs = &v
}

// GetFromFailed returns the FromFailed field value if set, zero value otherwise.
func (o *RerunWorkflowParameters) GetFromFailed() bool {
	if o == nil || o.FromFailed == nil {
		var ret bool
		return ret
	}
	return *o.FromFailed
}

// GetFromFailedOk returns a tuple with the FromFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RerunWorkflowParameters) GetFromFailedOk() (*bool, bool) {
	if o == nil || o.FromFailed == nil {
		return nil, false
	}
	return o.FromFailed, true
}

// HasFromFailed returns a boolean if a field has been set.
func (o *RerunWorkflowParameters) HasFromFailed() bool {
	if o != nil && o.FromFailed != nil {
		return true
	}

	return false
}

// SetFromFailed gets a reference to the given bool and assigns it to the FromFailed field.
func (o *RerunWorkflowParameters) SetFromFailed(v bool) {
	o.FromFailed = &v
}

func (o RerunWorkflowParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Jobs != nil {
		toSerialize["jobs"] = o.Jobs
	}
	if o.FromFailed != nil {
		toSerialize["from_failed"] = o.FromFailed
	}
	return json.Marshal(toSerialize)
}

type NullableRerunWorkflowParameters struct {
	value *RerunWorkflowParameters
	isSet bool
}

func (v NullableRerunWorkflowParameters) Get() *RerunWorkflowParameters {
	return v.value
}

func (v *NullableRerunWorkflowParameters) Set(val *RerunWorkflowParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableRerunWorkflowParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableRerunWorkflowParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRerunWorkflowParameters(val *RerunWorkflowParameters) *NullableRerunWorkflowParameters {
	return &NullableRerunWorkflowParameters{value: val, isSet: true}
}

func (v NullableRerunWorkflowParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRerunWorkflowParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
