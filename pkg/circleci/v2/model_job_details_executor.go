// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// JobDetailsExecutor Information about executor used for a job.
type JobDetailsExecutor struct {
	// Executor type.
	Type string `json:"type"`
	// Resource class name.
	ResourceClass string `json:"resource_class"`
}

// NewJobDetailsExecutor instantiates a new JobDetailsExecutor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobDetailsExecutor(type_ string, resourceClass string) *JobDetailsExecutor {
	this := JobDetailsExecutor{}
	this.Type = type_
	this.ResourceClass = resourceClass
	return &this
}

// NewJobDetailsExecutorWithDefaults instantiates a new JobDetailsExecutor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobDetailsExecutorWithDefaults() *JobDetailsExecutor {
	this := JobDetailsExecutor{}
	return &this
}

// GetType returns the Type field value
func (o *JobDetailsExecutor) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *JobDetailsExecutor) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *JobDetailsExecutor) SetType(v string) {
	o.Type = v
}

// GetResourceClass returns the ResourceClass field value
func (o *JobDetailsExecutor) GetResourceClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceClass
}

// GetResourceClassOk returns a tuple with the ResourceClass field value
// and a boolean to check if the value has been set.
func (o *JobDetailsExecutor) GetResourceClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceClass, true
}

// SetResourceClass sets field value
func (o *JobDetailsExecutor) SetResourceClass(v string) {
	o.ResourceClass = v
}

func (o JobDetailsExecutor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["resource_class"] = o.ResourceClass
	}
	return json.Marshal(toSerialize)
}

type NullableJobDetailsExecutor struct {
	value *JobDetailsExecutor
	isSet bool
}

func (v NullableJobDetailsExecutor) Get() *JobDetailsExecutor {
	return v.value
}

func (v *NullableJobDetailsExecutor) Set(val *JobDetailsExecutor) {
	v.value = val
	v.isSet = true
}

func (v NullableJobDetailsExecutor) IsSet() bool {
	return v.isSet
}

func (v *NullableJobDetailsExecutor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobDetailsExecutor(val *JobDetailsExecutor) *NullableJobDetailsExecutor {
	return &NullableJobDetailsExecutor{value: val, isSet: true}
}

func (v NullableJobDetailsExecutor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobDetailsExecutor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
