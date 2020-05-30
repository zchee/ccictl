// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// JobDetailsLatestWorkflow Info about the latest workflow the job was a part of.
type JobDetailsLatestWorkflow struct {
	// The unique ID of the workflow.
	Id string `json:"id"`
	// The name of the workflow.
	Name string `json:"name"`
}

// NewJobDetailsLatestWorkflow instantiates a new JobDetailsLatestWorkflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobDetailsLatestWorkflow(id string, name string) *JobDetailsLatestWorkflow {
	this := JobDetailsLatestWorkflow{}
	this.Id = id
	this.Name = name
	return &this
}

// NewJobDetailsLatestWorkflowWithDefaults instantiates a new JobDetailsLatestWorkflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobDetailsLatestWorkflowWithDefaults() *JobDetailsLatestWorkflow {
	this := JobDetailsLatestWorkflow{}
	return &this
}

// GetId returns the Id field value
func (o *JobDetailsLatestWorkflow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JobDetailsLatestWorkflow) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JobDetailsLatestWorkflow) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *JobDetailsLatestWorkflow) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobDetailsLatestWorkflow) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JobDetailsLatestWorkflow) SetName(v string) {
	o.Name = v
}

func (o JobDetailsLatestWorkflow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableJobDetailsLatestWorkflow struct {
	value *JobDetailsLatestWorkflow
	isSet bool
}

func (v NullableJobDetailsLatestWorkflow) Get() *JobDetailsLatestWorkflow {
	return v.value
}

func (v *NullableJobDetailsLatestWorkflow) Set(val *JobDetailsLatestWorkflow) {
	v.value = val
	v.isSet = true
}

func (v NullableJobDetailsLatestWorkflow) IsSet() bool {
	return v.isSet
}

func (v *NullableJobDetailsLatestWorkflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobDetailsLatestWorkflow(val *JobDetailsLatestWorkflow) *NullableJobDetailsLatestWorkflow {
	return &NullableJobDetailsLatestWorkflow{value: val, isSet: true}
}

func (v NullableJobDetailsLatestWorkflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobDetailsLatestWorkflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
