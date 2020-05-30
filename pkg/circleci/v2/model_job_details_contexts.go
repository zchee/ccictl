// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// JobDetailsContexts Information about the context.
type JobDetailsContexts struct {
	// The name of the context.
	Name string `json:"name"`
}

// NewJobDetailsContexts instantiates a new JobDetailsContexts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobDetailsContexts(name string) *JobDetailsContexts {
	this := JobDetailsContexts{}
	this.Name = name
	return &this
}

// NewJobDetailsContextsWithDefaults instantiates a new JobDetailsContexts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobDetailsContextsWithDefaults() *JobDetailsContexts {
	this := JobDetailsContexts{}
	return &this
}

// GetName returns the Name field value
func (o *JobDetailsContexts) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobDetailsContexts) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JobDetailsContexts) SetName(v string) {
	o.Name = v
}

func (o JobDetailsContexts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableJobDetailsContexts struct {
	value *JobDetailsContexts
	isSet bool
}

func (v NullableJobDetailsContexts) Get() *JobDetailsContexts {
	return v.value
}

func (v *NullableJobDetailsContexts) Set(val *JobDetailsContexts) {
	v.value = val
	v.isSet = true
}

func (v NullableJobDetailsContexts) IsSet() bool {
	return v.isSet
}

func (v *NullableJobDetailsContexts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobDetailsContexts(val *JobDetailsContexts) *NullableJobDetailsContexts {
	return &NullableJobDetailsContexts{value: val, isSet: true}
}

func (v NullableJobDetailsContexts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobDetailsContexts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
